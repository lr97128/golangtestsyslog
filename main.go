package main

import (
	"flag"
	"fmt"
	"log/syslog"
)

type SyslogServer struct {
	Network  string
	Address  string
	Port     uint
	Priority string
	Tag      string
}

func main() {
	var ss SyslogServer
	flag.StringVar(&ss.Network, "n", "udp", "The type of network: udp or tcp")
	flag.StringVar(&ss.Address, "a", "127.0.0.1", "The network address")
	flag.UintVar(&ss.Port, "p", 514, "The network port")
	flag.StringVar(&ss.Priority, "P", "debug", "The priority of syslog")
	flag.StringVar(&ss.Tag, "t", "demon", "The tag of syslog")
	flag.Parse()
	switch ss.Priority {
	case "debug":
		SendDebug(ss)
	case "info":
		SendInfo(ss)
	case "alert":
		SendAlert(ss)
	case "error":
		SendError(ss)
	case "critical":
		SendCritical(ss)
	default:
		fmt.Println("Please input your syslog's priority.")
	}
}

func SendDebug(ss SyslogServer) {
	s, err := syslog.Dial(ss.Network, fmt.Sprintf("%s:%d", ss.Address, ss.Port), syslog.LOG_DEBUG, ss.Tag)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.Debug("This is debug syslog")
}

func SendInfo(ss SyslogServer) {
	s, err := syslog.Dial(ss.Network, fmt.Sprintf("%s:%d", ss.Address, ss.Port), syslog.LOG_INFO, ss.Tag)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.Info("This is info syslog")
}

func SendAlert(ss SyslogServer) {
	s, err := syslog.Dial(ss.Network, fmt.Sprintf("%s:%d", ss.Address, ss.Port), syslog.LOG_ALERT, ss.Tag)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.Alert("This is alert syslog")
}

func SendError(ss SyslogServer) {
	s, err := syslog.Dial(ss.Network, fmt.Sprintf("%s:%d", ss.Address, ss.Port), syslog.LOG_ERR, ss.Tag)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.Err("This is error syslog")
}

func SendCritical(ss SyslogServer) {
	s, err := syslog.Dial(ss.Network, fmt.Sprintf("%s:%d", ss.Address, ss.Port), syslog.LOG_CRIT, ss.Tag)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.Crit("This is critical syslog")
}
