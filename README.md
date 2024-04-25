# WebSocketDemo

## 项目简介

利用 Gin+WebSocket 实现的在线聊天室 Demo 项目，支持加入/离开聊天室广播、给其他用户发送消息等。

## 如何使用

进入项目根目录，执行命令`go run .`命令，结果如下：

![image-20240425142642036](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425142642036.png)

可以看到们的 HTTP 服务已经启动成功并运行在了`8080`端口上。

接下来打开接口调试软件`Apifox`，也可以使用[在线的 WebSocket 接口调试网站](https://tool.xroy.net/websocket/)。

进入`Apifox`，新建`WebSocket`接口：

![image-20240425143032300](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425143032300.png)

输入`query`参数`user`，这个参数为必选，表示是谁加入了聊天室。

![image-20240425144233434](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425144233434.png)

输入接口地址`ws://127.0.0.1:8080/chat`，点击连接：

![image-20240425144429944](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425144429944.png)

结果如下：

![image-20240425144844987](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425144844987.png)

新建多个`WebSocket`接口，可以看到一有新成员加入，其他成员都会收到广播通知：

![image-20240425144942473](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425144942473.png)

![image-20240425145118970](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425145118970.png)

![image-20240425145139218](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425145139218.png)

当有人离开聊天室时，其他人也会受到下线通知：

![20240425145248_rec_-convert](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/20240425145248_rec_-convert.gif)

接下来我们将该用户重新上线，同时演示如何给其他成员发消息，需要将发送的消息格式设置为`json`：

![image-20240425145621983](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425145621983.png)

消息的正确格式如下：

```json
{
    "toUser": "赵丽颖",
    "content": "你真棒"
}
```

`toUser`代表给谁发消息，`content`代表发消息的内容。这两者有任何一者为空，均会收到提示内容缺失的消息。

![image-20240425145916418](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425145916418.png)

点击发送，可以看到收信人成功收到消息：

![20240425150210_rec_-convert](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/20240425150210_rec_-convert.gif)

如果该接收人不在线，发信人也会收到相应的提示信息：

![image-20240425150453002](https://cdn.jsdelivr.net/gh/palp1tate/ImgPicGo/img/image-20240425150453002.png)

`WebSocketDemo`算是一个比较经典的用于学习`WebSocket`的实战项目了吧，所有的介绍就到这里，欢迎 star！
