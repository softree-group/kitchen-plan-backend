package entity

type (
	Ingredient struct {
		Id       int     `json:"id"`
		Title    string  `json:"title"`
		Image    string  `json:"image,omitempty"`
		Measure  string  `json:"measure,omitempty"`
		Quantity float32 `json:"quantity,omitempty"`
	}
)
