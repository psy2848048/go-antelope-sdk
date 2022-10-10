package nodeLists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopRankedBPNodeList(t *testing.T) {
	nodeList, err := GetTopRankedBPNodeList(10)
	assert.NoError(t, err)
	assert.NotNil(t, nodeList)

	fmt.Println("List of API nodes from the right URL and the right-formed bp.json:")
	for _, addr := range nodeList {
		fmt.Println(addr)
	}
}
