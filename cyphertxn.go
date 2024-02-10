package cyphertxn

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DatabaseService struct {
	Ctx    context.Context
	Driver neo4j.DriverWithContext
	Name   string
}

type Statement struct {
	Query  string
	Params map[string]interface{}
}

func NewDBService(ctx context.Context, driver neo4j.DriverWithContext, name string) *DatabaseService {
	return &DatabaseService{
		Ctx:    ctx,
		Driver: driver,
		Name:   name,
	}
}
