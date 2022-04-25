# htcp
###### 轻量级的socket封装, 不粘包, 心跳检测, 并发安全



## 单方面传递心跳的弊端
若server收不到node1的心跳，泽说明node1失去了联系，但是并不一定是出现故障，也有可能出现node1服务处于繁忙状态，导致心跳传输超时。也有可能是server于node1之间的网络链路出现故障或者闪断，所以这种单方面传递的心跳不是万能的。

## 解决方案
使用周期性检测心跳机制：server每隔s秒向各个node发送检测请求，设定一个超时时间，如果超过超时时间，则进入死亡列表。
累计失效检测机制：在1 的基础之上，统计一定周期内节点的返回情况，以此来计算节点的死亡概率（超过超时次数/总检测次数）。对于死亡列表中的节点发起有限次数的重试，来做进一步判断。
对于设定的概率进行比对如果达到设定的概率可以进行一个真实踢出局的操作

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
