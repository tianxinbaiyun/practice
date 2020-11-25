@[TOC](influxdb)

## 简介

influxdb是一个开源分布式时序、时间和指标数据库，使用 Go 语言编写，无需外部依赖。其设计目标是实现分布式和水平伸缩扩展，是 InfluxData 的核心产品。
应用：性能监控，应用程序指标，物联网传感器数据和实时分析等的后端存储。
influxdb 完整的上下游产业还包括：Chronograf、Telegraf、Kapacitor

### 概念区分

|名称|描述|
|---|---|
|database|数据库|
|measurement|数据库中的表|
|point|表中的一行数据|



作者：楚_kw
链接：https://www.jianshu.com/p/f0905f36e9c3
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。