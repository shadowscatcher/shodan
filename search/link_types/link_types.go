package link_types

type LinkType string

const (
	EthernetOrModem LinkType = "\"Ethernet or modem\""
	TunnelOrVPN     LinkType = "\"generic tunnel or VPN\""
	DSL             LinkType = "DSL"
	IPIPorSIT       LinkType = "\"IPIP or SIT\""
	SLIP            LinkType = "SLIP"
	IPSecOrGRE      LinkType = "\"IPSec or GRE\""
	VLAN            LinkType = "VLAN"
	JumboEthernet   LinkType = "\"jumbo Ethernet\""
	Google          LinkType = "Google"
	GIF             LinkType = "GIF"
	PPTP            LinkType = "PPTP"
	Loopback        LinkType = "loopback"
	AX25            LinkType = "\"AX.25 radio modem\""
)
