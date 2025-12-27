package utils

type ListResponse[K any] struct {
	Data  []K `json:"data"`
	Count int `json:"count"`
}

func CreateListResponse[T any](data []T) ListResponse[T] {
	return ListResponse[T]{
		Data:  data,
		Count: len(data),
	}
}
