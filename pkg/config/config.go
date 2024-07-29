package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	configEnv := strings.ToLower(os.Getenv("CONFIG_ENV"))
	if len(configEnv) == 0 {
		configEnv = DevelopmentEnv
	}
	configPrefix := strings.ToUpper(configEnv)
	Config.SetEnvPrefix(configPrefix)
	Config.AutomaticEnv()
	configLoadFile()
}
func configLoadFile() {
	err := Config.ReadInConfig()
	if err != nil {
		log.Println("{\"label\":\"config-load-file\",\"level\":\"warning\",\"msg\":\"error loading config file, " +
			err.Error() + "\",\"service\":\"" + Config.GetString("SERVER_NAME") +
			"\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")
	}
}
