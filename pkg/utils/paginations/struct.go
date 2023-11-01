// utils/paginations/struct.go
package paginations

import "gorm.io/gorm"

type Param struct {
	DB      *gorm.DB
	Query   *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

type Pagination struct {
	Count    int64       `json:"count"`
	Pages    int         `json:"pages"`
	Records  interface{} `json:"records"`
	Offset   int         `json:"offset"`
	Limit    int         `json:"limit"`
	Page     int         `json:"page"`
	PrevPage int         `json:"prev_page"`
	NextPage int         `json:"next_page"`
}
