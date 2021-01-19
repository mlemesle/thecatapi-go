package thecatapi

import (
	"encoding/json"
	"net/http"
)

const imagesEndpoint string = theCatAPIBaseURL + "images"

type ImageOrder string

const (
	RANDOM ImageOrder = "RANDOM"
	ASC               = "ASC"
	DESC              = "DESC"
	EMPTY             = ""
)

type Image struct {
	ID         string     `json:"id"`
	URL        string     `json:"url"`
	Categories []Category `json:"categories"`
	Breeds     []Breed    `json:"breeds"`
	Height     int        `json:"height"`
	Width      int        `json:"width"`
}

func (tca TheCatAPI) GetRandomPublicImage() ([]Image, error) {
	resp, err := sendRequest(tca,
		http.MethodGet,
		imagesEndpoint+"/search",
		nil,
		nil)
	if err != nil {
		return nil, err
	}
	var images []Image
	err = json.Unmarshal(resp, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}
