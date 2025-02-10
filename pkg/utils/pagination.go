package utils

import (
	"line-bk-api/config"
	"math"
)

func GetOffset(page, limit int) (int, int) {
	// if page is less than 1, set it to default page
	if page < 1 {
		page = config.DefaultPage
	}

	// get limit
	limit = GetLimit(limit)

	// calculate offset
	return (page - 1) * limit, limit
}

func GetLimit(limit int) int {

	// if limit is less than 1, set it to default limit
	if limit < 1 {
		limit = config.DefaultLimit
	}

	// if limit is greater than max limit, set it to max limit
	if limit > config.MaxLimit {
		limit = config.MaxLimit
	}
	return limit
}

func GetTotalPages(total int, limit int) int {
	// calculate total pages
	return int(math.Ceil(float64(total) / float64(limit)))
}

func GetNextPage(page int, totalPages int) int {
	// if page is less than total pages, return next page
	if page < totalPages {
		return page + 1
	}
	return 0
}

func GetPreviousPage(page int) int {
	// if page is greater than 1, return previous page
	if page > 1 {
		return page - 1
	}
	return 0
}
