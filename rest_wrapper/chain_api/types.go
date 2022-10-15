package chainApi

import (
	"strings"
	"time"
)

type ResponseGetInfo struct {
	ServerVersion             string `json:"server_version"`
	ChainId                   string `json:"chain_id"`
	HeadBlockNum              uint64 `json:"head_block_num"`
	LastIrreversibleBlockNum  uint64 `json:"last_irreversible_block_num"`
	LastIrreversibleBlockId   string `json:"last_irreversible_block_id"`
	HeadBlockId               string `json:"head_block_id"`
	HeadBlockTime             Time   `json:"head_block_time"` // : "2022-10-10T13:34:03.500",
	HeadBlockProducer         string `json:"head_block_producer"`
	VirtualBlockCpuLimit      uint64 `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit      uint64 `json:"virtual_block_net_limit"`
	BlockCpuLimit             uint64 `json:"block_cpu_limit"`
	BlockNetLimit             uint64 `json:"block_net_limit"`
	ServerVersionString       string `json:"server_version_string"`
	ForkDbHeadBlockNum        uint64 `json:"fork_db_head_block_num"`
	ForkDbHeadBlockId         string `json:"fork_db_head_block_id"`
	ServerFullVersionString   string `json:"server_full_version_string"`
	TotalCpuWeight            string `json:"total_cpu_weight"`
	TotalNetWeight            string `json:"total_net_weight"`
	EarliestAvailableBlockNum uint64 `json:"earliest_available_block_num"`
	LastIrreversibleBlockTime Time   `json:"last_irreversible_block_time"`
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

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	timeFormat := "2006-01-02T15:04:05.999999999"
	timeStr := strings.ReplaceAll(string(b), `"`, "")
	appendedTime := strings.Join([]string{timeStr, "000000"}, "")

	ret, err := time.Parse(timeFormat, appendedTime)
	if err != nil {
		return err
	}

	*t = Time(ret)
	return nil
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
