package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/steinarvk/secrets"
)

var (
	dbSecretsFilename = flag.String("db_secrets", "", "database secrets YAML file")
)

func mainCore() error {
	dbSecrets := secrets.Postgres{}
	if err := secrets.FromYAML(*dbSecretsFilename, &dbSecrets); err != nil {
		return err
	}

	url, err := dbSecrets.AsURL()
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
