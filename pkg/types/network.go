package types

type NetworkAddress struct {
	Address string
	Type    NetworkAddressType
}

type NetworkAddressType int

const (
	PrivateNetwork NetworkAddressType = iota
	PublicNetwork
)
