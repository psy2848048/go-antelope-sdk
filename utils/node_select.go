package utils

import (
	"math/rand"
	"sync"
	"time"
)

var NodeRESTResponseStatus sync.Map

type UnitNodeStatus struct {
	Success int64 `json:"success"`
	Failure int64 `json:"failure"`
}

func PickRandomNodeFromNodeSets(nodeDomains []string) string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(nodeDomains)

	return nodeDomains[n]
}

func IncreaseSuccess(domain string) {
	if val, ok := NodeRESTResponseStatus.Load(domain); ok {
		castedVal := val.(UnitNodeStatus)
		castedVal.Success += 1

		NodeRESTResponseStatus.Store(domain, castedVal)
	} else {
		newVal := UnitNodeStatus{
			Success: 1,
			Failure: 0,
		}

		NodeRESTResponseStatus.Store(domain, newVal)
	}
}

func IncreaseFailure(domain string) {
	if val, ok := NodeRESTResponseStatus.Load(domain); ok {
		castedVal := val.(UnitNodeStatus)
		castedVal.Failure += 1

		NodeRESTResponseStatus.Store(domain, castedVal)
	} else {
		newVal := UnitNodeStatus{
			Success: 0,
			Failure: 1,
		}

		NodeRESTResponseStatus.Store(domain, newVal)
	}
}

func GetStat() map[string]UnitNodeStatus {
	ret := make(map[string]UnitNodeStatus)

	NodeRESTResponseStatus.Range(func(k, v interface{}) bool {
		ret[k.(string)] = v.(UnitNodeStatus)
		return true
	})

	return ret
}

func MakeStat(prevData map[string]UnitNodeStatus) {
	for k, uns := range prevData {
		NodeRESTResponseStatus.Store(k, uns)
	}
}

func Reset() {
	NodeRESTResponseStatus.Range(func(k, _ interface{}) bool {
		NodeRESTResponseStatus.Delete(k)
		return true
	})
}
