package config

type App struct {
	AppMode     string `envconfig:"APP_MODE"`
	AppLocation string `envconfig:"APP_LOCATION"`

	BotToken  string `envconfig:"TELEGRAM_BOT_TOKEN"`
	BotTimout int    `envconfig:"TELEGRAM_BOT_TIMEOUT"`
	BotDebug  bool   `envconfig:"TELEGRAM_BOT_DEBUG"`
}

func (a *App) IsProduction() bool {
	return a.AppMode == "production"
}
