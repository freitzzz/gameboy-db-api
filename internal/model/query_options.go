package model

type queryOrder int

const (
	QueryOrderNone queryOrder = iota
	QueryOrderRatingAsc
	QueryOrderRatingDesc
)

type QueryOptions struct {
	Count int
	Order queryOrder
}

func (o queryOrder) None() bool {
	return o == QueryOrderNone
}

func (o queryOrder) HighestRating() bool {
	return o == QueryOrderRatingDesc
}

func (o queryOrder) LowestRating() bool {
	return o == QueryOrderRatingAsc
}
