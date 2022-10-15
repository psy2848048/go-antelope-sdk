package chainApi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	mock "github.com/psy2848048/go-antelope-sdk/utils/mock_server"
)

func TestRESTGetProducers(t *testing.T) {
	testDomains := []string{"http://localhost"}
	limit := 64

	ret, err := RESTGetProducers(testDomains, uint(limit), "")

	assert.NoError(t, err)
	assert.Equal(t, len(ret.Rows), 64)
}

func TestRESTGetInfo(t *testing.T) {
	testDomains := []string{"http://localhost"}

	ret, err := RESTGetInfo(testDomains)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetAccount(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetAccount{
		AccountName: "teamgreymass",
	}

	ret, err := RESTGetAccount(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetBlock(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetBlock{
		BlockNumOrId: "273283700",
	}

	ret, err := RESTGetBlock(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetBlockInfo(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetBlockInfo{
		BlockNum: "273283700",
	}

	ret, err := RESTGetBlockInfo(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setUp() {
	mock.CreateAndActivateRestMockServer()
}

func tearDown() {
	mock.DeactivateMockServer()
}
