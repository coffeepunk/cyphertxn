package cyphertxn

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

type TargetCredentials struct {
	URI      string
	Username string
	Password string
	Realm    string
}

func BasicAuth(tc TargetCredentials) (neo4j.DriverWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(tc.URI, neo4j.BasicAuth(tc.Username, tc.Password, tc.Realm))

	if err != nil {
		log.Fatal("Invalid DB config:", err)
		return driver, err
	}

	return driver, nil
}
