package paginations

import (
	"math"

	"gorm.io/gorm"
)

var DEFAULT_PAGE_SIZE = 10
var DEFAULT_PAGE = 1

func Pagging(p *Param, data interface{}) (*Pagination, error) {
	db := p.DB
	query := p.Query

	if p.ShowSQL {
		query = query.Debug()
	}
	if p.Page < 1 {
		p.Page = DEFAULT_PAGE
	}
	if p.Limit == 0 {
		p.Limit = DEFAULT_PAGE_SIZE
	}
	if len(p.OrderBy) > 0 {
		for _, order := range p.OrderBy {
			query = query.Order(order)
		}
	}

	done := make(chan bool, 1)
	var pagination Pagination
	var count int64
	var offset int

	// count total record
	go CountRecords(db, data, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	if err := query.Limit(p.Limit).Offset(offset).Find(data).Error; err != nil {
		<-done
		return nil, err
	}
	<-done

	pagination.Count = count
	pagination.Records = data
	pagination.Page = p.Page

	pagination.Offset = offset
	pagination.Limit = p.Limit
	pagination.Pages = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		pagination.PrevPage = p.Page - 1
	} else {
		pagination.PrevPage = p.Page
	}

	if p.Page >= pagination.Pages {
		pagination.NextPage = p.Page
	} else {
		pagination.NextPage = p.Page + 1
	}
	return &pagination, nil
}

func CountRecords(db *gorm.DB, data interface{}, done chan bool, count *int64) {
	db.Model(data).Count(count)
	done <- true
}
