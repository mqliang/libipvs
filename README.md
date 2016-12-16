# goipvs: Pure Go client to work with IPVS
----

Netlink socket are used to communicate with various kernel subsystems as an RPC system.

This project provides a pure Go client to communicate with IPVS kernel module using generic netlink socket.


## Project Status
### Implemented Methods
* Flush() error
* GetInfo() (info Info, err error)
* ListServces() (services []Service, err error)
* NewService(s *Service) error
* UpdateService(s *Service) error
* DelService(s *Service) error
* ListDestinations(s *Service) (dsts []Destination, err error)
* NewDestination(s *Service, d *Destination) error
* UpdateDestination(s *Service, d *Destination) error
* DelDestination(s *Service, d *Destination) error

### TODO
* IPVS stats export: decode the `IPVS_SVC_ATTR_STATS` and `IPVS_DEST_ATTR_STATS` into a stats struct.
* IPVS state synchronization: support configuring the in-kernel IPVS sync daemon for supporting failover
  between IPVS routers, as done with keepalived `lvs_sync_daemon_interface`


## Example code

```Golang
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
```