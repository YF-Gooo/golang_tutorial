package main

import (
	"fmt"
	"go_socket/protocol"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn) {
	// for i := 0; i < 1000; i++ {
	// 	words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
	// 	conn.Write(protocol.Packet([]byte(words)))
	// }
	words := "2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。 2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。[56]  3月31日凌晨，文章通过新浪微博发布声明承认与姚笛的婚外情，并向家人诚恳道歉，称辜负了家庭 [57]  ；马伊琍则发微博回应：“恋爱虽易，婚姻不易，且行且珍惜” [58]  ；然而，文章声明里的4个“辜负”，相同的字繁简体不同，被质疑是代笔 [59-60]  ；同一天晚上，《南都娱乐周刊》执行主编谢晓通过微访谈在线与网友交流了他们追踪报道文章婚外情的幕后情况 [61]  ；当晚10点，文章再次发微博，质问《南都娱乐周刊》主编谢晓和出品人陈朝华。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。 [56]  3月31日凌晨，文章通过新浪微博发布声明承认与姚笛的婚外情，并向家人诚恳道歉，称辜负了家庭 [57]  ；马伊琍则发微博回应：“恋爱虽易，婚姻不易，且行且珍惜” [58]  ；然而，文章声明里的4个“辜负”，相同的字繁简体不同，被质疑是代笔 [59-60]  ；同一天晚上，《南都娱乐周刊》执行主编谢晓通过微访谈在线与网友交流了他们追踪报道文章婚外情的幕后情况 [61]  ；当晚10点，文章再次发微博，质问《南都娱乐周刊》主编谢晓和出品人陈朝华。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。 [56]  3月31日凌晨，文章通过新浪微博发布声明承认与姚笛的婚外情，并向家人诚恳道歉，称辜负了家庭 [57]  ；马伊琍则发微博回应：“恋爱虽易，婚姻不易，且行且珍惜” [58]  ；然而，文章声明里的4个“辜负”，相同的字繁简体不同，被质疑是代笔 [59-60]  ；同一天晚上，《南都娱乐周刊》执行主编谢晓通过微访谈在线与网友交流了他们追踪报道文章婚外情的幕后情况 [61]  ；当晚10点，文章再次发微博，质问《南都娱乐周刊》主编谢晓和出品人陈朝华。2014年3月28日，有媒体拍到了文章和姚笛在一起的照片。 [56]  3月31日凌晨，文章通过新浪微博发布声明承认与姚笛的婚外情，并向家人诚恳道歉，称辜负了家庭 [57]  ；马伊琍则发微博回应：“恋爱虽易，婚姻不易，且行且珍惜” [58]  ；然而，文章声明里的4个“辜负”，相同的字繁简体不同，被质疑是代笔 [59-60]  ；同一天晚上，《南都娱乐周刊》执行主编谢晓通过微访谈在线与网友交流了他们追踪报道文章婚外情的幕后情况 [61]  ；当晚10点，文章再次发微博，质问《南都娱乐周刊》主编谢晓和出品人陈朝华。"
	// words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
	conn.Write(protocol.Packet([]byte(words)))
	fmt.Println("send over")
}

func main() {
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")
	go sender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}
