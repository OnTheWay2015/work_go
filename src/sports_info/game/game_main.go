package game

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"proto/heathcheck"
	"proto/vedios_info"
	"sports_info/controllers"
	"sports_info/db"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"tsEngine/tsGrpc"
	"tsEngine/tsNacos"

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
	fmt.Println("svr_port:", G_config.Data.Svr_port)
	fmt.Println("db_collect uri:", G_config.Data.Db_collect.Uri)
}

func initGameLog() {
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

	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	G_log = logger.Sugar()
}
func initGameDb() {
	var err string
	G_DBCollects, err = db.NewDB(
		G_config.Data.Db_collect.User,
		G_config.Data.Db_collect.Pwd,
		G_config.Data.Db_collect.Crypto,
		G_config.Data.Db_collect.Uri,
		"COLLECT_INFOS",
	)
	if err != "" {
		G_log.Errorf("G_DBCollects init fail. err:%s")
	}

	//cur := G_DBCollects.Find("a1", bson.M{}, nil)
	//if cur == nil {
	//	return
	//}
	//for cur.Next(context.Background()) {
	//	var dd DBGameVedioInfo
	//	err := cur.Decode(&dd)
	//	if err == nil {
	//		fmt.Printf("infos:%v", dd)
	//		continue
	//	}
	//}

}

func initGameRouter() {
	gin.SetMode(gin.ReleaseMode)
	//_router = gin.Default()
	_router = gin.New()
	_router.Use(gin.Recovery())

	MidRegs(_router)
	ResRegs(_router)
}
func initgrpc() {

	regCenterServerUrl := G_config.Data.RegCenterServerUrl
	regCenterServerPort := G_config.Data.RegCenterServerPort
	regConfigCenterNamespaceId := G_config.Data.RegConfigCenterNamespaceId
	regConfigCenterGroupId := G_config.Data.RegConfigCenterGroupId

	nacosLocalHostAddress := G_config.Data.NacosLocalHostAddress
	nacoslocalHostPort := G_config.Data.Grpc_port

	// 注册到consul
	reqServer := tsNacos.NacosRegNameServerReq{
		ServerUrl:   regCenterServerUrl,
		ServerPort:  regCenterServerPort,
		ServiceName: tsGrpc.GameMicroGrpcVediosInfo,
		NamespaceId: regConfigCenterNamespaceId,
		GroupName:   regConfigCenterGroupId,
		Ip:          nacosLocalHostAddress,
		Port:        nacoslocalHostPort,
		RotateTime:  "24h",
		LogLevel:    "info",
	}
	err := tsNacos.RegisterNameService(reqServer)
	if err != nil {
		logs.Error("注册成功consul失败: ", err)
		panic(err)
	}

	// 启动RPC服务
	port := G_config.Data.Grpc_port
	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logs.Error("启动RPC服务失败: ", err)
		panic(err)
	}

	logs.Info("ConsulGrpcRegServer 启动RPC服务成功")

	srv := grpc.NewServer()
	heathcheck.RegisterHealthServer(srv, &controllers.HeathServerController{}) // 注册health_check服务
	vedios_info.RegisterVediosInfoServiceServer(srv, &controllers.GrpcVediosInfoServerController{})
	if err := srv.Serve(lis); err != nil {
		logs.Error("RPC服务绑定业务失败: ", err)
		panic(err)
	}

	logs.Info("ConsulGrpcRegServer RPC服务绑定业务成功")

}

func initGame() {
	initGameConfig()
	initGameLog()
	initGameDb()

	G_vedios = &GameVediosInfoManager{}
	G_vedios.init()

	if G_config.Data.Grpcflag {
		go initgrpc()
	}
	initGameRouter()

	a := map[string]interface{}{}
	aa, _ := json.Marshal(a)
	G_EMPTY_JSON = string(aa)

}

func startGame() {
	_router.Run(G_config.Data.Svr_port)
}

func startTimer() {
	for {
		time.Sleep(time.Second * 1)
		G_vedios.update(1)
	}
}

func getExitSignal() {
	c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, os.Kill)

	s := <-c
	fmt.Println("Got signal:", s)
}

func Start() {

	initGame()
	go startTimer()
	go getExitSignal()
	startGame()
}
