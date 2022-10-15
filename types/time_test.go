package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeUnmarshal(t *testing.T) {
	givenTime := []byte(`"2022-10-14T12:01:36.500"`)
	convTime := &Time{}

	err := json.Unmarshal(givenTime, convTime)
	fmt.Println(convTime)
	assert.NoError(t, err)
}
