package service

import (
	"time"
	conf "bitbucket.org/instinctools/gluten/master/config"
)

var (
	STATUS bool
	MESSAGE       string
	RESPONSE_TIME time.Duration
	EXIT_TIME     time.Duration

	nodes map[string]time.Duration
	config *conf.Config
)

type NodeStore interface {
	CheckNodes()
}

func (node *Node) CheckNodes() {

}

type Node struct {
	IP string
	Time time.Duration
}

func init() {
	STATUS = true
	nodes = make(map[string]time.Duration)

	//load variables from config
	config = conf.GetConfig()
	MESSAGE = config.Message
	RESPONSE_TIME = time.Second * time.Duration(config.ResponseTime)
	EXIT_TIME = time.Second * time.Duration(config.ExitTime)
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
