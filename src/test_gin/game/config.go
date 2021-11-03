package game

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"test_gin/utils"
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
	Filename   string
	MaxSize    int  //在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups int  //保留旧文件的最大个数
	MaxAge     int  //保留旧文件的最大天数
	Compress   bool //是否压缩/归档旧文件
}
type GameConfig struct {
	Db_collect     cfg_db
	Log            cfg_log
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
	//f, _ := os.Open("config.json")
	//r := io.Reader(f)
	//ret := &GameConfig{}
	//json.NewDecoder(r).Decode(ret)

	// 创建Json编码器,写文件
	//ff, _ := os.Create("config1.json")
	//encoder := json.NewEncoder(ff)
	//encoder.Encode(ret)

	f, _ := os.Open("config.json")
	r := io.Reader(f)
	err := json.NewDecoder(r).Decode(&inst.Data)
	if err != nil {
		fmt.Println("Parse config failed: ", err)
	}

	//f, err := ioutil.ReadFile("config.json")
	//if err != nil {
	//	fmt.Println("load config error: ", err)
	//}
	//err = json.Unmarshal(f, &inst.m_data)
	//if err != nil {
	//	fmt.Println("Para config failed: ", err)
	//}
}

func (inst *GameConfigManager) getKeyValueString(k string, def string) string {
	//if inst.m_data == nil || inst.m_data[k] == nil {
	//	return def
	//}
	//obj := inst.m_data[k]
	//objimpl := obj.(string)
	//if len(objimpl) <= 0 {
	//	return def
	//}
	//return objimpl
	return ""
}

func (inst *GameConfigManager) getKeyValueNumber(k string, def int32) int32 {
	//if inst.m_data == nil || inst.m_data[k] == nil {
	//	return def
	//}
	//obj := inst.m_data[k]
	//objimpl := obj.(int32)
	//return objimpl
	return 0
}
func (inst *GameConfigManager) getLogCfg() cfg_log {
	return inst.Data.Log
}
