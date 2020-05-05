package models

import (
	"encoding/json"
)

type SSL struct {
	// List of acceptable certificate authorities (CA)
	AcceptableCAs []SslAcceptableCA `json:"acceptable_cas"`

	// List of supported HTTP versions (ex. “h2” to indicate the server supports HTTP/2.0)
	Alpn []string `json:"alpn"`

	// The parsed certificate properties that includes information such as when it was issued, the SSL extensions,
	// the issuer, subject etc
	Cert SslCert `json:"cert"`

	// List of SSL certificates (PEM format) that are part of the certificate chain, including leaf and root
	// certificate.
	Chain []string `json:"chain"`

	// Preferred cipher for the SSL connection
	Cipher SslCipher `json:"cipher"`

	// The Diffie-Hellman parameters if available: "prime", "public_key", "bits", "generator" and an optional
	// "fingerprint" if we know which program generated these parameters
	DHparams *SslDHParams `json:"dhparams,omitempty"`

	// List of TLS extensions that the server supports
	TLSExt   []SslTlsExt `json:"tlsext"`
	Unstable []string    `json:"unstable,omitempty"`

	// List of supported SSL versions. If the value starts with a “-“ then the service does NOT support that version
	// (ex. “-SSLv2” means the service doesn’t support SSLv2).
	Versions []string `json:"versions"`
}

type SslAcceptableCA struct {
	Components SslCertComponents `json:"components"`
	Hash       int               `json:"hash"`
	Raw        string            `json:"raw"`
}

type SslCert struct {
	Expired     bool           `json:"expired"`
	Expires     string         `json:"expires"`
	Extensions  []SslExtension `json:"extensions"`
	Fingerprint SslFingerprint `json:"fingerprint"`
	Issued      string         `json:"issued"`
	Issuer      SslIssuer      `json:"issuer"`
	Pubkey      Pubkey         `json:"pubkey"`
	Serial      json.Number    `json:"serial,Number"`
	SigAlg      string         `json:"sig_alg"`
	Subject     SslSubject     `json:"subject"`
	Version     int            `json:"version"`
}

type SslCertComponents struct {
	C            string `json:"C,omitempty"`
	CN           string `json:"CN,omitempty"`
	DC           string `json:"DC,omitempty"`
	L            string `json:"L,omitempty"`
	O            string `json:"O,omitempty"`
	OU           string `json:"OU,omitempty"`
	SN           string `json:"SN,omitempty"`
	ST           string `json:"ST,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
}

type SslCipher struct {
	Bits    int    `json:"bits"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type SslDHParams struct {
	Bits        int         `json:"bits"`
	Fingerprint string      `json:"fingerprint,omitempty"`
	Generator   interface{} `json:"generator"`
	Prime       string      `json:"prime"`
	PublicKey   string      `json:"public_key"`
}

type SslExtension struct {
	Critical bool   `json:"critical,omitempty"`
	Data     string `json:"data"`
	Name     string `json:"name"`
}

type SslFingerprint struct {
	SHA1   string `json:"sha1"`
	SHA256 string `json:"sha256"`
}

type Pubkey struct {
	Bits int    `json:"bits"`
	Type string `json:"type"`
}

type SslSubject struct {
	SslCertComponents
	BusinessCategory string `json:"businessCategory,omitempty"`
	Description      string `json:"description,omitempty"`
	JurisdictionC    string `json:"jurisdictionC,omitempty"`
	JurisdictionSt   string `json:"jurisdictionST,omitempty"`
	PostalCode       string `json:"postalCode,omitempty"`
	Street           string `json:"street,omitempty"`
}

type SslIssuer struct {
	SslCertComponents
	Name                string `json:"name,omitempty"`
	UID                 string `json:"UID,omitempty"`
	DNQualifier         string `json:"dnQualifier,omitempty"`
	SubjectAltName      string `json:"subjectAltName,omitempty"`
	UnstructuredName    string `json:"unstructuredName,omitempty,omitempty"`
	UnstructuredAddress string `json:"unstructuredAddress,omitempty,omitempty"`
	PostalCode          string `json:"postalCode,omitempty,omitempty"`
	Street              string `json:"street,omitempty,omitempty"`
	Undef               string `json:"UNDEF,omitempty"`
}

type SslTlsExt struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
