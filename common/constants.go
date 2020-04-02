package common

import "sync"

// All the constants defined for communication
// between the peers
const (
	CommandSetConfig  = "set:config"
	CommandNewClient  = "new:client"
	CommandPingPeer   = "ping:peer"
	CommandPongPeer   = "pong:peer"
	CommandAuthServer = "new:authserver"
	CommandAuthClient = "new:authclient"

	CommandRegisterClient = "new:regserver"
	CommandRegisterServer = "new:regclient"
)

// Secretkey: To be used for encryption
var (
	Secretkey string
	mutex     sync.Mutex
)

// SetSecretKey  To be used to set encryption
func SetSecretKey(secretKey string) {
	Secretkey = secretKey
}
