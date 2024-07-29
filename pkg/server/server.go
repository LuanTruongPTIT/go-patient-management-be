package server

import "net/http"

type Server struct {
	srv *http.Server
}
type serverConfig struct {
	IP   string
	Port string
}
var ServerCfg *serverConfig
func init(){
	ServerCfg.IP = config.Config.Get
}
