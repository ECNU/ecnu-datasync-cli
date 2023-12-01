package g

import (
	"encoding/json"
	"github.com/ecnu/ecnu-openapi-sdk-go/sdk"
	"log"
	"sync"

	"github.com/toolkits/pkg/file"
)

/*
GlobalConfig 全局配置
*/
type GlobalConfig struct {
	OAuth2Config sdk.OAuth2Config `json:"oauth2_config"`
	APIConfig    sdk.APIConfig    `json:"api_config"`
	OutputFile   string           `json:"output_file"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

/*
Config 安全的读取和修改配置
*/
func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

/*
ParseConfig 加载配置
*/
func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -config to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c
}
