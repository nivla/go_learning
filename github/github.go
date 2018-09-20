package github

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const API_ENDPOINT = "https://api.github.com/users/"

//Error is an api error
type Error struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Error implementation

func (e Error) Error() string {
	return e.Message
}

//github user api
type User struct {
	Login      string `json:"login"`
	Id         int    `json:"id"`
	Avatar_url string `json:"avatar_url"`
	Name       string `json:"name"`
	Html_url   string `json:"html_url"`
	Location   string `json:"location"`
	Followers  int    `json:"followers"`
	Following  int    `json:"following"`
}

type Users struct{}

func New() *Client {
	return &Client{}
}

type Client struct {
	Users
}

func (u *Users) GetUserByName(user_name string) (v *User, err error) {
	res, err := http.Get(API_ENDPOINT + user_name)

	if err != nil {
		return nil, errors.Wrap(err, "http error")
	}

	if res.StatusCode >= 400 {
		err := Error{
			Status: res.StatusCode,
		}

		if err := json.NewDecoder(res.Body).Decode(&err); err != nil {
			return nil, errors.Wrap(err, "decoding error")
		}

		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, errors.Wrap(err, "decoding error")
	}

	return
}
