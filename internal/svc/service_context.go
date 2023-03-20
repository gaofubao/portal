package svc

import (
	"portal/internal/config"
	"portal/pkg/model/chatgpt"
)

type ServiceContext struct {
	Config config.Config
	Model  *chatgpt.GptClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  chatgpt.New(c.Model),
	}
}
