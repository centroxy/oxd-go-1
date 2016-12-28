package conf

import (
	"github.com/BurntSushi/toml"
	"oxd-client/utils"
	"fmt"
	"strings"
	"github.com/juju/loggo"
)

type Configuration struct {
	Host string `toml:"host"`
	OpHost string `toml:"opHost"`
	RedirectUrl string `toml:"redirectUrl"`
	PostLogoutRedirectUrl string `toml:"postLogoutRedirectUrl"`
	LogoutUrl string `toml:"logoutUrl"`
	UserId string `toml:"userId"`
	UserSecret string `toml:"userSecret"`
	ClientId string `toml:"clientId"`
	ClientSecret string `toml:"clientSecret"`
	Loglevel string `toml:"logLevel"`
	Scope string `toml:"scope"`
}

func  loadProperties(c *Configuration){
	_, err := toml.DecodeFile("../conf/main.toml", c)
	utils.CheckError("configuration","Config file error",err)
}

var TestConfiguration *Configuration

func init(){
	if TestConfiguration == nil {
		TestConfiguration = new(Configuration)
		loadProperties(TestConfiguration)
		fmt.Println(TestConfiguration.OpHost)
	}

	if strings.EqualFold(TestConfiguration.Loglevel, "Debug") {
		loggo.GetLogger("default").SetLogLevel(loggo.DEBUG)
	}
}