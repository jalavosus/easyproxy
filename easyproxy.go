package easyproxy

import (
	"fmt"
	"net/http"
	"net/url"
)

// NewClient returns a *http.Client using the provided host and port,
// using ProtocolHTTPS as the protocol for the connection.
func NewClient(host string, port int) (*http.Client, error) {
	return newClientForProtocol(host, port, ProtocolHTTPS)
}

// NewClientForProtocol constructs a *http.Client using the provided host, port, and protocol.
// Will return an error if one happens to pop up.
func NewClientForProtocol(host string, port int, protocol Protocol) (*http.Client, error) {
	return newClientForProtocol(host, port, protocol)
}

func newClientForProtocol(host string, port int, protocol Protocol) (*http.Client, error) {
	transport, err := newTransportForProtocol(host, port, protocol)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Transport: transport,
	}, nil
}

// NewTransport returns a *http.Transport using the provided host and port,
// using ProtocolHTTPS as the protocol for the connection.
func NewTransport(host string, port int) (*http.Transport, error) {
	return newTransportForProtocol(host, port, ProtocolHTTPS)
}

// NewTransportForProtocol constructs a *http.Transport using the provided host, port, and protocol.
// Will return an error if one happens to pop up.
func NewTransportForProtocol(host string, port int, protocol Protocol) (*http.Transport, error) {
	return newTransportForProtocol(host, port, protocol)
}

func newTransportForProtocol(host string, port int, protocol Protocol) (*http.Transport, error) {
	proxyUrl, err := newProxyURLForProtocol(host, port, protocol)
	if err != nil {
		return nil, err
	}

	return &http.Transport{
		Proxy: proxyUrl,
	}, nil
}

// NewProxyURL returns the result of calling http.ProxyURL using the provided host and port,
// using ProtocolHTTPS as the protocol for the connection.
func NewProxyURL(host string, port int) (func(*http.Request) (*url.URL, error), error) {
	return newProxyURLForProtocol(host, port, ProtocolHTTPS)
}

// NewProxyURLForProtocol returns the result of calling http.ProxyURL using the provided host, port, and protocol.
// Will return an error if one happens to pop up.
func NewProxyURLForProtocol(host string, port int, protocol Protocol) (func(*http.Request) (*url.URL, error), error) {
	return newProxyURLForProtocol(host, port, protocol)
}

func newProxyURLForProtocol(host string, port int, protocol Protocol) (func(*http.Request) (*url.URL, error), error) {
	rawUri := fmt.Sprintf("%[1]s://%[2]s:%[3]d", protocol.String(), host, port)
	parsedUri, err := url.Parse(rawUri)
	if err != nil {
		return nil, err
	}

	return http.ProxyURL(parsedUri), nil
}