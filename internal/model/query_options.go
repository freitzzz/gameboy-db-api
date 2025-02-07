package model

type queryOrder int

const (
	QueryOrderNone queryOrder = iota
	QueryOrderRatingAsc
	QueryOrderRatingDesc
)

type QueryOptions struct {
	Count int
	Page  int
	Order queryOrder
	Name  string
}

func (o QueryOptions) ListOnly() bool {
	return len(o.Name) == 0
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
