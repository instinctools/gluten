package clustering

import (
	"sync"
	"time"

	conf "bitbucket.org/instinctools/gluten/master/backend/config"
	log "bitbucket.org/instinctools/gluten/shared/logging"
)

var (
	RETRIEVE_TIMEOUT time.Duration
	EXIT_TIMEOUT     time.Duration

	nodes map[string]time.Time
	mutex sync.Mutex
)

func init() {
	nodes = make(map[string]time.Time)

	//load variables from nodesConfig
	nodesConfig := conf.GlobalConfig.Node
	RETRIEVE_TIMEOUT = time.Second * time.Duration(nodesConfig.RetrieveTimeout)
	EXIT_TIMEOUT = time.Second * time.Duration(nodesConfig.ExitTimeout)

	//call checking nodes
	go CheckExitTimeoutNodes()
}

func AddNode(address string) {
	mutex.Lock()
	defer mutex.Unlock()
	nodes[address] = time.Now()
}

func GetNodes() []string {
	keys := make([]string, 0, len(nodes))
	for k := range nodes {
		keys = append(keys, k)
	}
	return keys
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

func RegisterNode(address string) {
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
