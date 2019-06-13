package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/smallnest/rpcx"

	"bindolabs/bindocommon/env"
)

var allClient = make(map[string]*rpcx.Client)
var clientMutex sync.Mutex

//get a service connection,but u do,nt close it.
//because it a connection pool
func NewClient(name string) (*rpcx.Client, error) {
	clientMutex.Lock()
	if c, ok := allClient[name]; ok {
		clientMutex.Unlock()
		return c, nil
	}
	clientMutex.Unlock()
	config, err := env.ServiceGetConfig(name)
	if err != nil {
		return nil, err
	}
	newClient := rpcx.NewClient(&rpcx.DirectClientSelector{Network: config.Protocol,
		Address:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		DialTimeout: time.Duration(config.DialTimeout) * time.Millisecond,
	})
	newClient.Timeout = time.Duration(config.Timeout) * time.Millisecond
	clientMutex.Lock()
	allClient[name] = newClient
	clientMutex.Unlock()
	return newClient, nil
}
