package tags

/*

Tags provide additional context and information about the type of service that is provided. Most of them are only added
after further validation/analysis has been done that is outside the scope of the regular banner collection.

*/

const (
	// The cloud tag is added to banners where the IP is located in the network range of popular cloud hosting
	// providers such as Amazon AWS, Microsoft Azure and Google Cloud.
	Cloud = "cloud"

	// The compromised tag is added to services that appear to have been compromised by an attacker.
	// Currently, it is only being applied to NoSQL database instances that have been compromised for ransomware.
	Compromised = "compromised"

	// The cryptocurrency tag is added to services related to crypto currencies, such as Ethereum RPC, Bitcoin
	// and Monero.
	Cryptocurrency = "cryptocurrency"

	// The database tag is added to confirmed database instances.
	Database = "database"

	// The devops tag is added to services that are used by DevOps technology stacks, such as Docker.
	Devops = "devops"

	// The doublepulsar tag is added to SMB banners that are responding to the DoublePulsar backdoor handshake.
	Doublepulsar = "doublepulsar"

	// The honeypot tag is added to services that appear to be pretending to be a certain type of service/
	// device. The honeypot tag is mutually exclusive with other tags such as ics; a banner can’t have both an
	// ics tag and a honeypot tag.
	Honeypot = "honeypot"

	// The ics tag is added to services that are confirmed to be industrial control systems (ICS) based on the
	// protocol they’re exposing.  The crawlers only tag industrial protocols and don’t consider web servers or
	// VNC when classifying a device as ics.
	ICS = "ics"

	// The iot tag is added to services belonging to Internet of Things devices. This tag is still in active
	// development and will see significantly more usage as IoT products gain in popularity.
	IoT = "iot"

	// The malware tag is added to banners for verified command and control (C2) servers. It is part of the
	// special activity performed by Malware Hunter (https://malware-hunter.shodan.io).
	Malware = "malware"

	// The medical tag is added to services that belong to medical devices. At the moment, the Shodan
	// crawlers only consider DICOM responses for inclusion of the medical tag.
	Medical = "medical"

	// The onion tag is added to services that are providing content on the  Tor network but are leaking into
	// the Clearnet.
	Onion = "onion"

	// The scanner tag is added to banners that are coming from devices which have been seen scanning the Internet.
	Scanner = "scanner"

	// The self-signed tag is added to services where the SSL/ TLS certificate looks like it was self-signed.
	SelfSigned = "self-signed"

	// The starttls tag is added to services that support and successfully completed a StartTLS connection upgrade.
	StartTLS = "starttls"

	// The tor tag is added the services that look like they belong to a Tor node.
	Tor = "tor"

	// The videogame tag is added to banners that were gathered from a video game, such as Minecraft.
	Videogame = "videogame"

	// The vpn tag is added to VPN services that successfully complete the initial handshake. The following
	// protocols are currently tested: IKE, IKE NAT-T, PPTP.
	VPN = "vpn"
)
