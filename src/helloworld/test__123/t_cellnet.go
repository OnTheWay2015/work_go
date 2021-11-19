package test__123

import (
	"fmt"
	test "helloworld/pbtest"
	"reflect"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/cellnet/util"
)

type TestEchoACK struct {
	Msg   string
	Value int32
}

func regmsg() {
	cc := codec.MustGetCodec("binary")
	tt := reflect.TypeOf((*TestEchoACK)(nil)).Elem()
	ii := int(util.StringHash("TestEchoACK"))
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{Codec: cc, Type: tt, ID: ii})

}

const peerAddress = "127.0.0.1:17701"

// 服务器逻辑
func server_cellnet() {

	// 创建服务器的事件队列，所有的消息，事件都会被投入这个队列处理
	queue := cellnet.NewEventQueue()

	// 创建一个服务器的接受器(Acceptor)，接受客户端的连接
	peerIns := peer.NewGenericPeer("tcp.Acceptor", "server", peerAddress, queue)

	// 将接受器Peer与tcp.ltv的处理器绑定，并设置事件处理回调
	// tcp.ltv处理器负责处理消息收发，使用私有的封包格式以及日志，RPC等处理
	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {

		// 处理Peer收到的各种事件
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted: // 接受一个连接
			fmt.Println("server accepted")
		case *test.ContentACK: // 收到连接发送的消息

			fmt.Printf("server recv pb: %+v\n", msg)

		case *TestEchoACK: // 收到连接发送的消息

			fmt.Printf("server recv %+v\n", msg)

			// 发送回应消息
			ev.Session().Send(&TestEchoACK{
				Msg:   msg.Msg,
				Value: msg.Value,
			})

		case *cellnet.SessionClosed: // 会话连接断开
			fmt.Println("session closed: ", ev.Session().ID())
		}

	})

	// 启动Peer，服务器开始侦听
	peerIns.Start()

	// 开启事件队列，开始处理事件，此函数不阻塞
	queue.StartLoop()
}

// 模拟客户端逻辑
func client() {

	// 例子专用的完成标记
	//done := make(chan struct{})

	// 创建客户端的事件处理队列
	queue := cellnet.NewEventQueue()

	// 创建客户端的连接器
	peerIns := peer.NewGenericPeer("tcp.Connector", "client", peerAddress, queue)

	// 将客户端连接器Peer与tcp.ltv处理器绑定，并设置接收事件回调
	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected: // 已经连接上
			fmt.Println("client connected")
			go testSend(ev)
		case *TestEchoACK: //收到服务器发送的消息

			fmt.Printf("client recv %+v\n", msg)

			// 完成操作
			//done <- struct{}{}
			//fmt.Printf("client recv  okkkkkk") //当前边有写 channel 时,不会走到这里. 会阻碍,因为要等 channel读取.

		case *cellnet.SessionClosed:
			fmt.Println("client closed")
		}
	})

	// 开启客户端Peer
	peerIns.Start()

	// 开启客户端队列处理
	queue.StartLoop()

	// 等待客户端收到消息
	//<-done

}

func testSend(ev cellnet.Event) {
	for true {
		time.Sleep(time.Microsecond * 30)
		ev.Session().Send(&TestEchoACK{
			Msg:   "hello",
			Value: 1234,
		})
		ev.Session().Send(&test.ContentACK{
			Msg:   "hellopb",
			Value: 4567,
		})
		time.Sleep(time.Second * 15)
	}
}

func client_main() {
	time.Sleep(time.Microsecond * 15)
	client()
}

func svr_main() {
	server_cellnet()
	for true {
		time.Sleep(time.Microsecond * 30)
	}
}

func Test__cellnet() {
	regmsg()
	go svr_main()
	go client_main()
	for true {
		time.Sleep(time.Microsecond * 30)
	}
}
