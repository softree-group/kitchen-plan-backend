package entity

type Ingredient struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	Image         string  `json:"image,omitempty"`
	Measure       string  `json:"measure,omitempty"`
	Quantity      float32 `json:"quantity,omitempty"`
	Proteins      int     `json:"proteins,omitempty"`
	Fats          int     `json:"fats,omitempty"`
	Carbohydrates int     `json:"carbohydrates,omitempty"`
}

func (i *Ingredient) SetImageRoot(root string) {
	if i.Image == "" {
		return
	}

	i.Image = root + i.Image
}

var _ ImageRootSetter = &Ingredient{}
