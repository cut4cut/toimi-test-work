package entity

const itemsPerPage uint64 = 10

func NewPagination() Pagination {
	return Pagination{
		ItemsPerPage: itemsPerPage,
	}
}

type Pagination struct {
	Page         uint64 `json:"page" form:"num"`
	DescPrice    bool   `json:"price" form:"price"`
	DescCreated  bool   `json:"created" form:"created"`
	ItemsPerPage uint64
}
