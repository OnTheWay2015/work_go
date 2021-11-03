package game

import (
	"encoding/json"
	"test_gin/db"
	"test_gin/gin"
	"test_gin/yaag/yaag"
	yaag_gin "test_gin/yaag/yaag/gin"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _router *gin.Engine
var G_log *zap.SugaredLogger
var G_config *GameConfigManager
var G_vedios *GameVediosInfoManager
var G_DBCollects *db.DB_base

var G_EMPTY_JSON string

func initGameConfig() {
	G_config = &GameConfigManager{}
	G_config.init()
}

func initGameLog() {
	//logger, err := zap.NewProduction()
	//if err != nil {
	//	fmt.Printf("initGameLog err:%s", err.Error())
	//	return
	//}
	//G_log = logger.Sugar()
	//logfile := G_config.getKeyValueString("logfile", "logs/test.log")
	//file, err := os.Create(logfile)
	//if err != nil {
	//	fmt.Printf("initGameLog openfile err:%s\n", err.Error())
	//	return
	//}
	getEncoder := func() zapcore.Encoder {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	logcfg := G_config.getLogCfg()
	getLogWriter := func() zapcore.WriteSyncer {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   logcfg.Filename,
			MaxSize:    logcfg.MaxSize,    //在进行切割之前，日志文件的最大大小（以MB为单位）
			MaxBackups: logcfg.MaxBackups, //保留旧文件的最大个数
			MaxAge:     logcfg.MaxAge,     //保留旧文件的最大天数
			Compress:   logcfg.Compress,   //是否压缩/归档旧文件
		}
		return zapcore.AddSync(lumberJackLogger)
	}

	//var encoder zapcore.Encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())  //json
	//var encoder zapcore.Encoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()) //normal
	//var writeSyncer zapcore.WriteSyncer = zapcore.AddSync(file)
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	G_log = logger.Sugar()

	//for i := 0; i < 1024*50; i++ {
	//	G_log.Errorf("test log err")
	//	G_log.Infof("test log info")
	//}
}
func initGameDb() {
	G_DBCollects = db.NewDB(
		G_config.Data.Db_collect.User,
		G_config.Data.Db_collect.Pwd,
		G_config.Data.Db_collect.Crypto,
		G_config.Data.Db_collect.Uri,
		"COLLECT_INFOS",
	)
	//G_DBCollects

}

func initGameRouter() {
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "gin",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"": ""},
	})
	_router = gin.Default()

	_router.Use(yaag_gin.Document())

	MidRegs(_router)
	ResRegs(_router)
}

func initGame() {
	//a := InstConfig().getKeyValueString("db_collect", "10.2.2.11:27017")
	//fmt.Print(a)
	initGameConfig()
	initGameLog()
	initGameDb()

	G_vedios = &GameVediosInfoManager{}
	G_vedios.init()

	initGameRouter()

	a := map[string]interface{}{}
	aa, _ := json.Marshal(a)
	G_EMPTY_JSON = string(aa)

}

func startGame() {
	//logfile, err := os.Create("./logs/gin_http.log")
	//if err != nil {
	//	fmt.Println("Could not create log file")
	//}
	//gin.SetMode(gin.DebugMode)
	//gin.DefaultWriter = io.MultiWriter(logfile)
	//gin.DefaultErrorWriter = io.MultiWriter(logfile)

	_router.Run() // 默认时listen and serve on 0.0.0.0:8080
}

func startTimer() {
	for {
		time.Sleep(time.Second * 1)
		G_vedios.update(1)
	}
}

func Start() {
	initGame()
	go startTimer()
	startGame()
}
