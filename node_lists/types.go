package nodeLists

import "encoding/json"

type UnitBPInfo struct {
	ProducerAccountName string         `json:"producer_account_name"`
	Nodes               []UnitNodeInfo `json:"nodes"`
}

type UnitNodeInfo struct {
	Location    UnitNodeLocation `json:"location"`
	NodeType    NodeType         `json:"node_type,omitempty"`
	ApiEndpoint *string          `json:"api_endpoint,omitempty"`
	SslEndpoint *string          `json:"ssl_endpoint,omitempty"`
	Features    []string         `json:"features,omitempty"`
}

type UnitNodeLocation struct {
	Name      string   `json:"name"`
	Country   string   `json:"country"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type NodeType []string

// NodeType can be unit string or string array
func (n *NodeType) UnmarshalJSON(data []byte) (err error) {
	// 1. input is array
	ret := []string{}

	err = json.Unmarshal(data, &ret)
	if err == nil {
		*n = NodeType(ret)
		return nil
	}

	// 2. input is string
	ret = []string{string(data)}
	*n = NodeType(ret)
	return nil
}
