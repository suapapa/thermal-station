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
	Scheme             string
	Host               string
	Port               string
	Username, Password string
	CaCert             string
}

func connectBrokerByWSS(config *Config) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	broker := fmt.Sprintf("%s://%s:%s", config.Scheme, config.Host, config.Port)
	opts.AddBroker(broker)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	if config.CaCert != "" {
		certpool := x509.NewCertPool()
		ca, err := os.ReadFile(config.CaCert)
		if err != nil {
			return nil, errors.Wrap(err, "fail to connet broker")
		}
		certpool.AppendCertsFromPEM(ca)

		var tlsConfig tls.Config
		tlsConfig.RootCAs = certpool
		tlsConfig.SessionTicketsDisabled = true
		opts.SetTLSConfig(&tlsConfig)
	}

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
