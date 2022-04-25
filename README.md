# htcp
###### 轻量级的socket封装, 不粘包, 心跳检测, 并发安全

## 获得 htcp
`go get -u github.com/freetsdb/htcp`

## 使用 htcp

```go
package main

import (
    "fmt"
    "github.com/freetsdb/htcp/client"
    "github.com/freetsdb/htcp/server"
    "time"
)

func main() {
    {
        s, err := server.NewServer(
            server.WithBindPort(3333),
            server.WithClientConnectObserves(func(c *client.Client) {
                c.Send([]byte("你好"))
            }),
        )
        if err != nil {
            panic(err)
        }
        defer s.Close()
    }

    {
        c, err := client.NewClient(
            client.WithConnectAddr(":3333"),
            client.WithClientConnectObserves(func(c *client.Client) {
                fmt.Println("连接成功", c.LocalAddr(), c.RemoteAddr())
            }),
            client.WithClientGetDataObserves(func(c *client.Client, data []byte) {
                fmt.Println("收到数据", string(data))
            }),
            client.WithClientCloseObserves(func(c *client.Client, err error) {
                fmt.Println("连接关闭", err)
            }),
        )
        if err != nil {
            panic(err)
        }
        defer c.Close()
    }

    time.Sleep(1e9)
}

```
