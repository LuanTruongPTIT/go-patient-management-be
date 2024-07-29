package server

import (
	"net"
	"net/http"

	"github.com/LuanTruongPTIT/patient/pkg/config"
)

type Server struct {
	srv *http.Server
}
type serverConfig struct {
	IP   string
	Port string
}

var ServerCfg *serverConfig

func init() {
	ServerCfg.IP = config.Config.GetString("SERVER_IP")
	ServerCfg.Port = config.Config.GetString("SERVER_PORT")
}
func NewServer(handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:    net.JoinHostPort(ServerCfg.IP, ServerCfg.Port),
			Handler: handler,
		},
	}
}
