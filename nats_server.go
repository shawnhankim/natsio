// package natsio provides sample codes for NATS/STAN client/server.
package natsio

import (
	"fmt"
	"math/rand"
	"os"

	natsSvr "github.com/nats-io/nats-server/v2/server"
	natsTst "github.com/nats-io/nats-server/v2/test"
	stanSvr "github.com/nats-io/nats-streaming-server/server"
)

var (
	sampleCluster   = "sample-cluster"
	samplePubClient = "sample-pub-client"
	sampleSubClient = "sample-sub-client"
	sampleSubject   = "sample.subject"
	sampleServerIP  = "127.0.0.1"
	samplePortL     = 6000
	samplePortH     = 7000
	samplePortGap   = samplePortH - samplePortL
)

// EventReq is the data strucgture for testing event via STAN
type EventReq struct {
	code int
	data string
}

// PubSub is the type for publisher and subscriber
type PubSub int

// Constant for type
const (
	TypeReq = iota
	TypeRes
)

func (p PubSub) String() string {
	return [...]string{"[REQUEST ]", "[RECEIVER]"}[p]
}

// RunSampleNatsServer executes sample NATS server
func RunSampleNatsServer() (string, *natsSvr.Server, error) {
	opts := natsSvr.Options{
		Host:           sampleServerIP,
		Port:           rand.Intn(samplePortGap) + samplePortL,
		NoLog:          true,
		NoSigs:         true,
		MaxControlLine: 2048,
	}
	return RunSampleNatsServerWithOptions(&opts)
}

// RunSampleStanServer executes sample NATS streaming server
func RunSampleStanServer(clusterID, svrURL string) (*stanSvr.StanServer, error) {
	stanOpts := stanSvr.GetDefaultOptions()
	stanOpts.NATSServerURL = svrURL
	stanOpts.ID = clusterID
	return stanSvr.RunServerWithOpts(stanOpts, nil)
}

// RunSampleNatsServerWithOptions runs a NATS server
func RunSampleNatsServerWithOptions(opts *natsSvr.Options) (
	string, *natsSvr.Server, error) {
	svrURL := fmt.Sprintf("nats://%v:%v", opts.Host, opts.Port)
	return svrURL, natsTst.RunServer(opts), nil
}

// StartSampleServer starts sample NATS server with the parameter of clusterID
func StartSampleServer(clusterID string) (
	string, *natsSvr.Server, *stanSvr.StanServer, error) {
	serverURL, ns, err := RunSampleNatsServer()
	if err != nil {
		return serverURL, ns, nil, err
	}
	ss, err := RunSampleStanServer(clusterID, serverURL)
	return serverURL, ns, ss, err
}

// DispMsg is to display message between failure anc success
func DispMsg(t PubSub, msg string, err error) {
	if nil != err {
		fmt.Printf("- %s Failed  to %s: %v\n", t, msg, err)
		os.Exit(1)
	}
	fmt.Printf("- %s Succeed to %s\n", t, msg)
}

// RunServer runs NATS/STAN servers
func RunServer() {
	fmt.Printf("\n\n")
	fmt.Printf("+----------------------------------------------------------+\n")
	fmt.Printf("| Run NATS / STAN Servers                                  |\n")
	fmt.Printf("+----------------------------------------------------------+\n")

	// Start sample NATS server and STAN server
	url, ns, ss, err := StartSampleServer(sampleCluster)
	defer ns.Shutdown()
	defer ss.Shutdown()
	DispMsg(TypeReq, "run NATS / STAN server", err)
	fmt.Printf("  + URL         : %v\n", url)
	// fmt.Printf("  + NATS server : %v\n", ns)
	// fmt.Printf("  + STAN server : %v\n", ss)

	// Connect to the STAN server
	// cPub := stan.NewConn(logger)
	// err = cPub.Connect(url, testCluster, testPubClient)
	// defer cPub.Close()
	// dispMsg(TypeAPIReq, "connect STAN server for publishing event", err)

}
