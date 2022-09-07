package entity

import "fmt"

type AdminsQuery struct {
	Code       string   `json:"code,omitempty"`
	Additional []string `json:"additional,omitempty"`
}

func (q *AdminsQuery) Validate() error {
	if q.Code == "" {
		return fmt.Errorf("code is empty")
	}

	return nil
}
