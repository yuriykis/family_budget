package apiserver

import "net/http"

func Start(config *Config) error {
	server := newServer()
	return http.ListenAndServe(config.BindAddr, server)
}
