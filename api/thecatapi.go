package api

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

const theCatAPIBaseURL string = "https://api.thecatapi.com/v1/"

// QueryParams represents http params
type QueryParams map[string]string

// TheCatAPI holds the needed informations for the API
type TheCatAPI struct {
	baseURL string
	aPIKey  string
	client  *http.Client
}

// Interface represents the methods that TheCatAPI should implement
type Interface interface {
	ListBreeds() ([]Breed, error)
	SearchBreedByName(q string) ([]Breed, error)

	ListCategories(limit, page int) ([]Category, error)

	GetAllVotes(subID string, limit, page int) ([]Vote, error)
	PostVote(imageID, subID string, value int) (VoteResponse, error)

	GetRandomPublicImage() ([]Image, error)
}

// NewTheCatAPI creates a new instance using the api key given in parameter
func NewTheCatAPI(apiKey string) (*TheCatAPI, error) {
	if apiKey == "" {
		return nil, errors.New("No api key provided")
	}
	return &TheCatAPI{
		baseURL: theCatAPIBaseURL,
		aPIKey:  apiKey,
		client:  &http.Client{},
	}, nil
}

func (tca TheCatAPI) getNewRequest(method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("x-api-key", tca.aPIKey)
	return req
}

func sendRequest(tca TheCatAPI, method string, url string, params map[string]string, body io.Reader) ([]byte, error) {
	req := tca.getNewRequest(method, url, body)
	if params != nil {
		query := req.URL.Query()
		for key, value := range params {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}
	resp, err := tca.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
