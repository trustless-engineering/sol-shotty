package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ClusterMember struct {
	Pubkey  string `json:"pubkey"`
	Gossip  string `json:"gossip"`
	RPC     string `json:"rpc"`
	TPU     string `json:"tpu"`
	Version string `json:"version"`
}

func GetClusterEndpoints() ([]string, error) {
	endpoint := "https://api.mainnet-beta.solana.com"
	jsonBody := []byte(`{ "jsonrpc": "2.0", "id": 1, "method": "getClusterNodes"}`)

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, endpoint, bodyReader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	result, ok := response["result"].([]interface{})

	if !ok {
		return nil, fmt.Errorf("unexpected response structure")
	}

	var clusterEndpoints []string
	for _, rawMember := range result {
		member, ok := rawMember.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected entity structure")
		}
		rpc, ok := member["rpc"].(string)
		if !ok {
			continue
		}
		clusterEndpoints = append(clusterEndpoints, "http://"+rpc)
	}

	return clusterEndpoints, nil
}

func LoadEndpoints() ([]string, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	endpoints := config.Endpoints

	if config.UseClusterNodes {
		clusterEndpoints, err := GetClusterEndpoints()
		if err != nil {
			log.Print("error loading cluster endpoints")
		}

		endpoints = append(endpoints, clusterEndpoints...)
	}
	return endpoints, nil
}
