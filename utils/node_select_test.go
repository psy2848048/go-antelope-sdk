package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var givenNodes = map[string]UnitNodeStatus{
	"https://api.eoseoul.io": {
		Success: 1,
		Failure: 0,
	},

	"https://mainnet.genereos.io": {
		Success: 2,
		Failure: 1,
	},

	// "https://api1-eos.nodeone.network:8344": {
	// 	Success: 0,
	// 	Failure: 0,
	// },

	// "https://domain2.com": {
	// 	Success: 0,
	// 	Failure: 0,
	// },
}

func TestSuiteNodeStatsCRUD(t *testing.T) {
	// 1. MakeStat
	{
		MakeStat(givenNodes)
		currNodes := GetStat()
		assert.Equal(t, givenNodes, currNodes)
	}

	// 2. Pick random nodes
	{
		nodeSet := []string{}
		NodeRESTResponseStatus.Range(func(key, _ interface{}) bool {
			nodeSet = append(nodeSet, key.(string))
			return true
		})

		domain := PickRandomNodeFromNodeSets(nodeSet)
		assert.NotNil(t, domain)
	}

	// 3. Increase success (existing)
	{
		IncreaseSuccess("https://api.eoseoul.io")

		val := givenNodes["https://api.eoseoul.io"]
		val.Success += 1
		givenNodes["https://api.eoseoul.io"] = val

		currNodes := GetStat()
		assert.Equal(t, givenNodes, currNodes)
	}

	// 4. Increase failure (existing)
	{
		IncreaseFailure("https://mainnet.genereos.io")

		val := givenNodes["https://mainnet.genereos.io"]
		val.Failure += 1
		givenNodes["https://mainnet.genereos.io"] = val

		currNodes := GetStat()
		assert.Equal(t, givenNodes, currNodes)
	}

	// 5. Increase success (new)
	{
		IncreaseSuccess("https://api1-eos.nodeone.network:8344")

		givenNodes["https://api1-eos.nodeone.network:8344"] = UnitNodeStatus{
			Success: 1,
		}

		currNodes := GetStat()
		assert.Equal(t, givenNodes, currNodes)
	}

	// 6. Increase failure (new)
	{
		IncreaseFailure("https://domain2.com")

		givenNodes["https://domain2.com"] = UnitNodeStatus{
			Failure: 1,
		}

		currNodes := GetStat()
		assert.Equal(t, givenNodes, currNodes)
	}

	// 7. Reset
	{
		Reset()

		counter := 0
		NodeRESTResponseStatus.Range(func(key, _ interface{}) bool {
			counter += 1
			return true
		})

		assert.Equal(t, 0, counter)
	}
}
