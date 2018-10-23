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

type PostgresOptions struct {
	WithoutSSL bool
}

func (d *Postgres) check() error {
	switch {
	case d.Host == "":
		return errors.New("missing host")
	case d.Port == 0:
		return errors.New("missing port")
	case d.Database == "":
		return errors.New("missing database")
	case d.User == "":
		return errors.New("missing user")
	case d.Password == "":
		return errors.New("missing password")
	default:
		return nil
	}
}

func (d *Postgres) AsURLWithOptions(opts *PostgresOptions) (string, error) {
	if opts == nil {
		opts = &PostgresOptions{}
	}

	optionstring := ""
	if !opts.WithoutSSL {
		optionstring = "?ssl=true"
	}

	if err := d.check(); err != nil {
		return "", err
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s%s", d.User, d.Password, d.Host, d.Port, d.Database, optionstring), nil
}

func (d *Postgres) AsURL() (string, error) {
	return d.AsURLWithOptions(nil)
}

func (d *Postgres) AsOptionString() (string, error) {
	if err := d.check(); err != nil {
		return "", err
	}
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=require", d.Host, d.Port, d.Database, d.User, d.Password), nil
}
