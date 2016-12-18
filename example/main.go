package main

import (
	"fmt"
	"net"
	"syscall"

	"github.com/mqliang/goipvs"
)

func main() {
	h, err := goipvs.New()
	if err != nil {
		panic(err)
	}
	if err := h.Flush(); err != nil {
		panic(err)
	}

	info, err := h.GetInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", info)

	svcs, err := h.ListServces()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", svcs)

	svc := goipvs.Service{
		Address:       net.ParseIP("172.192.168.1"),
		AddressFamily: syscall.AF_INET,
		Protocol:      goipvs.Protocol(syscall.IPPROTO_TCP),
		Port:          80,
		SchedName:     goipvs.RoundRobin,
	}

	if err := h.NewService(&svc); err != nil {
		panic(err)
	}

	svcs, err = h.ListServces()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", svcs)

	dst := goipvs.Destination{
		Address: net.ParseIP("172.192.100.1"),
		Port:    80,
	}

	if err := h.NewDestination(&svc, &dst); err != nil {
		panic(err)
	}

	dsts, err := h.ListDestinations(&svc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", dsts)
}
