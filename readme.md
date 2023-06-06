<p align='center'>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/421327a5-1cf1-46ad-93b1-15c9c9d36490' alt='' width='800'/>
</p>


<details align='center'>
    <summary> 📷 Click to expand the full function of Connect-AI</summary>
    <br>
    <p align='center'>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/b993c610-1c91-40dd-bdcd-85a992c17b74' alt='语音对话' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/149f5fda-3fc4-49fa-8132-4825edfece1f' alt='角色扮演' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/7dae5661-2d4c-4584-934c-747a8c68d3e9' alt='角色扮演' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/942ffb30-fb48-4de4-a696-e0903a691665' alt='角色列表' width='800'/>
    </p>
</details>



<br>

<p align='center'>
   Lark ×（GPT-4 + DALL·E + Whisper）
<br>
<br>
    🚀 Lark OpenAI 🚀
</p>

<p align='center'>
   www.qiniai.com
</p>




<strong align="center">
<samp>

[**English**](./code/locales/en) · [**简体中文**](./code/locales/zh)· [**日本語**](./code/locales/ja)· [**Tiếng Việt**](./code/locales/vi)

</samp>
</strong>


## 👻 Feature

🗣Voice Communication: Private Direct Says with Robots

💬Multi-topic dialogue: support private and group chat multi-topic discussion, efficient and coherent

🖼Text graph: supports text graph and graph search

🛖Scene preset: built-in rich scene list, one-click switch AI role

🎭Role play: Support scene mode, add fun and creative discussion

🤖AI mode: Built-in 4 AI modes, feel the wisdom and creativity of AI

🔄Context preservation: reply dialog to continue the same topic discussion

⏰Automatic end: timeout automatically end the dialogue, support to clear the discussion history

📝Rich text card: support rich text card reply, more colorful information

👍Interactive Feedback: Instant access to robot processing results

🎰Balance query: obtain token consumption in real time

🔙History Back to File: Easily Back to File History Dialogue and Continue Topic Discussion🚧

🔒Administrator mode: built-in administrator mode, use more secure and reliable🚧

🌐Multi-token load balancing: Optimizing high-frequency call scenarios at the production level

↩️ Support reverse proxy: provide faster and more stable access experience for users in different regions

📚Interact with Flying Book Documents: Become a Super Assistant for Enterprise Employees🚧

🎥Topic Content Seconds to PPT: Make Your Report Simpler from Now on🚧

📊Table Analysis: Easily import flying book tables to improve data analysis efficiency🚧

🍊Private data training: use the company's product information for GPT secondary training to better meet the individual needs of customers.🚧



## 🌟 Base

- 🍏 The dialogue is based on OpenAI-[GPT4](https://platform.openai.com/account/api-keys) and [Lark](https://www.larksuite.com/)
- 🥒 support [Serverless ](https://github.com/serverless-devs/serverless-devs)、[local](https://dashboard.cpolar.com/login)、[Docker](https://www.docker.com/)、[binary package](https://github.com/Leizhenpeng/feishu-chatgpt/releases/)


## 🛵 Development

###### Description of configuration files for lark，**[➡︎ see more](#Detailed configuration steps)**

<details>
    <summary>Local Development</summary>
<br>

```bash
git clone git@github.com:ConnectAI-E/lark-openai.git
cd feishu-chatgpt/code
```

If your server does not have a public network IP, you can use a reverse proxy.

The server of Flying Book is very slow to access ngrok in China, so it is recommended to use some domestic reverse proxy service providers.

- [cpolar](https://dashboard.cpolar.com/)
- [natapp](https://natapp.cn/)

```bash
# 配置config.yaml
mv config.example.yaml config.yaml

//测试部署
go run main.go
cpolar http 9000

//正式部署
nohup cpolar http 9000 -log=stdout &

//查看服务器状态
https://dashboard.cpolar.com/status

// 下线服务
ps -ef | grep cpolar
kill -9 PID
```

<br>

</details>

<details>
    <summary>Serverless Development</summary>
<br>

```bash
git clone git@github.com:ConnectAI/lark-openai.git
cd feishu-chatgpt/code
```

install [severless](https://docs.serverless-devs.com/serverless-devs/quick_start)tool

```bash
# 配置config.yaml
mv config.example.yaml config.yaml
# 安装severless cli
npm install @serverless-devs/s -g
```

After the installation is complete, please deploy according to your local environment and the following tutorial`severless`

- local `linux`/`mac os` env

1. Modify the Deployment Region and Deployment Key in 's.yaml'

```
edition: 1.0.0
name: lark-openai
access: "aliyun" #  Modify the custom key name.

vars: # 全局变量
region: "cn-hongkong" # Modify the region where the cloud function wants to be deployed.

```

2. One-click deployment

```bash
cd ..
s deploy
```

- local `windows`

1. First open the local `cmd` command prompt tool, run `go env` to check the go environment variable settings on your computer, confirm the following variables and values

```cmd
set GO111MODULE=on
set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
```

If the value is incorrect, such as `set GOOS=windows` on your computer, please run the following command to set the `GOOS` variable value

```cmd
go env -w GOOS=linux
```

2. Modify the deployment region and deployment key in `s.yaml`

```
edition: 1.0.0
name: lark-openai
access: "aliyun" #  修改自定义的秘钥别称

vars: # 全局变量
  region: "cn-hongkong" #  修改云函数想要部署地区

```

3. Modify `pre-deploy` in `s.yaml`, remove the ring variable change part before the second step `run`

```
  pre-deploy:
        - run: go mod tidy
          path: ./code
        - run: go build -o
            target/main main.go  # 删除GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0
          path: ./code

```

4. One-click deployment

```bash
cd ..
s deploy
```

<br>
</details>

<details>
    <summary>Railway Deployment </summary>


Just configure environment variables on the platform. The process of deploying this project is as follows:

#### 1. Generate the Railway project

Click the button below to create a corresponding Railway project, which will automatically fork this project to your GitHub account.

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/10D-TF?referralCode=oMcVS2)

#### 2. Generate the Railway project

In the opened page, configure the environment variables. The description of each variable is shown in the figure below:

<img src='https://user-images.githubusercontent.com/50035229/225005602-88d8678f-9d17-4dc5-8d1e-4abf64fb84fd.png' alt='Railway 环境变量' width='500px'/>

#### 3. deployment project

After filling in the environment variables, click Deploy to complete the deployment of the project. After the deployment is complete, you need to obtain the corresponding domain name for the Feishu robot to access, as shown in the following figure:

<img src='https://user-images.githubusercontent.com/50035229/225006236-57cb3c8a-1b7d-4bfe-9c9b-099cb9179027.png' alt='Railway 域名' width='500px'/>

Uncertainty about success or failure of self-determination，can be passed through the above mentioned area name (https://xxxxxxxx.railway.app/ping)

，The result returned `pong`，The description department succeeded.。

</details>

<details>
    <summary>Docker Development</summary>
<br>

```bash
docker build -t lark-openai:latest .
docker run -d --name lark-openai -p 9000:9000 \
--env APP_LANG=en \
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

- `BOT_NAME` 为lark机器人名称，例如 `chatGpt`
- `OPENAI_KEY` 为openai key，多个key用逗号分隔，例如 `sk-xxx1,sk-xxx2,sk-xxx3`
- `HTTP_PROXY` 为宿主机的proxy地址，例如 `http://host.docker.internal:7890`,没有代理的话，可以不用设置
- `API_URL` 为openai api 接口地址，例如 `https://api.openai.com`, 没有反向代理的话，可以不用设置

--- 

部署azure版本

```bash
docker build -t lark-openai:latest .
docker run -d --name lark-openai -p 9000:9000 \
--env APP_LANG=en \
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

- `BOT_NAME` 为lark机器人名称，例如 `chatGpt`
- `AZURE_ON` 为是否使用azure ,请填写 `true`
- `AZURE_API_VERSION` 为azure api版本 例如 `2023-03-15-preview`
- `AZURE_RESOURCE_NAME` 为azure 资源名称 类似 `https://{AZURE_RESOURCE_NAME}.openai.azure.com`
- `AZURE_DEPLOYMENT_NAME` 为azure 部署名称 类似 `https://{AZURE_RESOURCE_NAME}.openai.azure.com/deployments/{AZURE_DEPLOYMENT_NAME}/chat/completions`
- `AZURE_OPENAI_TOKEN` 为azure openai token

</details>

<details>
    <summary>Docker-Compose Development</summary>
<br>

Edit docker-compose.yaml, configure the corresponding environment variable through environment (or mount the corresponding configuration file through volumes), and then run the following command
```bash
# 构建镜像
docker compose build

# 启动服务
docker compose up -d

# 停止服务
docker compose down
```

Event callback address: http://IP:9000/webhook/event
    
Card callback address: http://IP:9000/webhook/card

</details>


## Detailed configuration steps

<details align='left'>
    <summary> 📸 Click to expand the step-by-step screenshot guide for lark robot configuration</summary>
    <br>
    <p align='center'>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/a2bf0588-0fff-48a7-a253-25d237c37f0e' alt='' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/60b9dc76-3117-42c0-8086-6d5938161127' alt='' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/1f46d819-a063-42fd-bf28-6d31086e1912' alt='' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/47057139-1e09-48da-97ff-86f9021182f0' alt='' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/91dc1b09-664e-4dea-b6b8-ca7a656d1ac4' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/272bb80c-b9aa-49e2-9411-c0357ca03fe8' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/aed03155-22cd-446a-96d8-54cfd95e04fb' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/8f306e96-8767-480f-858f-41623038dbd2' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/b9a6d27f-f225-4d02-8b05-b8ab4dd55fa5' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/b2fbe22f-4920-4628-b673-011d036ae4fb' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/607db6b7-07d2-4307-9ded-9e959bd15fcf' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/fd8eded1-d248-4552-94d2-21e983d24dcd' alt='' width='800'/>
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/4d7b0a16-2871-4472-bd78-78d3989c050d' alt='' width='800'/>
    </p>
</details>


- 获取 [OpenAI](https://platform.openai.com/account/api-keys) 的 KEY( 🙉 下面有免费的 KEY 供大家测试部署 )
- 创建 [lark](https://open.larksuit.com/) 机器人
    1. 前往[开发者平台](https://open.larksuite.com/app?lang=zh-CN)创建应用,并获取到 APPID 和 Secret
    2. 前往`应用功能-机器人`, 创建机器人
    3. 从 cpolar、serverless 或 Railway 获得公网地址，在lark机器人后台的 `事件订阅` 板块填写。例如，
        - `http://xxxx.r6.cpolar.top`为 cpolar 暴露的公网地址
        - `/webhook/event`为统一的应用路由
        - 最终的回调地址为 `http://xxxx.r6.cpolar.top/webhook/event`
    4. 在lark机器人后台的 `机器人` 板块，填写消息卡片请求网址。例如，
        - `http://xxxx.r6.cpolar.top`为 cpolar 暴露的公网地址
        - `/webhook/card`为统一的应用路由
        - 最终的消息卡片请求网址为 `http://xxxx.r6.cpolar.top/webhook/card`
    5. 在事件订阅板块，搜索三个词`机器人进群`、 `接收消息`、 `消息已读`, 把他们后面所有的权限全部勾选。
       进入权限管理界面，搜索`图片`, 勾选`获取与上传图片或文件资源`。
       最终会添加下列回调事件
        - im:resource(获取与上传图片或文件资源)
        - im:message
        - im:message.group_at_msg(获取群组中所有消息)
        - im:message.group_at_msg:readonly(接收群聊中@机器人消息事件)
        - im:message.p2p_msg(获取用户发给机器人的单聊消息)
        - im:message.p2p_msg:readonly(读取用户发给机器人的单聊消息)
        - im:message:send_as_bot(获取用户在群组中@机器人的消息)
        - im:chat:readonly(获取群组信息)
        - im:chat(获取与更新群组信息)


5. Publish the version and wait for the approval of the enterprise administrator

## Connect-AI More


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




