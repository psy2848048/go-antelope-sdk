package chainApi

import "github.com/psy2848048/go-antelope-sdk/types"

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
	Weight     uint64 `json:"weight"`
	Permission struct {
		Actor      string `json:"actor"`
		Permission string `json:"permission"`
	} `json:"permission"`
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
