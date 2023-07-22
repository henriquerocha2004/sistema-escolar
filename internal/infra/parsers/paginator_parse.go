package parsers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"strconv"
)

func ParseRequestPaginator(ctx *fiber.Ctx) (*dto.PaginatorRequest, error) {
	columnSearch := ctx.Query("column_search")
	var columnSearchDto []dto.ColumnSearch

	if columnSearch != "" {
		err := json.Unmarshal([]byte(columnSearch), &columnSearchDto)
		if err != nil {
			return nil, errors.New("failed to parse column_search")
		}
	}

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	paginatorRequestDto := dto.PaginatorRequest{
		Limit:        limit,
		Page:         page,
		Search:       ctx.Query("search_term"),
		Sort:         ctx.Query("sort"),
		SortField:    ctx.Query("sort_field"),
		ColumnSearch: columnSearchDto,
	}

	return &paginatorRequestDto, nil
}
