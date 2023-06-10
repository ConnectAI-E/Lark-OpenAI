<p align='center'>
    <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/d855be1b-7d55-4b6f-aac1-2c7da5ab784c' alt='' width='800'/>
</p>

<details align='center'>
    <summary> 📷 點擊展開企聯AI完整功能</summary>
    <br>
    <p align='center'>
        <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/b3543ae5-23cc-4fd7-829b-74656a300901' width='800'/>
        <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/948b051a-c25a-4825-8a18-5b72f6660b87' width='800'/>
        <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/a2139cc3-021e-4820-8e95-53a70541f136' width='800'/>
        <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/8e700709-ec0d-4c5e-b4af-710e42a0ba69' width='800'/>
        <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/001cb314-5fe2-4c93-86a7-66ea747b7855' width='800'/>
        <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/4dc12ca6-5305-4bc5-bf1a-40354c17ef29' width='800'/>
        <img src='https://github.com/ConnectAI-E/Feishu-OpenAI/assets/110169811/0c3178a0-e648-4013-a042-c1f83dbbe847' width='800'/>   
    </p>
</details>

<br>

<p align='center'>
   飛書 ×（GPT-4 + DALL·E + Whisper）
<br>
<br>
    🚀 Lark OpenAI 🚀
</p>

<p align='center'>
   www.qiniai.com
</p>

<strong align="center">
<samp>

[**English**](./readme.md) · [**简体中文**](https://github.com/ConnectAI-E/Feishu-OpenAI)· [**繁體中文**](./readme_zh-hk.md) · [**日本語**](./readme_jp.md) · [**Tiếng Việt**](./readme_vi.md)

</samp>
</strong>


## 👻 機器人功能

🗣 語音交流：私人直接與機器人暢所欲言

💬 多話題對話：支持私人和群聊多話題討論，高效連貫

🖼 文本成圖：支持文本成圖和以圖搜圖

🛖 場景預設：內建豐富場景列表，一鍵切換AI角色

🎭 角色扮演：支持場景模式，增添討論樂趣和創意

🤖 AI模式：內建4種AI模式，感受AI的智慧與創意

🔄 上下文保留：回復對話框即可繼續同一話題討論

⏰ 自動結束：超時自動結束對話，支持清除討論歷史

📝 富文本卡片：支持富文本卡片回覆，資訊更豐富多彩

👍 交互式反饋：即時獲取機器人處理結果

🎰 餘額查詢：即時獲取token消耗情況

🔙 歷史回滾：輕鬆回滾歷史對話，繼續話題討論 🚧

🔒 管理員模式：內建管理員模式，使用更安全可靠 🚧

🌐 多token負載均衡：優化生產級別的高頻調用場景

↩️ 支持反向代理：為不同地區的使用者提供更快、更穩定的訪問體驗

📚 與飛書文檔互動：成為企業員工的超級助手 🚧

🎥 話題內容秒轉PPT：讓你的匯報從此變得更加簡單 🚧

📊 表格分析：輕鬆導入飛書表格，提升資料分析效率 🚧

🍊 私有數據訓練：利用公司產品信息對GPT二次訓練，更好地滿足客戶個性化需求 🚧



## 🌟 項目特點

- 🍏 對話基於 OpenAI-[gpt-4](https://platform.openai.com/account/api-keys) 和 [Lark](https://www.larksuite.com/)
- 🥒 支持[Serverless 雲函數](https://github.com/serverless-devs/serverless-devs)、[本地環境](https://dashboard.cpolar.com/login)、[Docker](https://www.docker.com/)、[二進制安裝包](https://github.com/Leizhenpeng/feishu-chatgpt/releases/)

## 項目部署

<details>
    <summary>本地部署</summary>
<br>

```bash
git clone git@github.com:ConnectAI-E/lark-openai.git
cd Lark-OpenAI/code
```
如果你的伺服器沒有公網IP，可以使用反向代理的方式。

飛書的伺服器在國內對ngrok的訪問速度很慢，所以推薦使用一些國內的反向代理服務商：

- [cpolar](https://dashboard.cpolar.com/)
- [natapp](https://natapp.cn/)

```bash
# 配置config.yaml
mv config.example.yaml config.yaml

// 測試部署
go run ./
cpolar http 9000

// 正式部署
nohup cpolar http 9000 -log=stdout &

// 查看伺服器狀態
https://dashboard.cpolar.com/status

// 下線服務
ps -ef | grep cpolar
kill -9 PID
```

<br>

</details>

<details>
    <summary>serverless雲函數部署</summary>
<br>

```bash
git clone git@github.com:ConnectAI/lark-openai.git
cd Lark-OpenAI/code
```

安裝[severless](https://docs.serverless-devs.com/serverless-devs/quick_start)工具

```bash
# 配置config.yaml
mv config.example.yaml config.yaml
# 安裝severless cli
npm install @serverless-devs/s -g
```

安裝完成後，請根據您本地環境，根據下面教程部署`severless`

- 本地 `linux`/`mac os` 環境

1. 修改`s.yaml`中的部署地區和部署秘鑰

```
edition: 1.0.0
name: feishuBot-chatGpt
access: "aliyun" #  修改自定義的秘鑰別稱

vars: # 全局變量
region: "cn-hongkong" # 修改雲函數想要部署地區

```

2. 一鍵部署

```bash
cd ..
s deploy
```

- 本地`windows`

1. 首先打開本地`cmd`命令提示符工具，運行`go env`檢查你電腦上 go 環境變量設置，確認以下變量和值

```cmd
set GO111MODULE=on
set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
```

如果值不正確，比如您電腦上為`set GOOS=windows`，請運行以下命令設置`GOOS`變量值。

```cmd
go env -w GOOS=linux
```

2. 修改`s.yaml`中的部署地区和部署秘钥

```
edition: 1.0.0
name: feishuBot-chatGpt
access: "aliyun" #  修改自訂的秘鑰別稱

vars: # 全局變量
  region: "cn-hongkong" #  修改雲函數想要部署地區

```

3. 修改`s.yaml`中的`pre-deploy`，去除第二步`run`前面的環境變量設置部分

```
  pre-deploy:
        - run: go mod tidy
          path: ./code
        - run: go build -o
            target/main main.go  # 刪除GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0
          path: ./code

```

4. 一键部署

```bash
cd ..
s deploy
```

<br>
</details>

<!-- <details>
    <summary>使用 Railway 平台一键部署</summary> -->

<!-- 
Railway 是一家國外的 Serverless 平台，支持多種語言，可以一鍵將 GitHub 上的代碼倉庫部署到 Railway 平台，然後在 Railway
平台上配置環境變量即可。部署本項目的流程如下：

#### 1. 生成 Railway 項目

點擊下方按鈕即可創建一個對應的 Railway 項目，其會自動 Fork 本項目到你的 GitHub 帳號下。

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/10D-TF?referralCode=oMcVS2)

#### 2. 配置環境變量

在打開的頁面中，配置環境變量，每個變量的說明如下圖所示：


<img src='https://user-images.githubusercontent.com/50035229/225005602-88d8678f-9d17-4dc5-8d1e-4abf64fb84fd.png' alt='Railway 環境變量' width='500px'/>

#### 3. 部署項目

填寫完環境變量後，點擊 Deploy 就完成了項目的部署。部署完成後還需獲取對應的域名用於飛書機器人訪問，如下圖所示：

<img src='https://user-images.githubusercontent.com/50035229/225006236-57cb3c8a-1b7d-4bfe-9c9b-099cb9179027.png' alt='Railway 域名' width='500px'/>

如果不確定自己部署是否成功，可以通過訪問上述獲取到的域名 (https://xxxxxxxx.railway.app/ping) 來查看是否返回了`pong`
，如果返回了`pong`，說明部署成功。

</details> -->

<details>
    <summary>docker部署</summary>
<br>

```bash
docker build -t feishu-chatgpt:latest .
docker run -d --name feishu-chatgpt -p 9000:9000 \
--env APP_ID=xxx \
--env APP_SECRET=xxx \
--env APP_ENCRYPT_KEY=xxx \
--env APP_VERIFICATION_TOKEN=xxx \
--env BOT_NAME=chatGpt \
--env OPENAI_KEY="sk-xxx1,sk-xxx2,sk-xxx3" \
--env API_URL="https://api.openai.com" \
--env HTTP_PROXY="" \
feishu-chatgpt:latest
```

注意:

- `APP_LANG` 是Lark机器人的语言，例如，`en`,`ja`,`vi`,`zh-hk`.
- `BOT_NAME` 為飛書機器人名稱，例如 `chatGpt`
- `OPENAI_KEY` 為openai key，多個key用逗號分隔，例如 `sk-xxx1,sk-xxx2,sk-xxx3`
- `HTTP_PROXY` 為宿主機的proxy地址，例如 `http://host.docker.internal:7890`,沒有代理的話，可以不用設置
- `API_URL` 為OpenAI API 接口地址，例如 `https://api.openai.com`, 沒有反向代理的話，可以不用設置

--- 

部署azure版本

```bash
docker build -t feishu-chatgpt:latest .
docker run -d --name feishu-chatgpt -p 9000:9000 \
--env APP_ID=xxx \
--env APP_SECRET=xxx \
--env APP_ENCRYPT_KEY=xxx \
--env APP_VERIFICATION_TOKEN=xxx \
--env BOT_NAME=chatGpt \
--env AZURE_ON=true \
--env AZURE_API_VERSION=xxx \
--env AZURE_RESOURCE_NAME=xxx \
--env AZURE_DEPLOYMENT_NAME=xxx \
--env AZURE_OPENAI_TOKEN=xxx \
feishu-chatgpt:latest
```

注意:

- `APP_LANG` 是Lark机器人的语言，例如， `en`,`ja`, `vi`,`zh-hk`.
- `BOT_NAME` 為飛書機器人名稱，例如 `chatGpt`
- `AZURE_ON` 為是否使用azure ,請填寫 `true`
- `AZURE_API_VERSION` 為azure api版本 例如 `2023-03-15-preview`
- `AZURE_RESOURCE_NAME` 為azure 資源名稱 類似 `https://{AZURE_RESOURCE_NAME}.openai.azure.com`
- `AZURE_DEPLOYMENT_NAME` 為azure 部署名稱 類似 `https://{AZURE_RESOURCE_NAME}.openai.azure.com/deployments/{AZURE_DEPLOYMENT_NAME}/chat/completions`
- `AZURE_OPENAI_TOKEN` 為azure openai token

</details>
<details>
    <summary>docker-compose 部署</summary>
<br>

編輯 docker-compose.yaml，透過 environment 配置相應環境變數（或透過 volumes 挂載相應配置文件），然後運行下面的命令即可

```bash
# 構建鏡像
docker compose build

# 啟動服務
docker compose up -d

# 停止服務
docker compose down
```

事件回調地址: http://IP:9000/webhook/event
卡片回調地址: http://IP:9000/webhook/card

</details>

## 詳細配置步驟

<details align='left'>
    <summary> 📸 點擊展開飛書機器人配置的分步截圖指導</summary>
    <br>
    <p align='center'>
    <img src='https://user-images.githubusercontent.com/50035229/223943381-39e0466f-2a5e-472a-9863-94eafb5f17b0.png' alt='' width='800'/>
    <img src='https://user-images.githubusercontent.com/50035229/223943448-228de5cb-0929-4d80-8087-8d8624dd6ddf.png' alt='' width='800'/>
    <img src='https://user-images.githubusercontent.com/50035229/223943485-ef331784-7940-4657-b128-70c98391e72f.png' alt='' width='800'/>
    <img src='https://user-images.githubusercontent.com/50035229/223943527-60e6653a-eb6e-4062-a076-b6c9da934352.png' alt='' width='800'/>
    <img src='https://user-images.githubusercontent.com/50035229/223943972-f49adf9f-af5f-463a-8c7a-c1f0cac0e8c3.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223944060-7ef630a4-4248-4509-852b-cad8bfffeefc.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223944230-aff586be-31cc-40de-9b1a-7d4e259d54dd.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223944350-917d115c-6c82-4d8b-9ec8-b5c82331a2dc.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223944381-97396156-f5e2-467f-aaf6-b1f6e1c446b2.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/230003546-36450f2f-b6e9-4292-8b40-3a4aa8a05a64.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223945122-f7ab3d9a-6742-43d2-970e-ddb0f284c7fa.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223944507-8d1a08d7-8b5b-4f32-a90d-fd338164ec82.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223944515-fb505e84-c840-484a-8df5-612f60bf27ea.png' alt='' width='800'/>
      <img src='https://user-images.githubusercontent.com/50035229/223944590-ad61320f-c14a-4542-80ad-dee2e6469b67.png' alt='' width='800'/>
    </p>
</details>


- 獲取 [OpenAI](https://platform.openai.com/account/api-keys) 的 KEY( 🙉 下面有免費的 KEY 供大家測試部署 )
- 創建 [飛書](https://open.larksuit.com/) 機器人
    1. 前往[開發者平台](https://open.feishu.cn/app?lang=zh-CN)創建應用，並獲取到 APPID 和 Secret
    2. 前往`應用功能-機器人`，創建機器人
    3. 從 cpolar、serverless 或 Railway 獲得公網地址，在飛書機器人後台的 `事件訂閱` 板塊填寫。例如，
        - `http://xxxx.r6.cpolar.top`為 cpolar 暴露的公網地址
        - `/webhook/event`為統一的應用路由
        - 最終的回調地址為 `http://xxxx.r6.cpolar.top/webhook/event`
    4. 在飛書機器人後台的 `機器人` 板塊，填寫消息卡片請求網址。例如，
        - `http://xxxx.r6.cpolar.top`為 cpolar 暴露的公網地址
        - `/webhook/card`為統一的應用路由
        - 最終的消息卡片請求網址為 `http://xxxx.r6.cpolar.top/webhook/card`
    5. 在事件訂閱板塊，搜索三個詞`機器人進群`、 `接收消息`、 `消息已讀`，把他們後面所有的權限全部勾選。
       進入權限管理界面，搜索`圖片`，勾選`獲取與上傳圖片或文件資源`。
       最終會添加下列回調事件
        - im:resource(獲取與上傳圖片或文件資源)
        - im:message
        - im:message.group_at_msg(獲取群組中所有消息)
        - im:message.group_at_msg:readonly(接收群聊中@機器人消息事件)
        - im:message.p2p_msg(獲取用戶發給機器人的單聊消息)
        - im:message.p2p_msg:readonly(讀取用戶發給機器人的單聊消息)
        - im:message:send_as_bot(獲取用戶在群組中@機器人的消息)
        - im:chat:readonly(獲取群組信息)
        - im:chat(獲取與更新群組信息)

## 企聯AI


|       <div style="width:300px">AI</div>        |                            <img width=100> SDK <img width=100>                             |                         Application                          |
| :---------------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
|      🎒OpenAI      |    [Go-OpenAI](https://github.com/ConnectAI-E/Go-OpenAI)     | [🏅Feishu-OpenAI](https://github.com/ConnectAI-E/Feishu-OpenAI), [Lark-OpenAI](https://github.com/ConnectAI-E/Lark-OpenAI), [🏅Feishu-EX-ChatGPT](https://github.com/ConnectAI-E/Feishu-EX-ChatGPT), [Feishu-OpenAI-Stream-Chatbot](https://github.com/ConnectAI-E/Feishu-OpenAI-Stream-Chatbot), [Feishu-OpenAI-Amazing](https://github.com/ConnectAI-E/Feishu-OpenAI-Amazing), [Feishu-Oral-Friend](https://github.com/ConnectAI-E/Feishu-Oral-Friend), [Feishu-OpenAI-Base-Helper](https://github.com/ConnectAI-E/Feishu-OpenAI-Base-Helper), [Feishu-Vector-Knowledge-Management](https://github.com/ConnectAI-E/Feishu-Vector-Knowledge-Management), [Feishu-OpenAI-PDF-Helper](https://github.com/ConnectAI-E/Feishu-OpenAI-PDF-Helper), [🏅Dingtalk-OpenAI](https://github.com/ConnectAI-E/Dingtalk-OpenAI), [Wework-OpenAI](https://github.com/ConnectAI-E/Wework-OpenAI), [WeWork-OpenAI-Node](https://github.com/ConnectAI-E/WeWork-OpenAI-Node), [llmplugin](https://github.com/ConnectAI-E/llmplugin) |
|  🎭 Stablediffusion  |                            ------                            | [🏅Feishu-Stablediffusion](https://github.com/ConnectAI-E/Feishu-Stablediffusion) |
|   🍎 Midjourney    | [Go-Midjourney](https://github.com/ConnectAI-E/Feishu-Midjourney/tree/main/midjourney) | [🏅Feishu-Midjourney](https://github.com/ConnectAI-E/Feishu-Midjourney), [MidJourney-Web](https://github.com/ConnectAI-E/MidJourney-Web), [Dingtalk-Midjourney](https://github.com/ConnectAI-E/Dingtalk-Midjourney) |
|    🍍 文心一言     |    [Go-Wenxin](https://github.com/ConnectAI-E/Go-Wenxin)     | [Feishu-Wenxin](https://github.com/ConnectAI-E/Feishu-Wenxin), [Dingtalk-Wenxin](https://github.com/ConnectAI-E/Dingtalk-Wenxin), [Wework-Wenxin](https://github.com/ConnectAI-E/Wework-Wenxin) |
|     💸 Minimax     |   [Go-Minimax](https://github.com/ConnectAI-E/Go-Minimax)    | [Feishu-Minimax](https://github.com/ConnectAI-E/Feishu-Minimax), [Dingtalk-Minimax](https://github.com/ConnectAI-E/Dingtalk-Minimax), [Wework-Minimax](https://github.com/ConnectAI-E/Wework-Minimax) |
|     ⛳️ CLAUDE      |    [Go-Claude](https://github.com/ConnectAI-E/Go-Claude)     | [Feishu-Claude](https://github.com/ConnectAI-E/Feishu-Claude), [DingTalk-Claude](https://github.com/ConnectAI-E/DingTalk-Claude), [Wework-Claude](https://github.com/ConnectAI-E/Wework-Claude) |
|     🎡 Prompt      |                                                              | [Prompt-Engineering-Tutior](https://github.com/ConnectAI-E/Prompt-Engineering-Tutior) |
|     🤖️ ChatGLM     |                                                              | [Feishu-ChatGLM](https://github.com/ConnectAI-E/Feishu-ChatGLM) |


