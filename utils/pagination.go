package utils

import (
	"net/http"
	"strconv"
)

func GetPagination(r *http.Request) (int64, int64) {

	page, _ := strconv.Atoi(
		r.URL.Query().Get("page"),
	)

	limit, _ := strconv.Atoi(
		r.URL.Query().Get("limit"),
	)

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	skip := (page - 1) * limit

	return int64(limit), int64(skip)
}