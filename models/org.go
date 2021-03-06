package models

import (
	"fmt"
)

type Organization struct {
	UUID         string   `json:"UUID"`         //
	Organization string   `json:"organization"` //
	Username     string   `json:"username"`     //
	Description  string   `json:"description"`  //
	Repositories []string `json:"repositories"` //
	Created      int64    `json:"created"`      //
	Updated      int64    `json:"updated"`      //
	Teams        []string `json:"teams"`        //
	Memo         []string `json:"memo"`         //
}

type Team struct {
	UUID              string       `json:"UUID"`           //
	Team              string       `json:"team"`           //
	Organization      string       `json:"organization"`   //
	Username          string       `json:"username"`       //
	Description       string       `json:"description"`    //
	Users             []string     `json:"users"`          //
	TeamPrivileges    []string     `json:"teamprivileges"` //
	Repositories      []string     `json:"repositories"`   //
	Memo              []string     `json:"memo"`           //
	RepositoryObjects []Repository `json:"repositoryobjects"`
	UserObjects       []User       `json:"userobjects"`
}

func (organization *Organization) Has(organizationName string) (bool, []byte, error) {
	UUID, err := GetUUID("organization", organizationName)
	if err != nil {
		return false, nil, err
	}
	if len(UUID) <= 0 {
		return false, nil, nil
	}

	err = Get(organization, UUID)

	return true, UUID, err
}

func (organization *Organization) Save() error {
	if err := Save(organization, []byte(organization.UUID)); err != nil {
		return err
	}

	if _, err := LedisDB.HSet([]byte(GLOBAL_ORGANIZATION_INDEX), []byte(organization.Organization), []byte(organization.UUID)); err != nil {
		return err
	}

	return nil
}

func (organization *Organization) Get(UUID string) error {
	if err := Get(organization, []byte(UUID)); err != nil {
		return err
	}

	return nil
}

func (organization *Organization) Remove() error {
	if _, err := LedisDB.HSet([]byte(fmt.Sprintf("%s_remove", GLOBAL_ORGANIZATION_INDEX)), []byte(organization.Organization), []byte(organization.UUID)); err != nil {
		return err
	}

	if _, err := LedisDB.HDel([]byte(GLOBAL_ORGANIZATION_INDEX), []byte(organization.UUID)); err != nil {
		return err
	}

	return nil
}

func (team *Team) Has(teamName string) (bool, []byte, error) {
	UUID, err := GetUUID("team", teamName)
	if err != nil {
		return false, nil, err
	}

	if len(UUID) <= 0 {
		return false, nil, nil
	}

	err = Get(team, UUID)

	return true, UUID, err
}

func (team *Team) Save() error {
	if err := Save(team, []byte(team.UUID)); err != nil {
		return err
	}

	if _, err := LedisDB.HSet([]byte(GLOBAL_TEAM_INDEX), []byte(team.Team), []byte(team.UUID)); err != nil {
		return err
	}

	return nil
}

func (team *Team) Get(UUID string) error {
	if err := Get(team, []byte(UUID)); err != nil {
		return err
	}

	return nil
}

func (team *Team) Remove() error {
	if _, err := LedisDB.HSet([]byte(fmt.Sprintf("%s_remove", GLOBAL_TEAM_INDEX)), []byte(team.Team), []byte(team.UUID)); err != nil {
		return err
	}

	if _, err := LedisDB.HDel([]byte(GLOBAL_TEAM_INDEX), []byte(team.UUID)); err != nil {
		return err
	}

	return nil
}
