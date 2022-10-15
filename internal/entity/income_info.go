package entity

import "fmt"

type IncomeInfo struct {
	UserID       int64  `json:"user_id,omitempty"`
	BotLink      string `json:"bot_link,omitempty"`
	BotName      string `json:"bot_name,omitempty"`
	IncomeSource string `json:"income_source,omitempty"`
	TypeBot      string `json:"type_bot,omitempty"`
}

func (i *IncomeInfo) ValidateAdd() error {
	if i.UserID == 0 {
		return fmt.Errorf("user_id is empty")
	}

	if i.IncomeSource == "" {
		return fmt.Errorf("income_source is empty")
	}

	if i.BotLink == "" {
		return fmt.Errorf("bot_link is empty")
	}

	return nil
}

func (i *IncomeInfo) ValidateGet() error {
	if i.UserID == 0 {
		return fmt.Errorf("user_id is empty")
	}

	if i.TypeBot == "" {
		return fmt.Errorf("type_bot is empty")
	}

	return nil
}
