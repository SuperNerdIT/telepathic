package configs

type Configs struct {
	Port string
	Host string
}

func MakeConfigs() *Configs {
	// we could trigger something to read envs or command line arguments here
	// so this could be dynamic
	return &Configs{
		Port: "8080",
		Host: "localhost",
	}
}
