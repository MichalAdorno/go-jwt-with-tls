package main

import (
	"crypto/tls"
	"flag"
	"jwt_auth/config"
	"jwt_auth/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var yamlConfig config.YamlConfig
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "/target/config.yaml", "path to config yaml file")
	flag.Parse()
	yamlConfig.ReadConf(configFilePath)

	sigChan := make(chan os.Signal)
	listenForOsSignals(sigChan)
	tlsRespHeaderHandler := handler.TlsResponseHeaderHandler
	logHandler := handler.LogHttpHandler

	mux := http.NewServeMux()
	mux.HandleFunc("/signin", logHandler(tlsRespHeaderHandler(handler.Signin)))
	mux.HandleFunc("/welcome", logHandler(tlsRespHeaderHandler(handler.Welcome)))
	mux.HandleFunc("/refresh", logHandler(tlsRespHeaderHandler(handler.Refresh)))

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         yamlConfig.GetServerPort(),
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS(yamlConfig.GetCertPath(), yamlConfig.GetKeyPath()))

}

func listenForOsSignals(sigChan chan os.Signal) {
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGINT:
			log.Printf("Exiting: received the SIGINT signal.")
		case syscall.SIGTERM:
			log.Printf("Exiting: received the SIGTERM signal.")
		}
		os.Exit(1)
	}()
}
