package http

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/bi-zone/sonar/pkg/server"
)

var defaultOptions = options{
	idleTimeout:       time.Second * 5,
	sessionTimeout:    time.Second * 5,
	tlsConfig:         nil,
	notifyStartedFunc: func() {},
	notifyRequestFunc: func(net.Addr, []byte, map[string]interface{}) {},
}

type options struct {
	idleTimeout       time.Duration
	sessionTimeout    time.Duration
	tlsConfig         *tls.Config
	notifyStartedFunc func()
	notifyRequestFunc server.NotifyRequestFunc
}

type Option func(*options)

func IdleTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.idleTimeout = d
	}
}

func SessionTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.sessionTimeout = d
	}
}

func TLSConfig(c *tls.Config) Option {
	return func(opts *options) {
		opts.tlsConfig = c
	}
}

func NotifyStartedFunc(f func()) Option {
	return func(opts *options) {
		opts.notifyStartedFunc = f
	}
}

func NotifyRequestFunc(f server.NotifyRequestFunc) Option {
	return func(opts *options) {
		opts.notifyRequestFunc = f
	}
}
