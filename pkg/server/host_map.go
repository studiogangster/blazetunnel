package server

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/lucas-clemente/quic-go"
)

// HostMap is a special type to map the subdomain
// configuration to the remote tunneling client
type HostMap struct {
	mx    *sync.Mutex
	hosts map[string]*TunnelState
}

// NewHostMap is a constructor used to create a new hostmap
func NewHostMap() *HostMap {
	return &HostMap{
		hosts: make(map[string]*TunnelState),
		mx:    &sync.Mutex{},
	}
}

// NewStreamFor creates a new stream of the Host Connection exists.
func (hm *HostMap) NewStreamFor(host string) (quic.Stream, error) {
	ts, ok := hm.Get(host)
	if !ok {
		return nil, errors.New("Host not found")
	}

	stream, err := ts.newStream()
	if err == nil {
		return stream, nil
	}

	neterr, ok := err.(net.Error)
	if !ok {
		hm.Delete(host)
		return nil, err
	}

	if neterr.Timeout() {
		log.Printf("[HostMap:NewStreamFor] Timeout occurred. Removing the configuration")
	}

	hm.Delete(host)

	return nil, err
}

// Put is used to map a new host to a tunnel state
// If the host configuration is already present, false is returned
func (hm *HostMap) Put(host string, ts *TunnelState) bool {
	log.Println("server client ", host)
	hm.mx.Lock()
	_, ok := hm.hosts[host]
	if ok {
		hm.mx.Unlock()
		return false
	}

	if ts == nil {
		hm.mx.Unlock()
		return false
	}

	hm.hosts[host] = ts
	hm.mx.Unlock()

	return true
}

// Get is used to get the tunnel state for a given host configuration
func (hm *HostMap) Get(host string) (*TunnelState, bool) {
	hm.mx.Lock()

	for tmp := range hm.hosts {
		log.Println("get", host, "from", tmp)
	}
	ts, ok := hm.hosts[host]
	hm.mx.Unlock()
	return ts, ok
}

// Delete is used to remove a specific host configuration
func (hm *HostMap) Delete(host string) {
	ts, ok := hm.Get(host)
	if ok {
		ts.Close()
	}
	hm.mx.Lock()
	delete(hm.hosts, host)
	hm.mx.Unlock()
}

// TunnelState stores the connection states in a struct
// this is further used to create a new connections to the client
// on demand
type TunnelState struct {
	session   quic.Session
	ctlStream quic.Stream
}

// NewStream is used to open a new stream from the quic session
func (ts *TunnelState) newStream() (quic.Stream, error) {
	return ts.session.OpenStreamSync(context.Background())
}

// Close closes the control stream and the session
func (ts *TunnelState) Close() {
	ts.ctlStream.Close()
	ts.session.Close()
}
