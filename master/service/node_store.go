package service

import (
	"time"
)

var STATUS bool
var nodes map[string]time.Duration

const (
	MESSAGE       string        = "200"
	RESPONSE_TIME time.Duration = time.Second * 5
	EXIT_TIME     time.Duration = time.Second * 30
)

func init() {
	STATUS = true
	nodes = make(map[string]time.Duration)
}

func ChangeStatus() {
	STATUS = false
}

func AddNode(address string) {
	nodes[address] = RESPONSE_TIME
}

func RemoveNode(address string) {
	delete(nodes, address)
}

func Size() int {
	return len(nodes)
}

func GetByKey(key string) time.Duration {
	return nodes[key]
}

func Exist(key string) bool {
	return GetByKey(key) != time.Second * 0
}

func AddResponseTimeForNode(key string) {
	if Exist(key) {
		nodes[key] = GetByKey(key) + RESPONSE_TIME
	}
}
