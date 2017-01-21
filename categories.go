package cas

// SubCategory such as "イケボ（男子）" in the "声" category
type SubCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// Number of current live Movies
	Count int `json:"count"`
}

// Category such as "声"
type Category struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	SubCategories []SubCategory `json:"sub_categories"`
}

// CategoriesContainer holds a list of Categories
type CategoriesContainer struct {
	Categories []Category `json:"categories"`
}

// Categories returns a list of Categories and its SubCategories
func (api *Client) Categories() (*CategoriesContainer, error) {
	categories := &CategoriesContainer{}
	err := get(api, "/categories?lang=ja", categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
