package chainApi

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/psy2848048/go-antelope-sdk/utils"
)

// implements `get_info` of Chain API
// Parameters:
//   - nodeDomains []string:
func RESTGetInfo(nodeDomains []string) (*ResponseGetInfo, error) {
	endpoint := "v1/chain/get_info"

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.GET, nil)
	if err != nil {
		err = errors.Wrap(err, "RESTGetInfo, rest call")
		return nil, err
	}

	ret := &ResponseGetInfo{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetInfo, unmarshal")
		return nil, err
	}

	return ret, nil
}

// implements `get_producers` of Chain API
// Parameters:
//   - nodeDomains []string:
//   - limit uint: set how many producer information are listed. If you input 0, the SDK thinks as nil
//   - lowerBound string: query the result from the parameter. The result includes the given param. If you input empty string, the SDK judges as nil
func RESTGetProducers(nodeDomains []string, limit uint, lowerBound string) (*ResponseGetProducers, error) {
	endpoint := "v1/chain/get_producers"
	unmarshaledParam := &RequestGetProducers{
		JSON: true,
	}

	if limit > 0 {
		inputLimit := fmt.Sprintf("%d", limit)
		unmarshaledParam.Limit = &inputLimit
	}

	if lowerBound != "" {
		lowerBoundParam := lowerBound // just copy
		unmarshaledParam.LowerBound = &lowerBoundParam
	}

	param, err := json.Marshal(unmarshaledParam)
	if err != nil {
		err = errors.Wrap(err, "go-antelope-sdk, RESTGetProducers, json marshal")
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, param)
	if err != nil {
		err = errors.Wrap(err, "go-antelope-sdk, RESTGetProducers, rest call")
		return nil, err
	}

	ret := &ResponseGetProducers{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "go-antelope-sdk, RESTGetProducers, unmarshal after REST call")
		return nil, err
	}

	return ret, nil
}
