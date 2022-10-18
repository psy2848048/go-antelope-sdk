package chainApi

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/psy2848048/go-antelope-sdk/types"
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

func TestRESTGetBlockHeaderState(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetBlockHeaderState{
		BlockNumOrId: "273457972",
	}

	ret, err := RESTGetBlockHeaderState(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetABI(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetAbi{
		AccountName: "tippedtipped",
	}

	ret, err := RESTGetAbi(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetCurrencyBalance(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetCurrencyBalance{
		Code:    "eosiochaince",
		Account: "haydinrvgage",
		Symbol:  "CET",
	}

	ret, err := RESTGetCurrencyBalance(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetCurrencyStats(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetCurrencyStats{
		Code:   "eosiochaince",
		Symbol: "CET",
	}

	ret, err := RESTGetCurrencyStats(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetRequiredKeys(t *testing.T) {
	testDomains := []string{"http://localhost"}
	currTime, _ := types.NewTimeFromString("2022-10-17T13:59:48")

	req := &RequestGetRequiredKeys{
		Transaction: types.TransactionInfo{
			Expiration:       currTime,
			RefBlockNum:      64629,
			RefBlockPrefix:   1239991166,
			MaxNetUsageWords: 0,
			MaxCpuUsageMs:    0,
			DelaySec:         0,

			Actions: []types.UnitAction{
				{
					Account: "everipediaiq",
					Name:    "transfer",
					Authorization: []types.UnitAuthorization{
						{
							Actor:      "haydinrvgage",
							Permission: "owner",
						},
					},
					Data: "a09861fb4e97bc69002ed6d99daa3155102700000000000003495100000000002431346534373536622d633630332d346537382d383966322d303663333364646430376534",
				},
			},
		},
		AvailableKeys: []string{
			"EOS5LYTegVed738P24stzRNG9BFdt5hxWHUo2VrnPE6LXDDcMa4dD",
		},
	}

	ret, err := RESTGetRequiredKeys(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)

	byteReq, _ := json.MarshalIndent(req, "", "    ")
	fmt.Println(string(byteReq))
}

func TestRESTGetRawCodeAndABI(t *testing.T) {
	testDomains := []string{"http://localhost"}
	req := &RequestGetRawCodeAndABI{
		AccountName: "tippedtipped",
	}

	ret, err := RESTGetRawCodeAndABI(testDomains, req)

	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestRESTGetTableByScope(t *testing.T) {
	testDomains := []string{"http://localhost"}

	var limit uint64 = 10
	req := &RequestGetTableByScope{
		Code:      "tippedtipped",
		Table:     "accounts",
		Limit:     &limit,
		Reverse:   false,
		ShowPayer: false,
	}

	ret, err := RESTGetTableByScope(testDomains, req)

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
