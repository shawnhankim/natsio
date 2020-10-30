// Package nats provides sample codes for NATS/STAN client/server.
package nats

// RunClient runs STAN client
func RunClient() {
	var err error
	DispMsg(TypeReq, "run STAN client", err)
}
