// Copyright (c) HopopOps
// SPDX-License-Identifier: MPL-2.0

package cloudbeaver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Team -
type Team struct {
	TeamId          string   `json:"teamId"`
	TeamName        string   `json:"teamName"`
	Description     string   `json:"description"`
	TeamPermissions []string `json:"teamPermissions"`
}

// GetAllTeams - Returns all teams
func (c *Client) GetAllTeams(authToken *string) (*[]Team, error) {
	rb, err := json.Marshal(NewGetAllTeams())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.HostURL, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	body, _, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gar := GetAllResponse{}
	err = json.Unmarshal(body, &gar)
	if err != nil {
		return nil, err
	}
	return &gar.Data.Teams, nil
}

// GetTeam - Returns a specifc team
func (c *Client) GetTeam(teamId string, authToken *string) (*Team, error) {
	rb, err := json.Marshal(NewGetTeams(&teamId))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.HostURL, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	body, _, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gr := GetResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}

	if len(gr.Data.Teams) != 1 {
		return nil, fmt.Errorf("expected 1 team, got %d", len(gr.Data.Teams))
	}

	return &gr.Data.Teams[0], nil
}

// CreateTeam - Create new team
func (c *Client) CreateTeam(teamId, teamName, description string, authToken *string) (*Team, error) {
	rb, err := json.Marshal(NewCreateTeam(&teamId, &teamName, &description))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.HostURL, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	body, _, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	r := CreateResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data.Team, nil
}

// UpdateTeam - Updates a team
func (c *Client) UpdateTeam(teamId, teamName, description string, authToken *string) (*Team, error) {
	rb, err := json.Marshal(NewUpdateTeam(&teamId, &teamName, &description))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.HostURL, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	body, _, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	r := UpdateResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data.Team, nil
}

// DeleteTeam - Deletes a team
func (c *Client) DeleteTeam(teamId string, authToken *string) error {
	rb, err := json.Marshal(NewDeleteTeam(&teamId))
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.HostURL, strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	body, _, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	r := DeleteResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}
	return nil
}
