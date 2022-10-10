package chainApi

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mock "github.com/psy2848048/go-antelope-sdk/utils/mock_server"
)

func TestRESTGetProducers(t *testing.T) {
	testDomains := []string{"http://localhost"}
	limit := 64

	mock.CreateAndActivateRestMockServer()
	defer mock.DeactivateMockServer()

	ret, err := RESTGetProducers(testDomains, uint(limit), "")

	assert.NoError(t, err)
	assert.Equal(t, len(ret.Rows), 64)
}
