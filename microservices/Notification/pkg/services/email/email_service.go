package email_service

import (
	"errors"
	"notificaiton/pkg/structs"
	"notificaiton/pkg/utils/common"
)

type EmailServiceInterface interface {
	SendEmail(emailFeatures structs.EmailFeatures) *common.ResponseMessage
	IsActiveProvider() bool
}

type EmailService struct {
	ActiveEmailService EmailServiceInterface
	Providers          []EmailServiceInterface
}

func (ss EmailService) addProviders() *EmailService {
	ss.Providers = append(ss.Providers, &SendGrid{})
	return &ss
}

func (ss EmailService) SetActiveEmailService() (*EmailService, error) {
	providers := ss.addProviders()
	for _, service := range providers.Providers {
		if service.IsActiveProvider() {
			return &EmailService{
				ActiveEmailService: service,
			}, nil
		}
	}

	return nil, errors.New("no active email service")
}
