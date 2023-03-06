# README

## 關於

這是業績系統的後端相關API

---------------------------------------

## 技術棧
1. golang
2. go-admin
3. mysql
4. docker

## 目前有
1. User管理：新增帳號與帳號基本資料設定。
2. 部門管理：部門管理。
3. 職位管理：職位管理。
4. 選單管理：配置系統選單，操作權限，按鈕權限，API權限等。
5. 角色管理：角色選單權限分配。
6. 字典管理：對系統中經常使用的一些較為固定的資料進行維護。
7. 基本的 sales explorers

## 如何開始開發
1. 安裝 golang
2. clone repo
3. cd repo 
4. 設定好 config/settings.yml
5. make init && make dev


## 目錄結構
```base 
├── app                     # 主要的程式碼
│   ├── admin               # admin為後台的主要程式碼
│   │   ├── apis
│   │   ├── models
│   │   ├── router
│   │   └── service
│   │      └── dto
│   └── other               # 與後台不相關的API均放這，結構與admin一樣    
├── common                  # 共用func
├── config                  # 該project 所使用的config
├── docs                    # api doc 放置位置
├── scripts                 # 部署時使用著腳本
├── static                  # 透過api上傳檔案時所放的靜態目錄
├── temp                    # log file
├── template                
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── README.md
```