package mockServer

import (
	"net/http"
	"os"
	"strings"

	"github.com/jarcoal/httpmock"
	"github.com/psy2848048/go-antelope-sdk/utils"
)

func CreateAndActivateRestMockServer() {
	httpmock.Activate()

	httpmock.RegisterResponder(
		utils.POST, "http://localhost/v1/chain/get_account",
		httpmock.NewStringResponder(http.StatusOK, openFile("chain_get_account.json")),
	)

	httpmock.RegisterResponder(
		utils.POST, "http://localhost/v1/chain/get_block",
		httpmock.NewStringResponder(http.StatusOK, openFile("chain_get_block.json")),
	)

	httpmock.RegisterResponder(
		utils.POST, "http://localhost/v1/chain/get_block_info",
		httpmock.NewStringResponder(http.StatusOK, openFile("chain_get_block_info.json")),
	)

	httpmock.RegisterResponder(
		utils.GET, "http://localhost/v1/chain/get_info",
		httpmock.NewStringResponder(http.StatusOK, openFile("chain_get_info.json")),
	)

	httpmock.RegisterResponder(
		utils.POST, "http://localhost/v1/chain/get_block_header_state",
		httpmock.NewStringResponder(http.StatusOK, openFile("chain_get_block_header_state.json")),
	)

	httpmock.RegisterResponder(
		utils.POST, "http://localhost/v1/chain/get_abi",
		httpmock.NewStringResponder(http.StatusOK, openFile("chain_get_abi.json")),
	)

	httpmock.RegisterResponder(
		utils.POST, "http://localhost/v1/chain/get_producers",
		httpmock.NewStringResponder(http.StatusOK, openFile("chain_get_producers.json")),
	)
}

func openFile(filename string) string {
	path := []string{"../../utils/mock_server", filename}
	body, err := os.ReadFile(strings.Join(path, "/"))
	if err != nil {
		panic(err)
	}

	return string(body)
}

func DeactivateMockServer() {
	httpmock.DeactivateAndReset()
}
