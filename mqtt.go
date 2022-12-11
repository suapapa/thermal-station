package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
)

type Config struct {
	ClientID           string
	Host               string
	Port               string
	Username, Password string
	CaCert             string
}

func connectBrokerByWSS(config *Config) (mqtt.Client, error) {
	if config.CaCert == "" {
		config.CaCert = "/etc/ssl/certs/ca-certificates.crt"
	}

	certpool := x509.NewCertPool()
	ca, err := os.ReadFile(config.CaCert)
	if err != nil {
		return nil, errors.Wrap(err, "fail to connet broker")
	}
	certpool.AppendCertsFromPEM(ca)

	var tlsConfig tls.Config
	tlsConfig.RootCAs = certpool
	tlsConfig.SessionTicketsDisabled = true

	opts := mqtt.NewClientOptions()
	broker := fmt.Sprintf("wss://%s:%s", config.Host, config.Port)
	opts.AddBroker(broker)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetTLSConfig(&tlsConfig)
	opts.SetOrderMatters(false)
	opts.SetClientID(config.ClientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, errors.Wrap(err, "fail to connet broker")
	}
	return client, nil
}
