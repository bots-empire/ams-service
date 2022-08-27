package entity

import "fmt"

type Access struct {
	UserID     int64    `json:"user_id,omitempty"`
	Code       string   `json:"code,omitempty"`
	Additional []string `json:"additional,omitempty"`

	UserName      string `json:"user_name,omitempty"`
	UserFirstName string `json:"user_first_name,omitempty"`
	UserLastName  string `json:"user_last_name,omitempty"`
}

func (a *Access) Validate(userInfo bool) error {
	if a.UserID == 0 {
		return fmt.Errorf("user_id is empty")
	}

	if a.Code == "" {
		return fmt.Errorf("code is empty")
	}

	if !userInfo {
		return nil
	}

	if a.UserName == "" {
		return fmt.Errorf("user_name is empty")
	}

	if a.UserFirstName == "" && a.UserLastName == "" {
		return fmt.Errorf("both of user_first_name and user_last_name is empty")
	}

	return nil
}

func (a *Access) MergeAccess(newRules *Access) {
	if len(a.Additional) != 0 {
		a.Additional = mergeStringArray(a.Additional, newRules.Additional)
	} else {
		a.Additional = newRules.Additional
	}

	a.UserName = newRules.UserName
	a.UserFirstName = newRules.UserFirstName
	a.UserLastName = newRules.UserLastName
}

func mergeStringArray(source, new []string) []string {
	uniqueLabels := make(map[string]struct{}, len(source)+len(new))

	for _, label := range source {
		uniqueLabels[label] = struct{}{}
	}
	for _, label := range new {
		uniqueLabels[label] = struct{}{}
	}

	uniqueArray := make([]string, len(uniqueLabels))
	i := 0
	for label := range uniqueLabels {
		uniqueArray[i] = label
		i++
	}

	return uniqueArray
}
