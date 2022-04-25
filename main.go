package main

import (
	"fmt"
	"strconv"
	"time"
)
import (
	"github.com/freetsdb/htcp/client"
	"github.com/freetsdb/htcp/server"
)

func main() {
	{
		s, err := server.NewServer(
			server.WithBindPort(3333),
			server.WithClientConnectObserves(func(c *client.Client) {
				c.Send([]byte(strconv.FormatUint(uint64(time.Now().UnixMicro()), 16)))
			}),
		)
		if err != nil {
			panic(err)
		}
		defer s.Close()
	}

	{
		for i := 0; i < 3; i++ {
			go func(p int) {
				c, err := client.NewClient(
					client.WithConnectAddr(":3333"),
					client.WithClientConnectObserves(func(c *client.Client) {
						fmt.Printf("连接成功, client:%s, server:%s\n", c.LocalAddr(), c.RemoteAddr())
					}),
					client.WithClientGetDataObserves(func(c *client.Client, data []byte) {
						fmt.Printf("client:%d 收到数据:%s\n", c.GetId(), string(data))
					}),
					client.WithClientCloseObserves(func(c *client.Client, err error) {
						fmt.Printf("client:%d 连接关闭", err)
					}),
				)
				if err != nil {
					panic(err)
				}
				defer c.Close()
			}(i)
		}
	}

	time.Sleep(5 * time.Second)
}
