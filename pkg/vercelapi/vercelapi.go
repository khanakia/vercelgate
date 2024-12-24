package vercelapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetUser(token string) (*User, error) {
	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}

	url := "https://api.vercel.com/v2/user"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("unable to fetch package info")
	}

	// fmt.Println(string(body))

	var record *GetUserResponse
	err = json.Unmarshal(body, &record)
	if err != nil {
		return nil, errors.New("unable to parse json")
	}

	if len(record.Error.Code) > 0 {
		return nil, errors.New(record.Error.Message)
	}

	return &record.User, nil
}

func GetTeams(token string) ([]Team, error) {
	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}

	url := "https://api.vercel.com/v2/teams"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("unable to fetch package info")
	}

	// fmt.Println(string(body))

	var record *GetTeamsResponse
	err = json.Unmarshal(body, &record)
	if err != nil {
		return nil, errors.New("unable to parse json")
	}

	if len(record.Error.Code) > 0 {
		return nil, errors.New(record.Error.Message)
	}

	return record.Teams, nil
}

type GetUserResponse struct {
	User  User  `json:"user"`
	Error Error `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Avatar        any    `json:"avatar"`
	DefaultTeamID string `json:"defaultTeamId"`
	Version       string `json:"version"`
	CreatedAt     int64  `json:"createdAt"`
}

type Team struct {
	ID        string    `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Avatar    any       `json:"avatar"`
	CreatedAt int64     `json:"createdAt"`
	Created   time.Time `json:"created"`
}

type Pagination struct {
	Count int   `json:"count"`
	Next  int64 `json:"next"`
	Prev  int64 `json:"prev"`
}

type GetTeamsResponse struct {
	Teams      []Team     `json:"teams"`
	Pagination Pagination `json:"pagination"`
	Error      Error      `json:"error,omitempty"`
}
