package freshchat

import (
	"net/url"
	"strconv"
)

const (
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

type Resp struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ListResp struct {
	Pagination Pagination `json:"pagination"`
	Links      Links      `json:"links"`
}

type Links struct {
	NextPage  Page `json:"next_page"`
	FirstPage Page `json:"first_page"`
	LastPage  Page `json:"last_page"`
}

type Page struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}

type Pagination struct {
	TotalItems   int64 `json:"total_items"`
	TotalPages   int64 `json:"total_pages"`
	CurrentPage  int64 `json:"current_page"`
	ItemsPerPage int64 `json:"items_per_page"`
}

type ListReq struct {
	Page         int
	ItemsPerPage int
	SortOrder    string
	SortBy       string
}

func (r *ListReq) Values() url.Values {
	values := make(url.Values, 4)
	if r.Page != 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.ItemsPerPage != 0 {
		values.Set("items_per_page", strconv.Itoa(r.ItemsPerPage))
	}
	if r.SortOrder != "" {
		values.Set("sort_order", r.SortOrder)
	}
	if r.SortBy != "" {
		values.Set("sort_by", r.SortBy)
	}
	return values
}
