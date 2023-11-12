package client

import (
	"net"
	"net/http"
	"time"
)

type customTimingTransport struct {
	rtp       http.RoundTripper
	dialer    *net.Dialer
	connStart time.Time
	connEnd   time.Time
	reqStart  time.Time
	reqEnd    time.Time
}

func NewTransport() *customTimingTransport {
	tr := &customTimingTransport{
		dialer: &net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: time.Second * 15,
		},
	}

	tr.rtp = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                tr.dialer.Dial,
		TLSHandshakeTimeout: time.Second * 10,
	}

	return tr
}

func (tr *customTimingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.reqStart = time.Now()
	resp, err := tr.rtp.RoundTrip(r)
	tr.reqEnd = time.Now()
	return resp, err
}

func (tr *customTimingTransport) dial(network, addr string) (net.Conn, error) {
	tr.connStart = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.connEnd = time.Now()
	return cn, err
}

func (tr *customTimingTransport) ReqDuration() time.Duration {
	return tr.Duration() - tr.ConnDuration()
}

func (tr *customTimingTransport) ConnDuration() time.Duration {
	return tr.connEnd.Sub(tr.connStart)
}

func (tr *customTimingTransport) Duration() time.Duration {
	return tr.reqEnd.Sub(tr.reqStart)
}
