package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Serve() {
	server := http.Server{
		Addr:    ":8083",
		Handler: initMux(),
	}
	configServer(&server)
	sertPath, keyPath := "./cert/certificate.crt", "./cert/privateKey.key"
	val, ok := os.LookupEnv("SERT_PATH")
	if ok {
		sertPath = val
	}
	val, ok = os.LookupEnv("_PATH")
	if ok {
		keyPath = val
	}
	if err := server.ListenAndServeTLS(sertPath, keyPath); err != nil {
		log.Fatal(err)
	}
}
func configServer(s *http.Server) {
	val, ok := os.LookupEnv("READ_TIMEOUT")
	if ok {
		time, err := time.ParseDuration(val)
		if err == nil {
			s.ReadTimeout = time
		}
	}
	val, ok = os.LookupEnv("WRITE_TIMEOUT")
	if ok {
		time, err := time.ParseDuration(val)
		if err == nil {
			s.WriteTimeout = time
		}
	}
}
