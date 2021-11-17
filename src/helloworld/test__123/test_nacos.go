package test__123

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

/*
//请求配置
//@tenant 为命名空间id
http://10.2.2.99:8848/nacos/v1/cs/configs?dataId=nacos_test_dataid_01&group=nacos_test_group_01&tenant=a2ca0c29-b7cc-4b56-bbcb-51e9f813526d


*/
/*
//nacos-sdk-go
https://github.com/zensh/nacos-sdk-go/blob/master/README_CN.md
	可以在 Go 程序中使用 Nacos Go SDK 管理 Nacos 配置，包括获取、监听、发布和删除配置。

示例
https://github.com/nacos-group/nacos-sdk-go/tree/master/example
*/
func Test_nacos() {
	normal()
}
func normal() {
	var namespaceId = "a2ca0c29-b7cc-4b56-bbcb-51e9f813526d"
	//创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	////创建clientConfig的另一种方式
	//clientConfig := *constant.NewClientConfig(
	//    constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"),
	//    constant.WithTimeoutMs(5000),
	//    constant.WithNotLoadCacheAtStart(true),
	//    constant.WithLogDir("/tmp/nacos/log"),
	//    constant.WithCacheDir("/tmp/nacos/cache"),
	//    constant.WithRotateTime("1h"),
	//    constant.WithMaxAge(3),
	//    constant.WithLogLevel("debug"),
	//)

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "10.2.2.99",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	//创建serverConfig的另一种方式
	//serverConfigs := []constant.ServerConfig{
	//    *constant.NewServerConfig(
	//        "console1.nacos.io",
	//        80,
	//        constant.WithScheme("http"),
	//        constant.WithContextPath("/nacos")
	//    ),
	//    *constant.NewServerConfig(
	//        "console2.nacos.io",
	//        80,
	//        constant.WithScheme("http"),
	//        constant.WithContextPath("/nacos")
	//    ),
	//}

	// 创建动态配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "com.alibaba.nacos.example.properties"
	//var group = "nacos_test_group_01"
	var group = "nacos_test_group_02"

	// 发布配置
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: "connectTimeoutInMills=30000000"})

	if success {
		fmt.Println("Publish config successfully.")
	} else {
		errstr := err.Error()
		fmt.Println("err:", errstr)
	}

	// 创建服务发现客户端
	//_, _ := clients.CreateNamingClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})

	//// 创建动态配置客户端
	//_, _ := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})

	//// 创建服务发现客户端的另一种方式 (推荐)
	//namingClient, err := clients.NewNamingClient(
	//	vo.NacosClientParam{
	//		ClientConfig:  &clientConfig,
	//		ServerConfigs: serverConfigs,
	//	},
	//)

	//// 创建动态配置客户端的另一种方式 (推荐)
	//configClient, err := clients.NewConfigClient(
	//	vo.NacosClientParam{
	//		ClientConfig:  &clientConfig,
	//		ServerConfigs: serverConfigs,
	//	},
	//)
}

func normal_() {
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	//var endpoint = "${endpoint}"
	var endpoint = "10.2.2.99"

	//var namespaceId = "${namespaceId}"
	var namespaceId = "a2ca0c29-b7cc-4b56-bbcb-51e9f813526d"

	//1. 推荐使用 RAM 用户的 accessKey、secretKey
	//2. nacos的启动配置文件里面可以配密钥，配了之后向nacos注册之类的请求的时候都要带上密钥
	//var accessKey = "${accessKey}"
	//var secretKey = "${secretKey}"
	//var accessKey = ""
	//var secretKey = ""

	clientConfig := constant.ClientConfig{
		//
		Endpoint:    endpoint + ":8848",
		NamespaceId: namespaceId,
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "com.alibaba.nacos.example.properties"
	var group = "nacos_test_group_01"

	// 发布配置
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: "connectTimeoutInMills=30000000"})

	if success {
		fmt.Println("Publish config successfully.")
	} else {
		errstr := err.Error()
		fmt.Println("err:", errstr)
	}

	time.Sleep(3 * time.Second)

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	fmt.Println("Get config：" + content)

	// 监听配置
	configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	// 删除配置
	success, err = configClient.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if success {
		fmt.Println("Delete config successfully.")
	}
}
