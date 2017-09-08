package clustering

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"sync"
	"time"
)

type SubmitF func(node_address string, execution *core.Execution, step core.TestStep)

type NodeStore struct {
	nodes             map[string]time.Time
	checkNodesTimeout time.Duration
	removeNodeTimeout time.Duration
	mutex             sync.Mutex
	onSubmit          SubmitF
}

func NewNodeStore(checkNodesTimeout time.Duration, removeNodeTimeout time.Duration, onSubmit SubmitF) *NodeStore {
	var store *NodeStore
	defer store.runCheckNodesJob()
	store = &NodeStore{
		checkNodesTimeout: time.Second * checkNodesTimeout,
		removeNodeTimeout: time.Second * removeNodeTimeout,
		nodes:             make(map[string]time.Time),
		onSubmit:          onSubmit,
	}
	return store
}

func (store *NodeStore) AddNode(address string) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.nodes[address] = time.Now()
}

func (store *NodeStore) GetNodes() []string {
	keys := make([]string, 0, len(store.nodes))
	for k := range store.nodes {
		keys = append(keys, k)
	}
	return keys
}

func (store *NodeStore) removeNode(address string) {
	delete(store.nodes, address)
}

func (store *NodeStore) Size() int {
	return len(store.nodes)
}

func (store *NodeStore) GetByKey(key string) time.Time {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	return store.nodes[key]
}

func (store *NodeStore) exist(key string) bool {
	return store.GetByKey(key).Second() != 0
}

func (store *NodeStore) runCheckNodesJob() {
	go func() {
		for {
			store.mutex.Lock()
			for key, value := range store.nodes {
				diff := time.Now().Sub(value)
				if diff.Nanoseconds() > store.removeNodeTimeout.Nanoseconds() {
					logging.WithFields(logging.Fields{
						"IP: ": key,
					}).Info("The node has been removed")
					store.removeNode(key)
				}
			}
			store.mutex.Unlock()
			time.Sleep(store.checkNodesTimeout)
		}
	}()
}

func (store *NodeStore) SubmitToAll(execution *core.Execution, step core.TestStep) {
	for _, node := range store.GetNodes() {
		store.onSubmit(node, execution, step)
	}

}

func (store *NodeStore) RegisterNode(address string) {
	if store.exist(address) {
		logging.WithFields(logging.Fields{
			"IP: ": address,
		}).Info("Available node")
	} else {
		store.AddNode(address)
		logging.WithFields(logging.Fields{
			"IP: ": address,
		}).Info("The node was registered")
	}
}
