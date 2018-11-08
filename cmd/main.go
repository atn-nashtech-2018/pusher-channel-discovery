package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adelowo/pusher-channel-discovery/registry"
	"github.com/google/uuid"
	pusher "github.com/pusher/pusher-http-go"
)

func main() {

	shutDownChan := make(chan os.Signal)
	signal.Notify(shutDownChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	port := flag.Int64("http.port", 3000, "Port to run HTTP server at")

	appID := os.Getenv("PUSHER_APP_ID")
	appKey := os.Getenv("PUSHER_APP_KEY")
	appSecret := os.Getenv("PUSHER_APP_SECRET")
	appCluster := os.Getenv("PUSHER_APP_CLUSTER")
	appIsSecure := os.Getenv("PUSHER_APP_SECURE")

	var isSecure bool
	if appIsSecure == "1" {
		isSecure = true
	}

	client := &pusher.Client{
		AppId:   appID,
		Key:     appKey,
		Secret:  appSecret,
		Cluster: appCluster,
		Secure:  isSecure,
	}

	reg := registry.New(client)

	ip, err := reg.IP()
	if err != nil {
		log.Fatalf("could not fetch public IP address... %v", err)
	}

	svc := registry.Service{
		Prefix:  "/v2",
		Address: ip,
		Port:    *port,
		Name:    "Url shortner",
		ID:      uuid.New(),
		HealthCheck: struct {
			URL       string `json:"url"`
			Method    string `json:"method"`
			TLSVerify bool   `json:"tlsVerify"`
		}{
			URL:       fmt.Sprintf("http://%s/health", ip.String()),
			Method:    http.MethodGet,
			TLSVerify: false,
		},
	}

	if err := reg.Register(svc); err != nil {
		log.Fatalf("Could not register service... %v", err)
	}

	<-shutDownChan
	reg.DeRegister(svc)
	fmt.Println("Application is shutting down")
}
