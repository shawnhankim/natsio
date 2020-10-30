// package natsio provides sample codes for NATS/STAN client/server.
package natsio

// RunClient runs STAN client
func RunClient() {
	var err error
	DispMsg(TypeReq, "run STAN client", err)
}
