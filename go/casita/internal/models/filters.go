package models

import "casita/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 1_000, "page", "must be a max of one thousand")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 50, "page_size", "must be a max of fifty")

	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}
