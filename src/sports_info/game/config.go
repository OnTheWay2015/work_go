package game

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sports_info/utils"
)

type carr struct {
	Mode  string
	Mode1 string
}
type cfg_db struct {
	Uri    string
	Crypto bool
	User   string
	Pwd    string
}

type cfg_log struct {
	Filename    string
	MaxSize     int  //在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups  int  //保留旧文件的最大个数
	MaxAge      int  //保留旧文件的最大天数
	Compress    bool //是否压缩/归档旧文件
	Log_resbody bool
}
type GameConfig struct {
	Db_collect     cfg_db
	Log            cfg_log
	Svr_port       string
	Version        string
	Configurations []carr
}

type GameConfigManager struct {
	//m_data map[string]interface{}
	Data GameConfig
}

func (inst *GameConfigManager) init() {
	utils.Once(func() {
		inst.initConfig()
	})
}

func (inst *GameConfigManager) initConfig() {

	f, _ := os.Open("sports_info_config.json")
	r := io.Reader(f)
	err := json.NewDecoder(r).Decode(&inst.Data)
	if err != nil {
		fmt.Println("Parse config failed: ", err)
	}
}

func (inst *GameConfigManager) getLogCfg() cfg_log {
	return inst.Data.Log
}
