package config

import (
	"github.com/zeromicro/go-zero/rest"
	"portal/pkg/model"
)

type Config struct {
	rest.RestConf
	Model model.Config
}
