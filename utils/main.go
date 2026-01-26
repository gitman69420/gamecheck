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

// GetPaginationStartIndex returns the starting index, provided the page number and page size.
//
// The returned uint is calculated with the formula: (pageNumber-1)*pageSize
func GetPaginationStartIndex(pageNumber uint, pageSize uint) uint {
	return (pageNumber - 1) * pageSize
}
