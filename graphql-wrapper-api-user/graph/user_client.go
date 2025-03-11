package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"poc-graphql/graph/model"
	"poc-graphql/pkg/config"
)

type UserClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewUserClient(cfg config.Config) *UserClient {
	return &UserClient{
		BaseURL:    cfg.UserApiURL,
		HTTPClient: &http.Client{},
	}
}

func (c *UserClient) getAllUsers() ([]*model.User, error) {
	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/api/users", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *UserClient) AddUser(dto model.UserInput) (*model.User, error) {
	url := fmt.Sprintf("%s/api/user", c.BaseURL)

	payload, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
