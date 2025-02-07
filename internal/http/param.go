package http

type ratingFilter string

const (
	ratingFilterHighestRated ratingFilter = "high"
	ratingFilterLowestRated  ratingFilter = "low"
)

const (
	ratingFilterQueryParam = "rating"
	nameFilterQueryParam   = "name"
	pageFilterQueryParam   = "page"
)
