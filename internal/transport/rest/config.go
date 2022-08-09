package rest

import "net"

type Config struct {
	Host      string
	Port      string
	TLSEnable bool
	CertFile  string
	KeyFile   string
}

func (c Config) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}
