package cyphertxn

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ReadTX(ds *DatabaseService, statement Statement) ([]*neo4j.Record, error) {
	session := ds.Driver.NewSession(ds.Ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: ds.Name,
	})
	defer session.Close(ds.Ctx)

	res, err := session.ExecuteRead(ds.Ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ds.Ctx, statement.Query, statement.Params)
		if err != nil {
			return nil, err
		}

		return result.Collect(ds.Ctx)
	})

	if err != nil {
		return []*neo4j.Record{}, err
	}

	return res.([]*neo4j.Record), nil
}

func ManagedTx(ds *DatabaseService, work neo4j.ManagedTransactionWork, session neo4j.SessionWithContext) (any, error) {
	result, err := session.ExecuteWrite(ds.Ctx, work)
	if err != nil {
		return nil, err
	}

	return result, nil
}
