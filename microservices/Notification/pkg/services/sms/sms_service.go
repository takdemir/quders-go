package sms_service

import (
	"errors"
	"notificaiton/pkg/utils/common"
)

type SmsServiceInterface interface {
	SendSms(service SmsService) *common.ResponseMessage
	IsActiveProvider() bool
}

type SmsService struct {
	ActiveSmsService SmsServiceInterface
	Providers        []SmsServiceInterface
	From             string   `json:"from"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	Url              string   `json:"url"`
	Text             string   `json:"text"`
	Phones           []string `json:"phones"`
}

func (ss SmsService) addProviders() *SmsService {
	ss.Providers = append(ss.Providers, &Pacific{})
	return &ss
}

func (ss SmsService) SetActiveSmsService() (*SmsService, error) {
	providers := ss.addProviders()
	for _, service := range providers.Providers {
		if service.IsActiveProvider() {
			return &SmsService{
				ActiveSmsService: service,
			}, nil
		}
	}

	return nil, errors.New("no active sms service")
}
