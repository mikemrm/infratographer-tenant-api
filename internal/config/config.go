package config

import (
	"go.infratographer.com/tenant-api/pkg/echox"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
)

var AppConfig struct {
	CRDB    crdbx.Config
	Logging loggingx.Config
	Server  echox.Config
	Tracing otelx.Config
	URNs    struct {
		ServiceAccounts string
		Tenants         string
		Tokens          string
		Users           string
	}
	Bootstrap struct {
		Tenant struct {
			Name string
		}
		Issuer struct {
			URI      string
			Audience string
			Claims   struct {
				Subject string
				Email   string
				Name    string
			}
		}
		User struct {
			Subject string
			Email   string
		}
	}
}
