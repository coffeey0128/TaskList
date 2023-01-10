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
`make testf`: 重置假資料並執行測試
   
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

# Code Gen
## 執行方式
1. 先建立好db schema
2. `make gen-crud`

<img width="810" alt="截圖 2023-01-11 上午4 10 58" src="https://user-images.githubusercontent.com/70966646/211652519-e60e1536-6bdf-4e63-af4b-d0eb047505d9.png">

## 說明
- 會去根據輸入的table name去抓取schema產生程式碼，包含api,service,repo及各自的測試範本
- 因為有使用到mock及wire，生成兩個檔案需要的程式後，會接著去下指令建立需要的檔案
- 測試需要自行補充細節，只生成起最基本能通過的測試
- 本專案使用該功能生成，並針對需求進行改寫
  
# Project Structure
```
├── api                   // controller
├── api_test              // api 整合測試
├── cmd
│   └──seeder             // 新增測試資料的程式
├── code_gen              // 自動生成程式碼的程式(make gen-crud)
├── config                
├── docs                  // swagger file(swag init 自動生成)
├── driver                // 初始化資料庫
├── internal              
│   └── task              // 實作與資料庫交互邏輯與單元測試，實作業務邏輯與單元測試
├── middleware            // 中介層
├── migration             // 資料庫初始化腳本(手動migrate-up)
├── mock                  // mock file(mockgen 自動生成)
├── models                // 定義 db schema 及用來操作 db 交互的 struct
│   ├── apireq            // 定義 API request struct
│   └── apires            // 定義 API response struct
├── pkg          
│   ├── errors            // 定義error struct
│   ├── query_condition   // 解析query，便於前端帶條件參數(本次未用到)
│   └── seeds             // 假資料設定
└── route                 // api 路由