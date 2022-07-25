package rest

type Config struct {
	Host      string
	Port      string
	TLSEnable bool
	CertFile  string
	KeyFile   string
	BotToken  string
}

func (c Config) Address() string {
	return c.Host + ":" + c.Port
}
