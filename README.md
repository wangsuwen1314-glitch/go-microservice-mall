# go-microservice-mall
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Gin](https://img.shields.io/badge/Gin-1.9-blue?style=flat)
![gRPC](https://img.shields.io/badge/gRPC-✔-2ca5a0?style=flat)
![MySQL](https://img.shields.io/badge/MySQL-8.0-orange?style=flat&logo=mysql)
![Redis](https://img.shields.io/badge/Redis-7.0-DC382D?style=flat&logo=redis)
![Docker](https://img.shields.io/badge/Docker-✔-2496ED?style=flat&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green?style=flat)
>基于一个go语言微服务后端商城开发系统
## 项目简介
**go-microservice-mall** 是一个基于 Go 语言实现的微服务电商系统，涵盖购物车、订单、第三方支付、用户地址、验证码与 RBAC 权限管理等功能模块，服务间采用 gRPC + Protobuf 通信。
项目以真实电商交易链路为核心，实现了订单状态机流转、支付幂等回调、默认地址事务控制等关键业务逻辑，并基于 Redis 缓存 Session 与热点数据，降低数据库压力。系统支持 Docker 容器化部署，各微服务可独立扩展。
>- 支付幂等设计：解决第三方支付重复通知导致的重复扣款风险。
>- 默认地址唯一性控制：防止并发设置多个默认地址，保证用户地址数据的一致性。
>- 状态机流转：规范订单生命周期，避免非法状态转换。
> 核心设计目标：保障交易一致性、解决并发数据冲突、实现服务高可用。
> 
