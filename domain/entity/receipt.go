package entity

type (
	Receipt struct {
		Id          int          `json:"id"`
		Title       string       `json:"title"`
		Type        string       `json:"type"`
		Steps       []string     `json:"steps"`
		TimeToCook  int          `json:"time_to_cook"`
		Image       string       `json:"image"`
		Ingredients []Ingredient `json:"ingredients,omitempty"`
	}
)
