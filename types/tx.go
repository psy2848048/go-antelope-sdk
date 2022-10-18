package types

type Tx struct {
	Transaction TransactionInfo `json:"transaction"`
	Compression string          `json:"compression"`
	Signatures  []string        `json:"signatures"`
}

type TransactionInfo struct {
	Expiration            Time         `json:"expiration"`
	RefBlockNum           uint64       `json:"ref_block_num"`
	RefBlockPrefix        uint64       `json:"ref_block_prefix"`
	MaxNetUsageWords      uint64       `json:"max_net_usage_words"`
	MaxCpuUsageMs         uint64       `json:"max_cpu_usage_ms"`
	DelaySec              uint64       `json:"delay_sec"`
	ContextFreeActions    []UnitAction `json:"context_free_actions"`
	Actions               []UnitAction `json:"actions"`
	TransactionExtensions []uint64     `json:"transaction_extensions"`
}

type UnitAction struct {
	Account       string              `json:"account"`
	Name          string              `json:"name"`
	Authorization []UnitAuthorization `json:"authorization"`
	Data          string              `json:"data"` // hex data
}

type UnitAuthorization struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
