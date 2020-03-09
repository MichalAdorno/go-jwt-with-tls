package main

import (
	"crypto/tls"
	"jwt_auth/config"
	"jwt_auth/handler"
	"log"
	"net/http"
)

func main() {
	var yamlConfig config.YamlConfig
	yamlConfig.ReadConf()

	tlsRespHeaderHandler := handler.TlsResponseHeaderHandler
	logHandler := handler.LogHttpHandler

	mux := http.NewServeMux()
	mux.HandleFunc("/signin", logHandler(tlsRespHeaderHandler(handler.Signin)))
	mux.HandleFunc("/welcome", logHandler(tlsRespHeaderHandler(handler.Welcome)))
	mux.HandleFunc("/refresh", logHandler(tlsRespHeaderHandler(handler.Refresh)))

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
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
