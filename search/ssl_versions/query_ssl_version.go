package ssl_versions

type SSLVersion string

const (
	SSLv2   SSLVersion = "sslv2"
	SSLv3   SSLVersion = "sslv3"
	TLSv1   SSLVersion = "tlsv1"
	TLSv1_1 SSLVersion = `"tlsv1.1"`
	TLSv1_2 SSLVersion = `"tlsv1.2"`
	TLSv1_3 SSLVersion = `"tlsv1.3"`
)
