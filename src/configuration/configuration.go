package configuration

import (
	"github.com/kelseyhightower/envconfig"
)

// Config
type Config struct {
	DSN       string `envconfig:"DSN" required:"true"`
	Port      string `envconfig:"PORT" required:"true"`
	Env       string `envconfig:"ENV" required:"true"`
	SentryDSN string `envconfig:"SENTRY_DSN" required:"false"`
}

var globalConfig Config

// Load は環境変数を globalConfig に読み込む。
func Load() {
	prefix := ""
	envconfig.MustProcess(prefix, &globalConfig)
}

// Get は Config を取得する。
func Get() Config {
	return globalConfig
}
