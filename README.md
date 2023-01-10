# TaskList
gogolook coding exercise



## 執行

兩種方式運行server

1. 手動執行
    1. 執行 `make init` 建立db和schema，並塞入假資料 (需要Mysql)
    2. 執行 `make run`
 
2. 執行 `docker-compose up -d` 並等待架設，再下`make init`

server運行在`localhost:8000`
   1. 可直接下`curl localhost:8000/{APIs}` 
   2. 或是使用swagger操作 http://localhost:8000/swagger/index.html#/ 

## 測試
`testf`: 重置假資料並執行測試
   
# Technologies
- Languange: GO 1.19
- Web Framework: [gin 1.8.2](https://github.com/gin-gonic/gin)
- Object Relational Mapping: [gorm 1.24.3](https://github.com/go-gorm/gorm)
- Dependency Injection: [wire 0.5.0](https://github.com/google/wire)
- API Doc: [swag 1.8.9](https://github.com/swaggo/swag)
- Mocking Framework: [gomock 1.6.0](https://github.com/golang/mock)

# API Routes
| Method | Path | Description |
| :----: | :--: | :---------- |
| GET | /tasks | 取得全部tasks|
| POST | /task | 新增task |
| PUT | /task/{id} | 更新該id的task |
| DELETE | /task/{id} | 刪除該id的task|