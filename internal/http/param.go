package http

type listingFilter string

const (
	listingFilterHighestRated listingFilter = "highest-rated"
	listingFilterLowestRated  listingFilter = "lowest-rated"
)

const (
	filterQueryParam = "filter"
)
