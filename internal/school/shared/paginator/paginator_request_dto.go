package paginator

import (
	"github.com/go-playground/validator"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
)

type ColumnSearch struct {
	Column string `json:"column" validate:"omitempty"`
	Value  string `json:"value" validate:"omitempty"`
}

type PaginatorRequest struct {
	Page         int            `json:"page" validate:"required,numeric"`
	Sort         string         `json:"sort" validate:"dbOrder"`
	Search       string         `json:"search_term" validate:"omitempty,alphanum"`
	SortField    string         `json:"sort_field" validate:"omitempty"`
	ColumnSearch []ColumnSearch `json:"column_search" validate:"omitempty,dive"`
	Limit        int            `json:"limit" validate:"required,numeric"`
}

func (p *PaginatorRequest) Validate() error {
	val := validator.New()
	_ = val.RegisterValidation("dbOrder", requestvalidator.DbOrder)
	return val.Struct(p)
}
