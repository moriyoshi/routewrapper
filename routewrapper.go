package routewrapper

import "net"

type Route struct {
	Destination net.IPNet
	Gateway     net.IP
	Interface   *net.Interface
	Flags       map[string]bool
	Expire      int
}

func (route *Route) IsDefaultRoute() bool {
	ones, _ := route.Destination.Mask.Size()
	return (&route.Destination.IP[0] == &net.IPv4zero[0] || &route.Destination.IP[0] == &net.IPv6zero[0]) && ones == 0
}

func (route *Route) DestinationIsNetwork() bool {
	if route.Destination.Mask == nil {
		return false
	}
	ones, bits := route.Destination.Mask.Size()
	if route.Destination.IP.To4() != nil {
		if ones == 32 && bits == 32 {
			return false
		}
	} else {
		if ones == 128 && bits == 128 {
			return false
		}
	}
	return true
}

type Routing interface {
	Routes() ([]Route, error)
	DefaultRoutes() ([]Route, error)
	AddRoute(Route) error
	GetInterface(name string) (*net.Interface, error)
}
