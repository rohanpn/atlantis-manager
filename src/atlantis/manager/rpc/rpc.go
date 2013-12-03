package rpc

import (
	atlantis "atlantis/common"
	"atlantis/manager/crypto"
	"atlantis/manager/manager"
	"atlantis/manager/supervisor"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strings"
	"time"
)

type ManagerRPC bool

var (
	lAddr                string
	lPort                string
	l                    net.Listener
	server               *rpc.Server
	config               *tls.Config
	CPUSharesIncrement   = uint(1) // default to no increment
	MemoryLimitIncrement = uint(1) // default to no increment
)

func Init(listenAddr string, supervisorPort uint16, cpuIncr, memIncr uint, resDuration time.Duration) error {
	var err error
	err = LoadEnvs()
	if err != nil {
		return err
	}
	CPUSharesIncrement = cpuIncr
	MemoryLimitIncrement = memIncr
	atlantis.Tracker.ResultDuration = resDuration
	// init rpc stuff here
	lAddr = listenAddr
	lPort = strings.Split(lAddr, ":")[1]
	supervisor.Init(fmt.Sprintf("%d", supervisorPort))
	manager.Init(lPort)
	manager := new(ManagerRPC)
	server = rpc.NewServer()
	server.Register(manager)
	config := &tls.Config{}
	config.InsecureSkipVerify = true
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.X509KeyPair(crypto.SERVER_CERT, crypto.SERVER_KEY)

	l, err = tls.Listen("tcp", lAddr, config)
	return err
}

func Listen() {
	if l == nil {
		panic("Not Initialized.")
	}
	log.Println("[RPC] Listening on", lAddr)
	server.Accept(l)
}
