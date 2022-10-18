package chainApi

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/psy2848048/go-antelope-sdk/types"
	"github.com/psy2848048/go-antelope-sdk/utils"
)

// implements `get_account` of Chain API
// Parameters:
//   - nodeDomains []string:
//   - req: RequestGetAccount{account: string}
func RESTGetAccount(nodeDomains []string, req *RequestGetAccount) (*ResponseGetAccount, error) {
	endpoint := "v1/chain/get_account"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetAccount, rest call")
		return nil, err
	}

	ret := &ResponseGetAccount{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetAccount, unmarshal")
		return nil, err
	}

	return ret, nil
}

// implements `get_block` of Chain API
// Parameters:
//   - nodeDomains []string:
//   - req: RequestGetBlock{block_num_or_id: string}
func RESTGetBlock(nodeDomains []string, req *RequestGetBlock) (*ResponseGetBlock, error) {
	endpoint := "v1/chain/get_block"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetBlock, rest call")
		return nil, err
	}

	ret := &ResponseGetBlock{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetBlock, unmarshal")
		return nil, err
	}

	return ret, nil
}

// implements `get_block_info` of Chain API
// Parameters:
//   - nodeDomains []string:
//   - req: RequestGetBlockInfo{block_num: int}
func RESTGetBlockInfo(nodeDomains []string, req *RequestGetBlockInfo) (*ResponseGetBlockInfo, error) {
	endpoint := "v1/chain/get_block"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetBlockInfo, rest call")
		return nil, err
	}

	ret := &ResponseGetBlockInfo{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetBlockInfo, unmarshal")
		return nil, err
	}

	return ret, nil
}

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

func RESTPushTransaction() {
	panic("not implemented")
}

func RESTPushTransactions() {
	panic("not implemented")
}

// implements `get_block_header_state` of Chain API
// Parameters:
//   - nodeDomains []string:
//   - req: RequestGetBlockHeaderState{block_num_or_id: string}
func RESTGetBlockHeaderState(nodeDomains []string, req *RequestGetBlockHeaderState) (*ResponseGetBlockHeaderState, error) {
	endpoint := "v1/chain/get_block_header_state"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetBlockHeaderState, rest call")
		return nil, err
	}

	ret := &ResponseGetBlockHeaderState{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetBlockHeaderState, unmarshal")
		return nil, err
	}

	return ret, nil
}

// implements `get_abi` of Chain API
// Parameters:
//   - nodeDomains []string:
//   - req: RequestGetAbi{account_name: string}
func RESTGetAbi(nodeDomains []string, req *RequestGetAbi) (*ResponseGetAbi, error) {
	endpoint := "v1/chain/get_abi"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetAbi, rest call")
		return nil, err
	}

	ret := &ResponseGetAbi{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetAbi, unmarshal")
		return nil, err
	}

	return ret, nil
}

// implements `get_currency_balance` of Chain API
// Parameters:
//   - nodeDomains []string:
//   - req: RequestGetCurrencyBalance{
//     code: string,          // account name of the token contract
//     account_name: string,  // account name of the token holder
//     symbol: string         // upper cased token symbol
//     }
func RESTGetCurrencyBalance(nodeDomains []string, req *RequestGetCurrencyBalance) (*ResponseGetCurrencyBalance, error) {
	endpoint := "v1/chain/get_currency_balance"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetCurrencyBalance, rest call")
		return nil, err
	}

	ret := &ResponseGetCurrencyBalance{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetCurrencyBalance, unmarshal")
		return nil, err
	}

	return ret, nil
}

func RESTGetCurrencyStats(nodeDomains []string, req *RequestGetCurrencyStats) (*ResponseGetCurrencyStats, error) {
	endpoint := "v1/chain/get_currency_stats"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetCurrencyStats, rest call")
		return nil, err
	}

	ret := &ResponseGetCurrencyStats{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetCurrencyStats, unmarshal")
		return nil, err
	}

	return ret, nil
}

func RESTGetRequiredKeys(nodeDomains []string, req *RequestGetRequiredKeys) (*ResponseGetRequiredKeys, error) {
	endpoint := "v1/chain/get_required_keys"

	// null check + force input empty array
	if req.Transaction.ContextFreeActions == nil {
		req.Transaction.ContextFreeActions = make([]types.UnitAction, 0)
	}

	if req.Transaction.TransactionExtensions == nil {
		req.Transaction.TransactionExtensions = make([]uint64, 0)
	}

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetRequiredKeys, rest call")
		return nil, err
	}

	ret := &ResponseGetRequiredKeys{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetRequiredKeys, unmarshal")
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

func RESTGetRawCodeAndABI(nodeDomains []string, req *RequestGetRawCodeAndABI) (*ResponseGetRawCodeAndABI, error) {
	endpoint := "v1/chain/get_raw_code_and_abi"

	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	byteResp, err := utils.RESTCallWithJson(nodeDomains, endpoint, utils.POST, byteReq)
	if err != nil {
		err = errors.Wrap(err, "RESTGetRawCodeAndABI, rest call")
		return nil, err
	}

	ret := &ResponseGetRawCodeAndABI{}
	err = json.Unmarshal(byteResp, ret)
	if err != nil {
		err = errors.Wrap(err, "RESTGetRawCodeAndABI, unmarshal")
		return nil, err
	}

	return ret, nil
}
