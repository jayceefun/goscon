package upstream

import (
	"net"
)

// Hook defines hook interface for upstream
type Hook interface {
	AfterConnected(local net.Conn, remote net.Conn) (err error)
}

// installed hook
var upstreamHook Hook

func setHook(hook Hook) {
	if upstreamHook != nil {
		panic("setHook again")
	}
	upstreamHook = hook
}

// OnAfterConnected call when upstream connection is connected
func OnAfterConnected(local net.Conn, remote net.Conn) (err error) {
	if upstreamHook == nil {
		return
	}
	err = upstreamHook.AfterConnected(local, remote)
	return
}
