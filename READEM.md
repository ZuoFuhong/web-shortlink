## 短地址服务

> 1.什么是短地址服务

将长地址缩短到一个很短的地址，用户访问这个短地址可以重新定位到原本的长地址。

![snap1](./doc/snap1.png)

> 2.短地址的使用场景

![snap2](./doc/snap2.png)

> 3.数据分析报表

![snap3](./doc/snap3.png)
![snap4](./doc/snap4.png)

示例图：

![snap5](./doc/snap5.png)

### 接口文档

1.生成短地址接口
  > POST /api/shorten
  Content-Type: application/json;charset=utf-8
  Params：
    {
    	"url": "https://www.baidu.com",
    	"expiration_in_minutes": 100
    }
  Response：
    {
        "code": "0",
        "errMsg": "OK",
        "data": {
            "shortUrl": "8dxu",
            "longUrl": "https://www.baidu.com"
        }
    }
  
2.短地址还原接口
  > GET /api/info?shortUrl=8dxu
  Response：
    {
        "code": "0",
        "errMsg": "OK",
        "data": {
            "url": "https://www.baidu.com",
            "created_at": "2020-02-24 12:21:02.151994 +0800 CST m=+94.081029585",
            "expiration_in_minutes": 100
        }
    }

3.短地址访问-重定向（307）
  > GET /8dxu


### RoadMap

1.接口调用-Token验证

2.流量桶算法

3.采集数据（客户端IP、设备类型、网络运行商等）

4.绘制统计图表