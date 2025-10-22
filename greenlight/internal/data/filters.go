package data

import (
	"github.com/rigoncs/lets-go-further/internal/validator"
	"slices"
	"strings"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSaftList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")

	v.Check(validator.PermittedValue(f.Sort, f.SortSaftList...), "sort", "invalid sort value")
}

func (f Filters) sortColumn() string {
	if slices.Contains(f.SortSaftList, f.Sort) {
		return strings.TrimPrefix(f.Sort, "-")
	}
	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}
