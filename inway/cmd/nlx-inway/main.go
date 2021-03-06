// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"fmt"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/VNG-Realisatie/nlx/common/nlxtls"
	"github.com/VNG-Realisatie/nlx/inway"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"localhost:2018" description:"Adress for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	nlxtls.TLSOptions
}

func main() {
	// Parse options
	args, err := flags.Parse(&options)
	if err != nil {
		if et, ok := err.(*flags.Error); ok {
			if et.Type == flags.ErrHelp {
				return
			}
		}
		log.Fatalf("error parsing flags: %v", err)
	}
	if len(args) > 0 {
		log.Fatalf("unexpected arguments: %v", args)
	}

	// Setup new zap logger
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	defer func() {
		syncErr := logger.Sync()
		if syncErr != nil {
			// notify the user that proper logging has failed
			fmt.Fprintf(os.Stderr, "failed to sync zap logger: %v\n", syncErr)
			// don't exit when we're in a panic
			if p := recover(); p != nil {
				panic(p)
			}
			os.Exit(1)
		}
	}()

	// Load certs
	roots, cert, err := nlxtls.Load(options.TLSOptions)
	if err != nil {
		fmt.Println(err)
		logger.Fatal("failed to load tls certs", zap.Error(err))
	}
	_ = roots
	logger.Info(fmt.Sprintf("According to my organization certificate, I am %s\n", cert.DNSNames[0]))

	// Create new inway and provide it with a hardcoded service.
	iw := inway.NewInway(logger, "DemoProviderOrganization")
	// NOTE: hardcoded service endpoint because we don't have any other means to register endpoints yet
	echoServiceEndpoint, err := inway.NewHTTPServiceEndpoint(logger, "PostmanEcho", "https://postman-echo.com/")
	if err != nil {
		logger.Fatal("failed to create PostmanEcho service", zap.Error(err))
	}
	iw.AddServiceEndpoint(echoServiceEndpoint)
	// Listen on the address provided in the options
	err = iw.ListenAndServe(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
