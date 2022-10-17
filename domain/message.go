package domain

import (
	"errors"
	"strings"
)

type Message struct {
	Condition                string                 `json:"condition,omitempty"`
	CollapseKey              string                 `json:"collapse_key,omitempty"`
	ContentAvailable         bool                   `json:"content_available,omitempty"`
	DelayWhileIdle           bool                   `json:"delay_while_idle,omitempty"`
	DeliveryReceiptRequested bool                   `json:"delivery_receipt_requested,omitempty"`
	DryRun                   bool                   `json:"dry_run,omitempty"`
	Data                     map[string]interface{} `json:"data,omitempty"`
	MutableContent           bool                   `json:"mutable_content,omitempty"`
	Notification             *Notification          `json:"notification,omitempty"`
	Priority                 string                 `json:"priority,omitempty"`
	RegistrationIDs          []string               `json:"registration_ids,omitempty"`
	RestrictedPackageName    string                 `json:"restricted_package_name,omitempty"`
	Token                    string                 `json:"to,omitempty"`
	TimeToLive               int                    `json:"time_to_live,omitempty"`
}

func (msg *Message) MessageValidate() error {

	if msg == nil {
		return errors.New("message is invalid")
	}

	inputCount := strings.Count(msg.Condition, "||") + strings.Count(msg.Condition, "&&")

	switch true {
	case msg.Token == "" && (msg.Condition == "" || inputCount > 2) && len(msg.RegistrationIDs) == 0:
		return errors.New("invalid topic or record identifiers not configured")
	case len(msg.RegistrationIDs) > 1000:
		return errors.New("too many record IDs")
	case msg.TimeToLive > 2419200:
		return errors.New("life time messages are not valid")
	default:
		return nil
	}

}
