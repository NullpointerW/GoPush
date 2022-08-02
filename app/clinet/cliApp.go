package main

import (
	"GoPush/cli"
	"GoPush/logger"
	"net"
	"strconv"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("connect to server %s\n", conn.RemoteAddr().String())

	var (
		id = 114514
	)
	pCli, _ := cli.NewClient(conn, int64(id))
	defer pCli.Close()
	msg := strconv.Itoa(id)
	_, wErr := pCli.Write(msg)
	if wErr != nil {
		logger.Fatalf("write error: %v", wErr)
	}
	logger.Debug("sendId:%d succeed \n", id)
	go cli.SendHeartbeat(pCli)
	go cli.HeartbeatCheck(pCli)

	for {
		msg, err := pCli.Read()
		if err != nil {
			logger.Fatal(err)
			pCli.Close()
			return
		}
		if strings.EqualFold(msg, "pong") {
			pCli.PongRecv()
		} else {
			logger.Infof(msg)
		}
	}
}
