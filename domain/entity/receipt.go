package entity

type Receipt struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Type        string       `json:"type"`
	Steps       []string     `json:"steps"`
	TimeToCook  int          `json:"time_to_cook"`
	Image       string       `json:"image"`
	Ingredients []Ingredient `json:"ingredients,omitempty"`
}

func (r *Receipt) SetImageRoot(root string) {
	r.Image = root + r.Image
}

var _ ImageRootSetter = &Receipt{}

type ReceiptFilter struct {
	Since       int
	Limit       int
	Title       string
	Type        string
	Ingredients []int
}
