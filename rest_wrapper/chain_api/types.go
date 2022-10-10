package chainApi

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
