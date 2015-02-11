package main

// Container response from describe
// generated struct using: http://mholt.github.io/json-to-go/
type Container struct {
	ID          string `json:"id"`
	ContainerID string `json:"containerId"`
	Name        string `json:"name"`
	PublicDNS   string `json:"publicDns"`
	IP          string `json:"ip"`
	Source      string `json:"source"`
	State       string `json:"state"`
	CPUShares   int    `json:"cpuShares"`
	Memory      int    `json:"memory"`
	DiskSize    int    `json:"diskSize"`
	Ordered     int64  `json:"ordered"`
	Created     int64  `json:"created"`
	Terminated  int64  `json:"terminated"`
	OrderedBy   string `json:"orderedBy"`
	Plan        string `json:"plan"`
	Ports []struct {
		ID          string `json:"id"`
		PrivatePort int    `json:"privatePort"`
		PublicPort  int    `json:"publicPort"`
		Scheme      string `json:"scheme"`
	} `json:"ports"`
	Env            []interface{} `json:"env"`
	Cmd []interface{} `json:"cmd"`
	PricePerMinute struct {
		Amount   float32 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"pricePerMinute"`
	PricePerOrder struct {
		Amount   float32 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"pricePerOrder"`
	Cost struct {
		Amount float32 `json:"amount"`
		Currency string `json:"currency"`
		} `json:"cost"`
	BuildLog    interface{} `json:"buildLog"`
}

type createContainerRequest struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Plan   string `json:"plan"`
}
