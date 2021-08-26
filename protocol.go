package easyproxy

//go:generate stringer -type=Protocol -linecomment
type Protocol uint8

const (
	ProtocolHTTP Protocol = iota // http
	ProtocolHTTPS // https
	ProtocolSOCKS5 // socks5
)