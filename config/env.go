package config

import (
	"app/libs/dotenv"
	"fmt"
	"os"
)

var Env = env{}

func init() {
	var de dotenv.DotEnv

	err := de.Load()

	if err != nil {
		fmt.Println("Cannot load .env file")
		os.Exit(-1)
	}

	Env = env{
		App: App{
			Stage:         de.GetString("APP_STAGE", "DEV", ""),
			Version:       de.GetString("APP_VERSION", "v1", ""),
			Address:       de.GetString("APP_ADDRESS", "0.0.0.0", ""),
			Protocol:      de.GetString("APP_PROTOCOL", "http", ""),
			Name:          de.GetString("APP_NAME", "LAN-ACCOUNTS", ""),
			Port:          de.GetUint16("APP_PORT", "3000", ""),
			ServerTimeout: de.GetUint16("APP_SERVER_TIMEOUT", "4", ""),
			ExitTimeout:   de.GetUint16("APP_EXIT_TIMEOUT", "1", ""),
		},
	}
}
