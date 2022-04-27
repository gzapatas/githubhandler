package config

type (
	// Env variables de entorno de la aplicacion
	env struct {
		App App
	}

	App struct {
		Stage         string
		Version       string
		Address       string
		Port          uint16
		Protocol      string
		Name          string
		ServerTimeout uint16
		ExitTimeout   uint16
	}
)
