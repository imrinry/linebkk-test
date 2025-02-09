package utils

import "line-bk-api/config"

func GetOffset(page, limit int) int {
	if page < 1 {
		page = config.DefaultPage
	}

	if limit < 1 {
		limit = config.DefaultLimit
	}
	return (page - 1) * limit
}
