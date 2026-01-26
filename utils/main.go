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

// GetPaginationStartAndEndIndices returns the starting and ending index, provided the page number and page size.
//
// The start index uint is calculated using the formula: (pageNumber - 1) * pageSize
//
// The end index uint is calculated using the formula: startIndex + pageSize
func GetPaginationStartAndEndIndices(pageNumber uint, pageSize uint) (uint, uint) {
	startIndex := (pageNumber - 1) * pageSize
	endIndex := startIndex + pageSize
	return startIndex, endIndex
}
