package cyphertxn

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

type MockDriver struct {
	bookmarkManager neo4j.BookmarkManager
	url             url.URL
	session         neo4j.SessionWithContext
	serverInfo      neo4j.ServerInfo
}

func (md *MockDriver) ExecuteQueryBookmarkManager() neo4j.BookmarkManager {
	return md.bookmarkManager
}

func (md *MockDriver) Target() url.URL {
	return md.url
}

func (md *MockDriver) NewSession(ctx context.Context, config neo4j.SessionConfig) neo4j.SessionWithContext {
	return md.session
}

func (md *MockDriver) VerifyConnectivity(ctx context.Context) error {
	return nil
}

func (md *MockDriver) VerifyAuthentication(ctx context.Context, auth *neo4j.AuthToken) error {
	return nil
}

func (md *MockDriver) Close(ctx context.Context) error {
	return nil
}

func (md *MockDriver) IsEncrypted() bool {
	return false
}

func (md *MockDriver) GetServerInfo(ctx context.Context) (neo4j.ServerInfo, error) {
	return md.serverInfo, nil
}

func (md *MockDriver) SessionWithContext(ctx context.Context, access neo4j.AccessMode) (neo4j.SessionWithContext, error) {
	return md.session, nil
}

func TestNewDBService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		ctx    context.Context
		driver neo4j.DriverWithContext
		dbName string
	}{
		{
			name:   "Valid",
			ctx:    context.Background(),
			driver: &MockDriver{},
			dbName: "TestDB",
		},
		{
			name:   "WithCancelContext",
			ctx:    context.WithValue(context.Background(), "key", "value"),
			driver: &MockDriver{},
			dbName: "TestDB",
		},
		{
			name:   "WithEmptyName",
			ctx:    context.Background(),
			driver: &MockDriver{},
			dbName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewDBService(tt.ctx, tt.driver, tt.dbName)

			assert.NotNil(t, result)
			assert.Equal(t, tt.ctx, result.Ctx)
			assert.Equal(t, tt.driver, result.Driver)
			assert.Equal(t, tt.dbName, result.Name)
		})
	}
}
