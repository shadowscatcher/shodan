package models

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

type NotifierProviderType string
type ProviderArgs map[string]string

const (
	EmailProviderType     NotifierProviderType = "email"
	GitterProviderType    NotifierProviderType = "gitter"
	PagerdutyProviderType NotifierProviderType = "pagerduty"
	SlackProviderType     NotifierProviderType = "slack"
	TelegramProviderType  NotifierProviderType = "telegram"
	WebhookProviderType   NotifierProviderType = "webhook"
	PhoneProviderType     NotifierProviderType = "phone"
)

type NotifierProvider struct {
	Description  string
	ProviderType NotifierProviderType
	ProviderArgs
}

type NotifierResponse struct {
	SimpleResponse
	NotifierID string `json:"id"`
}

type NotifierDescriptor struct {
	Description string            `json:"description"`
	Provider    string            `json:"provider"`
	NotifierID  string            `json:"id"`
	Args        map[string]string `json:"args"`
}

type NotifierList struct {
	Notifiers []NotifierDescriptor `json:"matches"`
	Total     int                  `json:"total"`
}

type ProviderRequirements struct {
	RequiredFields []string `json:"required"`
}

func (p NotifierProvider) ToRequestBody() io.Reader {
	values := make(url.Values)
	if p.Description != "" {
		values.Set("description", p.Description)
	}
	values.Set("provider", string(p.ProviderType))

	for key, value := range p.ProviderArgs {
		values.Set(key, value)
	}

	return strings.NewReader(values.Encode())
}

func CreateEmailProvider(to, description string) (p NotifierProvider, err error) {
	return CreateProvider(EmailProviderType, description, map[string]string{
		"to": to,
	})
}

func CreatePhoneProvider(to, description string) (p NotifierProvider, err error) {
	return CreateProvider(PhoneProviderType, description, map[string]string{
		"to": to,
	})
}

func CreatePagerdutyProvider(routingKey, description string) (p NotifierProvider, err error) {
	return CreateProvider(PagerdutyProviderType, description, map[string]string{
		"routing_key": routingKey,
	})
}
func CreateGitterProvider(roomID, token, description string) (p NotifierProvider, err error) {
	return CreateProvider(GitterProviderType, description, map[string]string{
		"room_id": roomID,
		"token":   token,
	})
}

func CreateSlackProvider(webhookURL, description string) (p NotifierProvider, err error) {
	return CreateProvider(SlackProviderType, description, map[string]string{
		"webhook_url": webhookURL,
	})
}

func CreateTelegramProvider(chatID, token, description string) (p NotifierProvider, err error) {
	return CreateProvider(TelegramProviderType, description, map[string]string{
		"chat_id": chatID,
		"token":   token,
	})
}

func CreateWebhookProvider(URL, description string) (p NotifierProvider, err error) {
	return CreateProvider(WebhookProviderType, description, map[string]string{
		"url": URL,
	})
}

func CreateProvider(t NotifierProviderType, description string, requiredArgs ProviderArgs) (p NotifierProvider, err error) {
	for argName, arg := range requiredArgs {
		if arg == "" {
			err = fmt.Errorf("%s is required", argName)
		}
	}

	p.ProviderType = t
	p.Description = description
	p.ProviderArgs = requiredArgs

	return
}
