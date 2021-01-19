package api

import (
	"encoding/json"
	"net/http"
)

const categoriesEndpoint string = theCatAPIBaseURL + "categories"

// Category represents a cat category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ListCategories represents all the categories available
// parameters limit and page must be greater than 0 to be taken into account
func (tca TheCatAPI) ListCategories(limit, page int) ([]Category, error) {
	queryParams := QueryParams{}
	if limit > 0 {
		queryParams["limit"] = string(limit)
	}
	if page > 0 {
		queryParams["page"] = string(page)
	}
	resp, err := sendRequest(tca,
		http.MethodGet,
		categoriesEndpoint,
		queryParams,
		nil)
	if err != nil {
		return nil, err
	}
	var categories []Category
	err = json.Unmarshal(resp, &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
