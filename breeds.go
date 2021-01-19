package thecatapi

import (
	"encoding/json"
	"net/http"
)

const breedsEndpoint string = theCatAPIBaseURL + "breeds"

// Breed represents a breed from TheCatApi
// https://docs.thecatapi.com/api-reference/breeds
type Breed struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Temperament      string `json:"temperament"`
	LifeSpan         string `json:"life_span"`
	AltNames         string `json:"alt_names"`
	WikipediaURL     string `json:"wikipedia_url"`
	Origin           string `json:"origin"`
	WeightImperial   string `json:"weight_imperial"`
	Experimental     int    `json:"experimental"`
	Hairless         int    `json:"hairless"`
	Natural          int    `json:"natural"`
	Rare             int    `json:"rare"`
	Rex              int    `json:"rex"`
	SuppressTail     int    `json:"suppress_tail"`
	ShortLegs        int    `json:"short_legs"`
	Hypoallergenic   int    `json:"hypoallergenic"`
	Adaptability     int    `json:"adaptability"`
	AffectionLevel   int    `json:"affection_level"`
	CountryCode      string `json:"country_code"`
	ChildFriendly    int    `json:"child_friendly"`
	DogFriendly      int    `json:"dog_friendly"`
	EnergyLevel      int    `json:"energy_level"`
	Grooming         int    `json:"grooming"`
	HealthIssues     int    `json:"health_issues"`
	Intelligence     int    `json:"intelligence"`
	SheddingLevel    int    `json:"shedding_level"`
	SocialNeeds      int    `json:"social_needs"`
	StrangerFriendly int    `json:"strangers_friendly"`
	Vocalisation     int    `json:"vocalisation"`
}

// ListBreeds lists all breeds possible
func (tca TheCatAPI) ListBreeds() ([]Breed, error) {
	resp, err := sendRequest(tca, http.MethodGet, breedsEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	var breeds []Breed
	err = json.Unmarshal(resp, &breeds)
	if err != nil {
		return nil, err
	}
	return breeds, nil
}

// SearchBreedByName searches a breed by name
func (tca TheCatAPI) SearchBreedByName(q string) ([]Breed, error) {
	resp, err := sendRequest(tca, http.MethodGet, breedsEndpoint+"/search", QueryParams{"q": q}, nil)
	if err != nil {
		return nil, err
	}
	var breeds []Breed
	err = json.Unmarshal(resp, &breeds)
	if err != nil {
		return nil, err
	}
	return breeds, nil
}
