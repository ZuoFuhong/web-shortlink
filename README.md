## Web-shortlink

### 设计思路

1. 通过 MurmurHash 哈希算法生成短网址，需要关心哈希算法的计算速度和冲突概率
2. 为了让短网址更短，可以将 10 进制的哈希值转换成 62 进制的哈希值
3. 解决哈希冲突，可以搜索数据库发现冲突，修改 Murmurhash 种子，重新计算哈希值。

### 接口文档

**1. 生成短地址**

```shell
curl --location --request POST 'http://127.0.0.1:8080/api/v1/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://www.google.com/",
    "expiration_in_minutes": 1
}'
```

**2. 短地址还原**

```shell
curl --location --request GET 'http://127.0.0.1:8080/api/v1/short_info/1tJiBJ'
```

**3. 访问短地址**

```shell
curl --location --request GET 'http://127.0.0.1:8080/1tJiBJ'
```

### 接口压测

```shell
wrk -t1 -c10 -d2 --script=post.lua --latency http://127.0.0.1:8080/api/v1/shorten
```

post.lua

```lua
wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json; charset=utf-8"
wrk.body = "{\"url\":\"https://www.google.com/\",\"expiration_in_minutes\":1}"
```

macOS 双核四线程连接本地 Redis 压测结果: 

```text
Running 2s test @ http://127.0.0.1:8080/api/v1/shorten
  1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.85ms    1.11ms  16.27ms   88.98%
    Req/Sec     5.60k     1.20k    6.88k    71.43%
  Latency Distribution
     50%    1.59ms
     75%    2.14ms
     90%    2.86ms
     99%    5.94ms
  11699 requests in 2.10s, 2.54MB read
Requests/sec:   5568.57
Transfer/sec:      1.21MB
```

### 参考资料

1. [Design Pastebin.com (or Bit.ly)](https://github.com/ZuoFuhong/system-design-primer/blob/master/solutions/system_design/pastebin/README.md)
2. [硬核课堂-短网址系统设计](https://www.bilibili.com/video/BV1UB4y1s7iZ)