package registry

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	pusher "github.com/pusher/pusher-http-go"
)

type Event string

const (
	Register Event = "register"
	Exit           = "exit"
)

func (e Event) String() string {
	switch e {
	case Register:
		return string(Register)
	case Exit:
		return string(Exit)
	default:
		return ""
	}
}

const (
	Channel = "mapped-discovery"
)

type Registrar struct {
	pusher *pusher.Client
}

type Service struct {
	// The path that is links to this service
	Prefix string `json:"prefix"`

	// Public IP of the host running this service
	Address net.IP `json:"address"`

	Port int64 `json:"port"`

	Name     string    `json:"name"`
	ID       uuid.UUID `json:"id"`
	Hostname string    `json:"hostName"`

	HealthCheck struct {
		URL       string `json:"url"`
		Method    string `json:"method"`
		TLSVerify bool   `json:"tlsVerify"`
	} `json:"healthCheck"`
}

func (s Service) Validate() error {
	if s.Address == nil {
		return errors.New("addr is nil")
	}

	if s.Port <= 0 {
		return errors.New("invalid HTTP port")
	}

	_, err := url.Parse(s.HealthCheck.URL)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(s.Hostname)) == 0 {
		return errors.New("please provide the hostname of this service")
	}

	switch s.HealthCheck.Method {
	case http.MethodGet:
		return nil
	default:
		return errors.New("Only GET is supported for Health check")
	}
}

func New(client *pusher.Client) *Registrar {
	return &Registrar{client}
}

func (r *Registrar) do(svc Service, event Event) error {
	if err := svc.Validate(); err != nil {
		return err
	}

	_, err := r.pusher.Trigger(Channel, event.String(), svc)
	return err

}

func (r *Registrar) Register(svc Service) error {
	return r.do(svc, Register)
}

func (r *Registrar) DeRegister(svc Service) error {
	return r.do(svc, Exit)
}

func (r *Registrar) IP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
				return ipnet.IP, nil
			}
		}
	}

	return nil, nil
}
