package service

import (
	"time"

	"sync"

	conf "bitbucket.org/instinctools/gluten/master/config"
	log "bitbucket.org/instinctools/gluten/shared/logging"
)

var (
	MESSAGE          string
	RETRIEVE_TIMEOUT time.Duration
	EXIT_TIMEOUT     time.Duration

	nodes map[string]time.Time
	mutex sync.Mutex
)

type Node struct {
	IP   string
	Time time.Duration
}

func init() {
	nodes = make(map[string]time.Time)

	//load variables from config
	config := conf.GetConfig().Node
	RETRIEVE_TIMEOUT = time.Second * time.Duration(config.RetrieveTimeout)
	EXIT_TIMEOUT = time.Second * time.Duration(config.ExitTimeout)

	//call checking nodes
	go CheckExitTimeoutNodes()
}

func AddNode(address string) {
	mutex.Lock()
	defer mutex.Unlock()
	nodes[address] = time.Now()
}

func RemoveNode(address string) {
	delete(nodes, address)
}

func Size() int {
	return len(nodes)
}

func GetByKey(key string) time.Time {
	mutex.Lock()
	defer mutex.Unlock()
	return nodes[key]
}

func Exist(key string) bool {
	return GetByKey(key).Second() != 0
}

func CheckExitTimeoutNodes() {
	for {
		mutex.Lock()
		for key, value := range nodes {
			diff := time.Now().Sub(value)
			if diff.Nanoseconds() > EXIT_TIMEOUT.Nanoseconds() {
				log.WithFields(log.Fields{
					"IP: ": key,
				}).Info("The node has been removed")
				RemoveNode(key)
			}
		}
		mutex.Unlock()
		time.Sleep(RETRIEVE_TIMEOUT)
	}
}

func RegisterNode(in string, address string) {
	if Exist(address) {
		log.WithFields(log.Fields{
			"IP: ": address,
		}).Info("Available node")
	} else {
		AddNode(address)
		log.WithFields(log.Fields{
			"IP: ": address,
		}).Info("The node was registered")
	}
}
