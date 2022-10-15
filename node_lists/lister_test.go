package nodeLists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopRankedBPNodeList(t *testing.T) {
	nodeList, err := GetTopRankedBPNodeList(100)
	assert.NoError(t, err)
	assert.NotNil(t, nodeList)

	filteredList, err := FilterByAPINodeCallable(nodeList)
	assert.NoError(t, err)
	assert.NotNil(t, filteredList)

	fmt.Println("These are reachable node list!:")
	for _, addr := range filteredList {
		fmt.Println(addr)
	}
}
