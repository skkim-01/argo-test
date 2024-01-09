package instance

import (
	"sync"
)

type AgentInstance struct {
	// TODO: define instance data
}

var instance *AgentInstance
var once sync.Once

func GetInstance() *AgentInstance {
	once.Do(func() {
		// TODO: initialize code
		instance = &AgentInstance{}
	})
	return instance
}

func Init() {
	GetInstance()
}
