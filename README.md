# libipvs: Pure Go lib to work with IPVS
----

Netlink socket are used to communicate with various kernel subsystems as an RPC system.

This project provides a pure Go client to communicate with IPVS kernel module using generic netlink socket.


## Project Status
### Implemented Methods
```Golang
type IPVSHandle interface {
	Flush() error
	GetInfo() (info Info, err error)
	ListServices() (services []*Service, err error)
	NewService(s *Service) error
	UpdateService(s *Service) error
	DelService(s *Service) error
	ListDestinations(s *Service) (dsts []*Destination, err error)
	NewDestination(s *Service, d *Destination) error
	UpdateDestination(s *Service, d *Destination) error
	DelDestination(s *Service, d *Destination) error
}
```

### TODO
* IPVS state synchronization: support configuring the in-kernel IPVS sync daemon for supporting failover
  between IPVS routers, as done with keepalived `lvs_sync_daemon_interface`

### Acknowledgments
* The code is first “borrowed” from https://github.com/qmsk/clusterf, so all kudos goes @SpComb. I moved the code out into this dedicated project, with a better API, proper tests and concurrency safety, so everyone would benefit from having a good and well-tested package.

### Alternatives
Other pure go implementation of IPVS that maybe useful:
* https://github.com/tehnerd/gnl2go/

## Example code

```Golang
package main

import (
	"fmt"
	"net"
	"syscall"

	"github.com/mqliang/libipvs"
)

func main() {
	h, err := libipvs.New()
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

	svcs, err := h.ListServices()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", svcs)

	svc := libipvs.Service{
		Address:       net.ParseIP("172.192.168.1"),
		AddressFamily: syscall.AF_INET,
		Protocol:      libipvs.Protocol(syscall.IPPROTO_TCP),
		Port:          80,
		SchedName:     libipvs.RoundRobin,
	}

	if err := h.NewService(&svc); err != nil {
		panic(err)
	}

	svcs, err = h.ListServces()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", svcs)

	dst := libipvs.Destination{
		Address:       net.ParseIP("172.192.100.1"),
		AddressFamily: syscall.AF_INET,
		Port:          80,
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
