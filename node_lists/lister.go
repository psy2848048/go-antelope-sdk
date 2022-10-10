package nodeLists

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"

	chainApi "github.com/psy2848048/go-antelope-sdk/rest_wrapper/chain_api"
	"github.com/psy2848048/go-antelope-sdk/utils"
)

// trusted Korean BPs
// will use as seed for fetching API endpoints of top 64 BPs
var initTrustedNodes = []string{
	"https://api.eoseoul.io",
	"https://api1-eos.nodeone.network:8344",
}

func GetTopRankedBPNodeList(numberOfNodes uint) ([]string, error) {
	nodeLists, err := chainApi.RESTGetProducers(initTrustedNodes, numberOfNodes, "")
	if err != nil {
		err = errors.Wrap(err, "GetTopRankedBPNodeList, RESTGetProducers")
		return nil, err
	}

	// make "<url>/bp.json"
	nodeBpJsonLists := []string{}
	for _, val := range nodeLists.Rows {
		if val.Url != "" {
			nodeBpJsonLists = append(nodeBpJsonLists, val.Url)
		}
	}

	apiNodes := []string{}
	var wg sync.WaitGroup

	for _, unitUrl := range nodeBpJsonLists {
		wg.Add(1)

		go func(unitUrl string) {
			defer wg.Done()

			bpJsonInfoByte, err := utils.RESTCallWithJson([]string{unitUrl}, "bp.json", utils.GET, nil)
			if err != nil {
				wrapMsg := fmt.Sprintf("GetTopRankedBPNodeList, receiving each %s/bp.json info", unitUrl)
				err = errors.Wrap(err, wrapMsg)
				logger.Println(err)

				return
			}

			bpJsonInfo := &UnitBPInfo{}
			err = json.Unmarshal(bpJsonInfoByte, bpJsonInfo)
			if err != nil {
				wrapMsg := fmt.Sprintf("GetTopRankedBPNodeList, unmarshaling each %s/bp.json info", unitUrl)
				err = errors.Wrap(err, wrapMsg)
				logger.Println(err)

				return
			}

			for _, val := range bpJsonInfo.Nodes {
				if val.ApiEndpoint != nil && *val.ApiEndpoint != "" {
					unit := *val.ApiEndpoint

					if strings.HasPrefix(unit, "http://") {
						unit = strings.Replace(unit, "http://", "https://", 1)
					}

					apiNodes = append(apiNodes, unit)
				}
			}
		}(unitUrl)
	}

	wg.Wait()

	apiNodes = utils.RemoveDuplicateValues(apiNodes)

	return apiNodes, nil
}
