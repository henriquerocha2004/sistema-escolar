package common

import (
	"fmt"
	"strconv"

	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
)

type Pagination struct {
	Limit        int
	offSet       int
	Sort         string
	Search       string
	SortField    string
	ColumnSearch []dto.ColumnSearch
}

func (p *Pagination) SetPage(page int) {
	if page == 0 {
		page = 1
	}

	p.offSet = (page * p.Limit) - p.Limit
}

func (p *Pagination) GetOffset() int {
	return p.offSet
}

func (p *Pagination) FillFromDto(dtoRequest dto.PaginatorRequest) {
	p.ColumnSearch = dtoRequest.ColumnSearch
	p.Limit = dtoRequest.Limit
	p.Search = dtoRequest.Search
	p.SetPage(dtoRequest.Page)
	p.SortField = dtoRequest.SortField
	p.Sort = dtoRequest.Sort
}

func (p *Pagination) FiltersInSql() string {
	query := ""

	if len(p.ColumnSearch) > 0 {
		for _, column := range p.ColumnSearch {
			query += fmt.Sprintf(" AND %s = '%s' ", column.Column, column.Value)
		}
	}

	if p.Sort != "" && p.SortField != "" {
		order := fmt.Sprintf(" ORDER BY %s %s", p.SortField, p.Sort)
		query += order
	}

	if p.Limit != 0 {
		limit := fmt.Sprintf(" LIMIT %s", strconv.Itoa(p.Limit))
		query += limit
	}

	if p.GetOffset() != 0 {
		offset := fmt.Sprintf(" OFFSET %s", strconv.Itoa(p.GetOffset()))
		query += offset
	}

	return query
}
