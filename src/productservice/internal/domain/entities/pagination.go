package domain

type Pagination struct {
	CurrentPage int64
	TotalItems  int64
	TotalPages  int64
}

func NewPagination(currentPage, totalItems, totalPages int64) *Pagination {
	return &Pagination{
		CurrentPage: currentPage,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
	}
}
