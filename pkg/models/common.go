package models

import (
	"math"

	"github.com/AlekSi/pointer"
)

type PaginationOptions struct {
	Page         *int     `query:"page" json:"page"`
	Limit        *int     `query:"limit" json:"limit"`
	Order        *string  `query:"order" json:"order"`
	TotalPage    *int     `query:"total_pages" json:"total_page"`
	Total        *int64   `query:"total" json:"total"`
	Search       *string  `query:"search" json:"search"`
	SearchFields []string `json:"-"`
}

func (p *PaginationOptions) DefaultLimit(limit int) {
	if p.Limit == nil || p.GetLimit() == 0 {
		p.Limit = pointer.ToInt(limit)
	} else if *p.Limit > 100 {
		p.Limit = pointer.ToInt(100)
	}
}

func (p *PaginationOptions) DefaultPage(page int) {
	if p.Page == nil {
		p.Page = pointer.ToInt(page)
	}
}

func (p *PaginationOptions) DefaultTotal() {
	p.Total = pointer.ToInt64(0)
}

func (p *PaginationOptions) DefaultTotalPage() {
	p.TotalPage = pointer.ToInt(0)
}

func (p *PaginationOptions) GetLimit() int {
	if p.Limit == nil {
		return 0
	}
	return *p.Limit
}

func (p *PaginationOptions) GetPage() int {
	if p.Page == nil {
		return 1
	}
	return *p.Page
}

func (p *PaginationOptions) GetTotal() int64 {
	if p.Total == nil {
		return int64(*p.Limit)
	}
	return *p.Total
}

func (p *PaginationOptions) GetTotalPage() int {
	if p.TotalPage == nil {
		return 1
	}
	return *p.TotalPage
}

func (p *PaginationOptions) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationOptions) SetPage(page int) {
	p.Page = &page
}

func (p *PaginationOptions) SetLimit(limit int) {
	p.Limit = &limit
}

func (p *PaginationOptions) SetOrder(order string) {
	p.Order = &order
}

func (p *PaginationOptions) SetSearch(search string) {
	p.Search = &search
}

func (p *PaginationOptions) SetTotal(total int64) {
	p.Total = &total
}

func (p *PaginationOptions) SetTotalPage(totalPage int) {
	p.TotalPage = &totalPage
}

func (p *PaginationOptions) SetSearchFields(fields []string) {
	p.SearchFields = fields
}

func (p *PaginationOptions) AddSearchFields(fields []string) {
	if p.SearchFields == nil {
		p.SearchFields = fields
	} else {
		p.SearchFields = append(p.SearchFields, fields...)
	}
}

func (p *PaginationOptions) CalTotalPage() {
	totalF := float64(p.GetTotal())
	limitF := float64(p.GetLimit())
	if limitF > totalF || limitF <= 0 {
		p.SetTotalPage(1)
	} else {
		totalPageF := math.Ceil(totalF / limitF)
		p.SetTotalPage(int(totalPageF))
	}
}
