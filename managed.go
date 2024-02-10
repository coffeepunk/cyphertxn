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

func Transactions(ds *DatabaseService, statements ...Statement) (any, error) {
	session := ds.Driver.NewSession(ds.Ctx, neo4j.SessionConfig{
		DatabaseName: ds.Name,
	})
	defer session.Close(ds.Ctx)

	result, err := session.ExecuteWrite(ds.Ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		var results []*neo4j.Record
		for _, s := range statements {
			result, runErr := tx.Run(ds.Ctx, s.Query, s.Params)
			if runErr != nil {
				return nil, runErr
			}

			records, errRes := result.Collect(ds.Ctx)
			if errRes != nil {
				return []*neo4j.Record{}, errRes
			}

			results = append(results, records...)
		}

		return results, nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}
