package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/steinarvk/secrets"
)

var (
	dbSecretsFilename = flag.String("db_secrets", "", "database secrets YAML file")

	noEnsureSSL = flag.Bool("nossl", false, "don't add SSL option")
)

func mainCore() error {
	dbSecrets := secrets.Postgres{}
	if err := secrets.FromYAML(*dbSecretsFilename, &dbSecrets); err != nil {
		return err
	}

	opts := &secrets.PostgresOptions{}
	if *noEnsureSSL {
		opts.WithoutSSL = true
	}

	url, err := dbSecrets.AsURLWithOptions(opts)
	if err != nil {
		return err
	}

	fmt.Print(url)

	return nil
}

func main() {
	flag.Parse()

	if *dbSecretsFilename == "" && len(flag.Args()) == 1 {
		*dbSecretsFilename = flag.Args()[0]
	}

	if err := mainCore(); err != nil {
		log.Printf("fatal: %v", err)
	}
}
