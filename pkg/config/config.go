package config

type Config struct {
	Debug                        bool
	RefreshRateSeconds           int
	PushEnabled                  bool
	PushIntervalSeconds          int
	PushPath                     string
	PushUrl                      string
	Token                        string
	Endpoint                     string
	MetricsFile                  string
	CollectorEnabled             bool
	CollectorRetentionPeriodDays int
	BindAddr                     string
}

func NewDefaultConfig() *Config {
	return &Config{
		Debug:                        false,
		RefreshRateSeconds:           5,
		PushEnabled:                  true,
		PushIntervalSeconds:          60,
		PushPath:                     "/api/v1/sentinel/push",
		PushUrl:                      "",
		Token:                        "",
		Endpoint:                     "",
		MetricsFile:                  "/app/db/metrics.sqlite",
		CollectorEnabled:             false,
		CollectorRetentionPeriodDays: 7,
		BindAddr:                     ":8888",
	}
}
