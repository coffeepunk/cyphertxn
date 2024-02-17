package cyphertxn

import (
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"testing"
)

type mockDriverStruct struct {
	neo4j.DriverWithContext
}

func (mds mockDriverStruct) ExecuteQuery(ctx context.Context, query string, params map[string]interface{}, resultTransformer neo4j.ResultTransformer[any], settings ...neo4j.ExecuteQueryConfigurationOption) (*neo4j.EagerResult, error) {
	if query == "" {
		return nil, errors.New("query is empty")
	}
	return &neo4j.EagerResult{}, nil
}

func TestQueryWrite(t *testing.T) {
	testCases := []struct {
		name      string
		statement Statement
	}{
		{
			name: "Normal Input",
			statement: Statement{
				Query:  "MATCH (n:Person) RETURN n.name as name",
				Params: map[string]interface{}{},
			},
		},
		{
			name: "Empty Query",
			statement: Statement{
				Query:  "",
				Params: map[string]interface{}{},
			},
		},
		{
			name: "Query with Params",
			statement: Statement{
				Query:  "MATCH (n:Person {name: $name}) RETURN n.age as age",
				Params: map[string]interface{}{"name": "Alice"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ds := DatabaseService{
				Ctx:    context.Background(),
				Driver: &MockDriver{},
				Name:   "test",
			}
			_, err := QueryWrite(ds, tc.statement)
			if err != nil && tc.statement.Query != "" {
				t.Fatalf("failed test "+tc.name+" with error: %v", err)
			}
			if tc.statement.Query == "" && err == nil {
				t.Fatalf("failed test " + tc.name + ". Expected an error but got no error.")
			}
		})
	}
}
