package chainApi

import (
	"encoding/json"

	"github.com/psy2848048/go-antelope-sdk/types"
)

type RequestGetAccount struct {
	AccountName string `json:"account_name"`
}

type ResponseGetAccount struct {
	AccountName            string                  `json:"account_name"`
	HeadBlockNum           uint64                  `json:"head_block_num"`
	HeadBlockTime          types.Time              `json:"head_block_time"`
	LastCodeUpdate         types.Time              `json:"last_code_update"`
	Created                types.Time              `json:"created"`
	RefundRequest          *RefundRequest          `json:"refund_request,omitempty"`
	RamQuota               uint64                  `json:"ram_quota"`
	NetLimit               *Limit                  `json:"net_limit,omitempty"`
	CpuLimit               *Limit                  `json:"cpu_limit"`
	TotalResources         *Resources              `json:"total_resources"`
	CoreLiquidBalance      *types.Coin             `json:"core_liquid_balance,omitempty"`
	SelfDelegatedBandwidth *SelfDelegatedBandwidth `json:"self_delegated_bandwidth,omitempty"`
	NetWeight              uint64                  `json:"net_weight"`
	CpuWeight              types.Int               `json:"cpu_weight"`
	RamUsage               uint64                  `json:"ram_usage"`
	Privileged             bool                    `json:"privileged"`
	VoterInfo              *VoterInfo              `json:"voter_info,omitempty"`
}

type RequestGetBlock struct {
	BlockNumOrId string `json:"block_num_or_id"`
}

type ResponseGetBlock struct {
	Timestamp         types.Time    `json:"timestamp"`
	Producer          string        `json:"producer"`
	Confirmed         uint64        `json:"confirmed"`
	Previous          string        `json:"previous"`
	TransactionMroot  string        `json:"transaction_mroot"`
	ActionMroot       string        `json:"action_mroot"`
	ScheduleVersion   uint64        `json:"schedule_version"`
	NewProducers      *NewProducers `json:"new_producers,omitempty"`
	ProducerSignature string        `json:"producer_signature"`
	Id                string        `json:"id"`
	BlockNum          uint64        `json:"block_num"`
	RefBlockPrefix    uint64        `json:"ref_block_prefix"`

	Transactions []Tx `json:"transactions"`
}

type RequestGetBlockInfo struct {
	BlockNum string `json:"block_num"`
}

type ResponseGetBlockInfo struct {
	Timestamp         types.Time `json:"timestamp"`
	Producer          string     `json:"producer"`
	Confirmed         uint64     `json:"confirmed"`
	Previous          string     `json:"previous"`
	TransactionMroot  string     `json:"transaction_mroot"`
	ActionMroot       string     `json:"action_mroot"`
	ScheduleVersion   uint64     `json:"schedule_version"`
	ProducerSignature string     `json:"producer_signature"`
	Id                string     `json:"id"`
	BlockNum          uint64     `json:"block_num,omitempty"`
	RefBlockPrefix    uint64     `json:"ref_block_prefix,omitempty"`
}

type ResponseGetInfo struct {
	ServerVersion             string     `json:"server_version"`
	ChainId                   string     `json:"chain_id"`
	HeadBlockNum              uint64     `json:"head_block_num"`
	LastIrreversibleBlockNum  uint64     `json:"last_irreversible_block_num"`
	LastIrreversibleBlockId   string     `json:"last_irreversible_block_id"`
	HeadBlockId               string     `json:"head_block_id"`
	HeadBlockTime             types.Time `json:"head_block_time"` // : "2022-10-10T13:34:03.500",
	HeadBlockProducer         string     `json:"head_block_producer"`
	VirtualBlockCpuLimit      uint64     `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit      uint64     `json:"virtual_block_net_limit"`
	BlockCpuLimit             uint64     `json:"block_cpu_limit"`
	BlockNetLimit             uint64     `json:"block_net_limit"`
	ServerVersionString       string     `json:"server_version_string"`
	ForkDbHeadBlockNum        uint64     `json:"fork_db_head_block_num"`
	ForkDbHeadBlockId         string     `json:"fork_db_head_block_id"`
	ServerFullVersionString   string     `json:"server_full_version_string"`
	TotalCpuWeight            string     `json:"total_cpu_weight"`
	TotalNetWeight            string     `json:"total_net_weight"`
	EarliestAvailableBlockNum uint64     `json:"earliest_available_block_num"`
	LastIrreversibleBlockTime types.Time `json:"last_irreversible_block_time"`
}

type RequestGetAbi struct {
	AccountName string `json:"account_name"`
}

type ResponseGetAbi struct {
	AccountName string `json:"account_name"`
	ABI         *ABI   `json:"abi,omitempty"`
}

type RequestGetProducers struct {
	Limit      *string `json:"limit,omitempty"`
	LowerBound *string `json:"lower_bound,omitempty"`
	JSON       bool    `json:"json"`
}

type ResponseGetProducers struct {
	Rows                    []UnitProducerInfo `json:"rows"`
	TotalProducerVoteWeight string             `json:"total_producer_vote_weight"`
	More                    string             `json:"more"`
}

type RequestGetBlockHeaderState struct {
	BlockNumOrId string `json:"block_num_or_id"`
}

type ResponseGetBlockHeaderState struct {
	Id                                     string                      `json:"id"`
	Header                                 *ResponseGetBlockInfo       `json:"header"`
	BlockNum                               uint64                      `json:"block_num"`
	DposProposedIrreversibleBlockNum       uint64                      `json:"dpos_proposed_irreversible_blocknum"`
	DposIrreversibleBlockNum               uint64                      `json:"dpos_irreversible_blocknum"`
	ActiveSchedule                         *ActiveSchedule             `json:"active_schedule"`
	BlockrootMerkle                        interface{}                 `json:"blockroot_merkle"`
	ProducerToLastProduced                 []UnitProducerInHeaderState `json:"producer_to_last_produced"`
	ProducerToLastImpliedIrreversibleBlock []UnitProducerInHeaderState `json:"producer_to_last_implied_irb"`
	ValidBlockSigningAuthority             *UnitAuthority              `json:"valid_block_signing_authority,omitempty"`
	ConfirmCount                           []uint64                    `json:"confirm_count"`
	PendingSchedule                        *PendingSchedule            `json:"pending_schedule,omitempty"`
	ActivatedProtocolFeatures              *ProtocolFeatures           `json:"activated_protocol_features,omitempty"`
	AdditionalSignatures                   []interface{}               `json:"additional_signatures"`
}

type RequestGetCurrencyBalance struct {
	Code    string `json:"code"`
	Account string `json:"account"`
	Symbol  string `json:"symbol"`
}

type ResponseGetCurrencyBalance struct {
	Amount types.Coin `json:"amount"`
}

func (r *ResponseGetCurrencyBalance) UnmarshalJSON(b []byte) error {
	interimRet := []interface{}{}

	err := json.Unmarshal(b, &interimRet)
	if err != nil {
		return err
	}

	bytePartial, err := json.Marshal(interimRet[0])
	if err != nil {
		return err
	}

	ret := &types.Coin{}
	err = json.Unmarshal(bytePartial, ret)
	if err != nil {
		return err
	}

	r.Amount = *ret

	return nil
}

type RequestGetCurrencyStats struct {
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
}

type ResponseGetCurrencyStats map[string]UnitCurrencyStats

type RequestGetRequiredKeys struct {
	Transaction   types.TransactionInfo `json:"transaction"`
	AvailableKeys []string              `json:"available_keys"`
}

type ResponseGetRequiredKeys struct {
	RequiredKeys []string `json:"required_keys"`
}

type RequestGetRawCodeAndABI struct {
	AccountName string `json:"account_name"`
}

type ResponseGetRawCodeAndABI struct {
	AccountName string `json:"account_name"`
	WASM        string `json:"wasm"`
	ABI         string `json:"abi"`
}

type RequestGetTableByScope struct {
	Code       string  `json:"code"`
	Table      string  `json:"table"`
	LowerBound *string `json:"lower_bound,omitempty"`
	UpperBound *string `json:"upper_bound,omitempty"`
	Limit      *uint64 `json:"limit"`
	Reverse    bool    `json:"reverse"`
	ShowPayer  bool    `json:"show_payer"`
}

type ResponseGetTableByScope struct {
	Rows []AggregatedTableRow `json:"rows"`
	More string               `json:"more"`
}

// Subtypes

type RefundRequest struct {
	Owner       string     `json:"owner"`
	RequestTime types.Time `json:"request_time"`
	NetAmount   uint64     `json:"net_amount"`
	CpuAmount   uint64     `json:"cpu_amount"`
}

type Limit struct {
	Max       uint64 `json:"max"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
}

type Resources struct {
	Owner     string     `json:"owner"`
	RamBytes  uint64     `json:"ram_bytes"`
	NetWeight types.Coin `json:"net_weight"`
	CpuWeight types.Coin `json:"cpu_weight"`
}

type SelfDelegatedBandwidth struct {
	From      string     `json:"from"`
	To        string     `json:"to"`
	NetWeight types.Coin `json:"net_weight"`
	CpuWeight types.Coin `json:"cpu_weight"`
}

type Permission struct {
	Parent        string         `json:"parent"`
	PermName      string         `json:"perm_name"`
	RequiredAuth  *UnitAuth      `json:"required_auth,omitempty"`
	LinkedActions []LinkedAction `json:"linked_actions,omitempty"`
}

type UnitAuth struct {
	Waits     []Wait    `json:"waits"`
	Keys      []AuthKey `json:"keys"`
	Threshold int       `json:"threshold"`
	Accounts  []Account `json:"accounts"`
}

type Wait struct {
	WaitSec uint   `json:"wait_sec"`
	Weight  uint64 `json:"weight"`
}

type AuthKey struct {
	Key    string `json:"key"`
	Weight uint64 `json:"weight"`
}

type Account struct {
	// TODO: To be modified
	Weight     uint64             `json:"weight"`
	Permission *UnitAuthorization `json:"permission,omitempty"`
}

type LinkedAction struct {
	Account string `json:"account"`
	Action  string `json:"action"`
}

type VoterInfo struct {
	Owner             string        `json:"owner"`
	Proxy             string        `json:"proxy"`
	Producers         []string      `json:"producers"`
	Staked            uint64        `json:"staked"`
	LastVoteWeight    types.Decimal `json:"last_vote_weight"`
	ProxiedVoteWeight types.Decimal `json:"proxied_vote_weight"`
	IsProxy           int           `json:"is_proxy"`
	Flags1            int           `json:"flags1"`
	Reserved2         int           `json:"reserved2"`
	Reserved3         types.Coin    `json:"reserved3"`
}

type NewProducers struct {
	Version   string            `json:"version"`
	Producers []UnitNewProducer `json:"producers"`
}

type UnitNewProducer struct {
	ProducerName    string `json:"producer_name"`
	BlockSigningKey string `json:"block_signing_key"`
}

type Tx struct {
	Status        string `json:"status"`
	CpuUsageUs    uint64 `json:"cpu_usage_us"`
	NetUsageWords uint64 `json:"net_usage_words"`
	Trx           *Trx   `json:"trx"`
}

type Trx struct {
	Id                    string           `json:"id"`
	Signatures            []string         `json:"signatures"`
	Compression           string           `json:"compression"`
	PackedContextFreeData string           `json:"packed_context_free_data"`
	ContextFreeData       []interface{}    `json:"context_free_data,omitempty"` // TODO: <- check
	PackedTrx             string           `json:"packed_trx"`
	TransactionInfo       *TransactionInfo `json:"transaction"`
}

type TransactionInfo struct {
	Expiration         types.Time    `json:"expiration"`
	RefBlockNum        uint64        `json:"ref_block_num"`
	RefBlockPrefix     uint64        `json:"ref_block_prefix"`
	MaxNetUsageWords   uint64        `json:"max_net_usage_words"`
	MaxCpuUsageMs      uint64        `json:"max_cpu_usage_ms"`
	DelaySec           uint64        `json:"delay_sec"`
	ContextFreeActions []interface{} `json:"context_free_actions,omitempty"`
	Actions            []UnitAction  `json:"actions"`
}

type UnitAction struct {
	Account       string              `json:"account"`
	Name          string              `json:"name"`
	Authorization []UnitAuthorization `json:"authorization"`
	Data          interface{}         `json:"data"`
	HexData       string              `json:"hex_data"`
}

type UnitAuthorization struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

type ActiveSchedule struct {
	Version   uint64                       `json:"version"`
	Producers []ActiveScheduleUnitProducer `json:"producers"`
}

type ActiveScheduleUnitProducer struct {
	ProducerName string         `json:"producer_name"`
	Authority    *UnitAuthority `json:"authority,omitempty"`
}

type UnitAuthority struct {
	Threshold uint64    `json:"threshold"`
	Keys      []AuthKey `json:"keys"`
}

type DuplicatedUnitAuthority UnitAuthority

func (u *UnitAuthority) UnmarshalJSON(b []byte) error {
	interimRet := []interface{}{}

	err := json.Unmarshal(b, &interimRet)
	if err != nil {
		return err
	}

	bytePartial, err := json.Marshal(interimRet[1])
	if err != nil {
		return err
	}

	ret := &DuplicatedUnitAuthority{}
	err = json.Unmarshal(bytePartial, ret)
	if err != nil {
		return err
	}

	*u = UnitAuthority(*ret)

	return nil
}

type UnitProducerInHeaderState struct {
	ProducerName string `json:"producer_name"`
	BlockNum     uint64 `json:"block_num"`
}

func (u *UnitProducerInHeaderState) UnmarshalJSON(b []byte) error {
	interimRet := [2]interface{}{}

	err := json.Unmarshal(b, &interimRet)
	if err != nil {
		return err
	}

	u.ProducerName = interimRet[0].(string)

	blockNum := interimRet[1].(float64)
	u.BlockNum = uint64(blockNum)

	return nil
}

type PendingSchedule struct {
	ScheduleLibNum uint64 `json:"schedule_lib_num"`
	ScheduleHash   string `json:"schedule_hash"`
	Schedule       struct {
		Version   uint64                       `json:"version"`
		Producers []ActiveScheduleUnitProducer `json:"producers"`
	} `json:"schedule,omitempty"`
}

type ProtocolFeatures struct {
	ProtocolFeatures []string `json:"protocol_features"`
}

type ABI struct {
	Version          string                 `json:"version"`
	Types            []interface{}          `json:"types"`
	Structs          []ActionArgsDefinition `json:"structs"`
	Actions          []UnitActionDefinition `json:"actions"`
	Tables           []UnitTable            `json:"tables"`
	RicardianClauses []string               `json:"ricardian_clauses"`
	ErrorMessages    []string               `json:"error_messages"`
	AbiExtensions    []interface{}          `json:"abi_extensions"`
	Variants         []string               `json:"variants"`
	ActionResults    []interface{}          `json:"action_results"`
}

type ActionArgsDefinition struct {
	Name   string `json:"name"`
	Base   string `json:"base"`
	Fields []Args `json:"fields"`
}

type Args struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type UnitActionDefinition struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	RicardianContract string `json:"ricardian_contract"`
}

type UnitTable struct {
	Name      string        `json:"name"`
	IndexType string        `json:"i64"`
	KeyNames  []interface{} `json:"key_names"`
	KeyTypes  []interface{} `json:"key_types"`
	Type      string        `json:"type"`
}

type UnitCurrencyStats struct {
	Supply    types.Coin `json:"supply"`
	MaxSupply types.Coin `json:"max_supply"`
	Issuer    string     `json:"issuer"`
}

type UnitProducerInfo struct {
	Owner         string `json:"owner"`
	TotalVotes    string `json:"total_votes"`
	ProducerKey   string `json:"producer_key"`
	IsActive      uint   `json:"is_active"`
	Url           string `json:"url"`
	UnpaidBlocks  uint64 `json:"unpaid_blocks"`
	LastClaimTime string `json:"last_claim_time"`
	Location      uint   `json:"location"`
}

type AggregatedTableRow struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Table string `json:"table"`
	Payer string `json:"payer"`
	Count uint64 `json:"count"`
}
