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
> ## ✨ 功能特性
### 🛒 购物车服务
- 支持添加、删除、修改商品数量，结合库存预校验，防止超卖
- 购物车数据缓存至 Redis，减少数据库读压力，提升响应速度
### 📦 订单中心
- 提供订单创建、查询、取消等核心操作，核心流程在事务中完成，保障数据一致性
- 内置订单状态机，严格定义状态流转路径：**待支付 → 已支付 → 已发货 → 已完成 / 已取消**
- 通过状态前置校验，拦截非法状态变更，防止订单数据污染
### 💰 支付服务
- 对接第三方支付渠道，实现支付下单、异步通知回调处理
- **幂等性设计**：以 `(订单ID + 支付流水号)` 为幂等键，结合 Redis 记录回调处理状态，有效防止重复扣款与订单状态异常
- 在支付回调中触发订单状态机推进，保证支付与订单状态实时同步
### 🏠 地址管理
- 支持地址增删改查，并通过数据库事务实现**默认地址唯一性控制**
- 采用 `SELECT ... FOR UPDATE` 解决并发修改场景下可能出现多个默认地址的数据一致性问题
### 🔐 RBAC 权限管理（微服务）
- 支持角色、菜单、权限的灵活配置，实现接口级权限控制
- 独立微服务部署，通过 gRPC 对外提供权限鉴权接口，与业务服务解耦
### 🤖 验证码服务（微服务）
- 生成图形验证码并存储于 Redis，设定有效期，支持验证码校验
- 独立服务部署，gRPC 调用，可横向扩展
### ⚡ 高性能与高可用
- **Redis 缓存**：缓存 Session 登录态及商品热点数据，采用 Cache-Aside 模式，防止缓存穿透
- **gRPC 通信**：服务间使用 gRPC + Protobuf 强类型契约，高效且文档化
- **容器化部署**：提供 Dockerfile 与 docker-compose，一键启动所有服务
### 🧩 系统架构 
<img width="1268" height="1295" alt="micro_framework01" src="https://github.com/user-attachments/assets/bc7f5aa3-2b51-4f46-a110-672a29fc6983" />
<img width="1218" height="1295" alt="micro_framework02" src="https://github.com/user-attachments/assets/e81b313f-762e-48be-817d-32d78831b457" />

