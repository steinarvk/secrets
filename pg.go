package secrets

import (
	"errors"
	"fmt"
)

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (d *Postgres) AsURL() (string, error) {
	switch {
	case d.Host == "":
		return "", errors.New("missing host")
	case d.Port == 0:
		return "", errors.New("missing port")
	case d.Database == "":
		return "", errors.New("missing database")
	case d.User == "":
		return "", errors.New("missing user")
	case d.Password == "":
		return "", errors.New("missing password")
	default:
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.User, d.Password, d.Host, d.Port, d.Database), nil
	}
}
