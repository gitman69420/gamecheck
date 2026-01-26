# How we're caching RAWG API search responses

A RAWG API call to `/games?search="GTA"` looks like this:
```json
{
    "count": 20,
    "next": "...",
    "prev": "...",
    "results": [
        {
            "id": 1234,
            "name": "Grand Theft Auto V"
            // ...
        },
        {
            "id": 123,
            "name": "Grand Theft Auto IV"
            // ...
        }
    ]
}
```

## Game Cache Key

We also create **game cache keys** for each game, like so:
```go
"game:1234" =>  "{\"id\": 1234,\"name\": \"Grand Theft Auto V\"}"

"game:123"  =>  "{\"id\": 123,\"name\": \"Grand Theft Auto IV\"}"
```
Each value is a stringified JSON object


## Search Cache Key

For each search, we create a **search cache key**
```go
// For e.g. if the user calls the API with the keyword "GTA", we create a cache key
"search:gta"
```


## Search Cache Value (Sorted Set)

To maintain the order of the search responses, we store a [sorted set](https://valkey.io/topics/sorted-sets/) of **game cache keys** for the **search cache key** where the scores are the index of the search and the values are the **game cache keys**.

```go
// a rough representation
type SortedSet ...


cache := map[string]SortedSet {
    "search:gta": SortedSet{
        0: "game:1234",
        1: "game:123",
    },
    "search:gtav": SortedSet{
        0: "game:1234",
    },
}
```


## How we fetch data in cache

We get pagination parameters from the request: `page` and `size`.
If not, we have constant defaults these parameters

We calculate a start and end index from these 2 values: 
```go
startIndex := (page - 1) * size
endIndex := startIndex + size
```

We access the **search cache key** sorted set ranged by the above start and end indeces.

If we get an empty result set, then we continue to fetch from the API.


## How we avoid empty API responses due to large pagination params

There's a possibility where a user can basically keep doing `/search/games?keyword=gta&page=9999&limit=1000000`.

Even though this may return empty responses, it still exhausts our RAWG API credits.

So we make a new cache key called "Search Maximum Index key"
```go
"search:gta:max-index"
```

This key will hold a value which is potentially the last index value for that particular search. So if the keyword "gta" has 2000 results overall, this value would be either 2000 or a value above it.


### How we should update the **Search Maximum Index**
Initially this key would be empty.

Whenever we make a call to the RAWG API, and get an empty response due to the page and size parameters exceeding the maximum index, we use it as a trigger to update this key.

```go
"/search/games?search=gta&page=20&size=10"
// existing search maximum index is null
// make RAWG API call => empty response
// so we update search maximum index to start index value = 190

"/search/games?search=gta&page=15&size=10"
// existing search maximum index is 190
// start index from request params is 140
// make RAWG API call => non empty response
// do nothing

"/search/games?search=gta&page=21&size=10"
// existing search maximum index is 190
// start index would be (21-1)*10 = 200
// since 200 > 190, we avoid the RAWG API call and return an empty response

"/search/games?search=gta&page=17&size=10"
// existing search maximum index is 190
// start index would be (17-1)*10 = 160
// make RAWG API call => empty response
// we update the search maximum index to 160 and return an empty response
```


### Final edge case to update **Search Maximum Index**
Whenever we make a valid RAWG API call, we also check the `count` field in the response.

We compare this with the `size` param we pass, and we can conclude that we've seen the last record.

For e.g. 
```
"api.rawg.com/games?search=gta&page=12&size=20" => [17 records ...]
```

We can safely conclude that we reached the last page because the page size and the response size are different, but not empty.

We do one last update to the Search Maximum Index, setting it to `(startIndex + responseCount)` which is (12-1)*20) + 17 = **237**

```
SET search:gta:max-index 237
```
We'd only need to invalidate this key only if more data is added to the RAWG API.