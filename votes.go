package thecatapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const votesEndpoint string = theCatAPIBaseURL + "votes"

type Vote struct {
	Value       int    `json:"value"`
	ImageID     string `json:"image_id"`
	SubID       string `json:"sub_id"`
	CreatedAt   string `json:"created_at"`
	ID          string `json:"id"`
	CountryCode string `json:"country_code"`
}

type VoteResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func (tca TheCatAPI) GetAllVotes(subID string, limit, page int) ([]Vote, error) {
	queryParams := QueryParams{}
	if subID == "" {
		queryParams["sub_id"] = string(subID)
	}
	if limit > 0 {
		queryParams["limit"] = string(limit)
	}
	if page > 0 {
		queryParams["page"] = string(page)
	}
	resp, err := sendRequest(tca,
		http.MethodGet,
		votesEndpoint,
		queryParams,
		nil)
	if err != nil {
		return nil, err
	}
	var votes []Vote
	err = json.Unmarshal(resp, &votes)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

type postVoteBody struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
	Value   int    `json:"value"`
}

func (tca TheCatAPI) PostVote(imageID, subID string, value int) (VoteResponse, error) {
	if imageID == "" || (value != 0 && value != 1) {
		return VoteResponse{}, errors.New("imageID and value are mandatory")
	}
	postVoteBody := postVoteBody{ImageID: imageID,
		SubID: subID,
		Value: value}
	postVoteBodyBytes, err := json.Marshal(postVoteBody)
	if err != nil {
		return VoteResponse{}, err
	}
	body := bytes.NewReader(postVoteBodyBytes)
	resp, err := sendRequest(tca,
		http.MethodPost,
		votesEndpoint,
		nil,
		body)
	if err != nil {
		return VoteResponse{}, err
	}
	var voteResponse VoteResponse
	err = json.Unmarshal(resp, &voteResponse)
	if err != nil {
		return VoteResponse{}, err
	}
	return voteResponse, nil
}
