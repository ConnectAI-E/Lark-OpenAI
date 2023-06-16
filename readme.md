<p align='center'>
    <a href='https://www.qiniai.com' target="_blank" rel="noopener noreferrer">
    <img src='https://github-production-user-asset-6210df.s3.amazonaws.com/50035229/243593939-421327a5-1cf1-46ad-93b1-15c9c9d36490.png' alt='' width='800'/>
    </a>
</p>


<details align='center'>
    <summary> 📷 Click to expand the full function of Connect-AI</summary>
    <br>
    <p align='center'>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/b993c610-1c91-40dd-bdcd-85a992c17b74' alt='' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/149f5fda-3fc4-49fa-8132-4825edfece1f' alt='' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/7dae5661-2d4c-4584-934c-747a8c68d3e9' alt='' width='800'/>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/942ffb30-fb48-4de4-a696-e0903a691665' alt='' width='800'/>
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
   www.connectai-e.com
</p>




<strong align="center">
<samp>

[**English**](./readme.md) · [**简体中文**](https://github.com/ConnectAI-E/Feishu-OpenAI)· [**繁體中文**](./readme_zh-hk.md) · [**日本語**](./readme_jp.md) · [**Tiếng Việt**](./readme_vi.md)

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
- 🥒 support [Serverless](https://github.com/serverless-devs/serverless-devs)、[local](https://dashboard.cpolar.com/login)、[Docker](https://www.docker.com/)、[binary package](https://github.com/Leizhenpeng/feishu-chatgpt/releases/)


## 🛵 Development

<details>
    <summary>Local Development</summary>
<br>

```bash
git clone git@github.com:ConnectAI-E/lark-openai.git
cd Lark-OpenAI/code
```

If your server does not have a public network IP, you can use a reverse proxy.

The server of Flying Book is very slow to access ngrok in China, so it is recommended to use some domestic reverse proxy service providers.

- [cpolar](https://dashboard.cpolar.com/)
- [natapp](https://natapp.cn/)

```bash
# Configure config.yaml
mv config.example.yaml config.yaml

// Testing deployment.
go run ./
cpolar http 9000

//Production deployment
nohup cpolar http 9000 -log=stdout &

//Check server status
https://dashboard.cpolar.com/status

// Take down the service
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
cd Lark-OpenAI/code
```

install [severless](https://docs.serverless-devs.com/serverless-devs/quick_start)tool

```bash
# Configure config.yaml
mv config.example.yaml config.yaml
# install severless cli
npm install @serverless-devs/s -g
```

After the installation is complete, please deploy according to your local environment and the following tutorial`severless`

- local `linux`/`mac os` env

1. Modify the Deployment Region and Deployment Key in 's.yaml'

```
edition: 1.0.0
name: lark-openai
access: "aliyun" #  Modify the custom key name.

vars: # Global variables
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
access: "aliyun" #  Modify the custom key alias

vars: # Global variables
  region: "ap-southeast-1" # Modify the desired deployment region for the cloud functions

```

3. Modify `pre-deploy` in `s.yaml`, remove the ring variable change part before the second step `run`

```
  pre-deploy:
        - run: go mod tidy
          path: ./code
        - run: go build -o
            target/main main.go  # del GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0
          path: ./code

```

4. One-click deployment

```bash
cd ..
s deploy
```

<br>
</details>

<!-- <details>
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

</details> -->

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

Attention:

- `APP_LANG` is the language of the Lark bot, for example, `en`,`ja`, `vi`,`zh-hk`.
- `BOT_NAME` is the name of the Lark bot, for example, `chatGpt`.
- `OPENAI_KEY` is the OpenAI key. If you have multiple keys, separate them with commas, for example, `sk-xxx1,sk-xxx2,sk-xxx3`.
- `HTTP_PROXY` is the proxy address of the host machine, for example, `http://host.docker.internal:7890`. If you don't have a proxy, you can leave this unset.
- `API_URL` is the OpenAI API endpoint address, for example, `https://api.openai.com`. If you don't have a reverse proxy, you can leave this unset.

--- 

To deploy the Azure version

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
    
Attention:

- `APP_LANG` is the language of the Lark bot, for example, `en`,`ja`, `vi`,`zh-hk`.
- `BOT_NAME` is the name of the Lark bot, for example, `chatGpt`.
- `AZURE_ON` indicates whether to use Azure. Please set it to `true`.
- `AZURE_API_VERSION` is the Azure API version, for example, `2023-03-15-preview`.
- `AZURE_RESOURCE_NAME` is the Azure resource name, similar to `https://{AZURE_RESOURCE_NAME}.openai.azure.com`.
- `AZURE_DEPLOYMENT_NAME` is the Azure deployment name, similar to `https://{AZURE_RESOURCE_NAME}.openai.azure.com/deployments/{AZURE_DEPLOYMENT_NAME}/chat/completions`.
- `AZURE_OPENAI_TOKEN` is the Azure OpenAI token.

</details>

<details>
    <summary>Docker-Compose Development</summary>
<br>

Edit docker-compose.yaml, configure the corresponding environment variable through environment (or mount the corresponding configuration file through volumes), and then run the following command
```bash
# Build the image
docker compose build

# Start the service
docker compose up -d

# Stop the service
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


- Get [OpenAI](https://platform.openai.com/account/api-keys) KEY( 🙉 Below are free keys available for everyone to test deployment )
- Create [lark](https://open.larksuit.com/) Bot
    1. Go [Lark Open Platform](https://open.larksuite.com/?lang=en-US) creat app , get APPID and Secret
    2. Go `Features-Bot` , creat bot
    3. Obtain the public address from cpolar, serverless, or Railway, and fill it in the "Event Subscription" section of the Lark bot backend. For example,
   - `http://xxxx.r6.cpolar.top` is the public address exposed by cpolar.
   - `/webhook/event` is the unified application route.
   - The final callback address is `http://xxxx.r6.cpolar.top/webhook/event`.
    4. In the "Bot" section of the Lark bot backend, fill in the request URL for message cards. For example,
   - `http://xxxx.r6.cpolar.top` is the public address exposed by cpolar.
   - `/webhook/card` is the unified application route.
   - The final request URL for message cards is `http://xxxx.r6.cpolar.top/webhook/card`.

    5. In the "Event Subscription" section, search for the three terms: "Bot Join Group," "Receive Messages," and "Messages Read." Check all the permissions behind them.
   Go to the permission management interface, search for "Image," and check "Get and upload image or file resources."
   Finally, the following callback events will be added.
        - im:resource(Read and upload images or other files)
        - im:message
        - im:message.group_at_msg(Read group chat messages mentioning the bot)
        - im:message.group_at_msg:readonly(Obtain group messages mentioning the bot)
        - im:message.p2p_msg(Read private messages sent to the bot)
        - im:message.p2p_msg:readonly(Obtain private messages sent to the bot)
        - im:message:send_as_bot(Send messages as an app)
        - im:chat:readonly(Obtain group information)
        - im:chat(Obtain and update group information)


5. Publish the version and wait for the approval of the enterprise administrator

## Connect-AI More

| <div style="width:200px">AI</div> |             <img width=110> SDK <img width=110>              |                         Application                          |
| :-------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
|              🎒OpenAI              |    [Go-OpenAI](https://github.com/ConnectAI-E/Go-OpenAI)     | [🏅Feishu-OpenAI](https://github.com/ConnectAI-E/Feishu-OpenAI), [🎖Lark-OpenAI](https://github.com/ConnectAI-E/Lark-OpenAI), [Feishu-EX-ChatGPT](https://github.com/ConnectAI-E/Feishu-EX-ChatGPT), [🎖Feishu-OpenAI-Stream-Chatbot](https://github.com/ConnectAI-E/Feishu-OpenAI-Stream-Chatbot), [Feishu-TLDR](https://github.com/ConnectAI-E/Feishu-TLDR),[Feishu-OpenAI-Amazing](https://github.com/ConnectAI-E/Feishu-OpenAI-Amazing), [Feishu-Oral-Friend](https://github.com/ConnectAI-E/Feishu-Oral-Friend), [Feishu-OpenAI-Base-Helper](https://github.com/ConnectAI-E/Feishu-OpenAI-Base-Helper), [Feishu-Vector-Knowledge-Management](https://github.com/ConnectAI-E/Feishu-Vector-Knowledge-Management), [Feishu-OpenAI-PDF-Helper](https://github.com/ConnectAI-E/Feishu-OpenAI-PDF-Helper), [🏅Dingtalk-OpenAI](https://github.com/ConnectAI-E/Dingtalk-OpenAI), [Wework-OpenAI](https://github.com/ConnectAI-E/Wework-OpenAI), [WeWork-OpenAI-Node](https://github.com/ConnectAI-E/WeWork-OpenAI-Node), [llmplugin](https://github.com/ConnectAI-E/llmplugin) |
|             🤖 AutoGPT             |                            ------                            | [🏅AutoGPT-Next-Web](https://github.com/ConnectAI-E/AutoGPT-Next-Web) |
|         🎭 Stablediffusion         |                            ------                            | [🎖Feishu-Stablediffusion](https://github.com/ConnectAI-E/Feishu-Stablediffusion) |
|           🍎 Midjourney            | [Go-Midjourney](https://github.com/ConnectAI-E/Feishu-Midjourney/tree/main/midjourney) | [🏅Feishu-Midjourney](https://github.com/ConnectAI-E/Feishu-Midjourney), [MidJourney-Web](https://github.com/ConnectAI-E/MidJourney-Web), [Dingtalk-Midjourney](https://github.com/ConnectAI-E/Dingtalk-Midjourney) |
|            🍍 文心一言             |    [Go-Wenxin](https://github.com/ConnectAI-E/Go-Wenxin)     | [Feishu-Wenxin](https://github.com/ConnectAI-E/Feishu-Wenxin), [Dingtalk-Wenxin](https://github.com/ConnectAI-E/Dingtalk-Wenxin), [Wework-Wenxin](https://github.com/ConnectAI-E/Wework-Wenxin) |
|             💸 Minimax             |   [Go-Minimax](https://github.com/ConnectAI-E/Go-Minimax)    | [Feishu-Minimax](https://github.com/ConnectAI-E/Feishu-Minimax), [Dingtalk-Minimax](https://github.com/ConnectAI-E/Dingtalk-Minimax), [Wework-Minimax](https://github.com/ConnectAI-E/Wework-Minimax) |
|             ⛳️ CLAUDE              |    [Go-Claude](https://github.com/ConnectAI-E/Go-Claude)     | [Feishu-Claude](https://github.com/ConnectAI-E/Feishu-Claude), [DingTalk-Claude](https://github.com/ConnectAI-E/DingTalk-Claude), [Wework-Claude](https://github.com/ConnectAI-E/Wework-Claude) |
|              🥁 PaLM               |      [Go-PaLM](https://github.com/ConnectAI-E/go-PaLM)       | [Feishu-PaLM](https://github.com/ConnectAI-E/Feishu-PaLM),[DingTalk-PaLM](https://github.com/ConnectAI-E/DingTalk-PaLM),[Wework-PaLM](https://github.com/ConnectAI-E/Wework-PaLM) |
|             🎡 Prompt              |                            ------                            | [📖 Prompt-Engineering-Tutior](https://github.com/ConnectAI-E/Prompt-Engineering-Tutior) |
|             🍋 ChatGLM             |                            ------                            | [Feishu-ChatGLM](https://github.com/ConnectAI-E/Feishu-ChatGLM) |
|            ⛓ LangChain            |                            ------                            | [📖 LangChain-Tutior](https://github.com/ConnectAI-E/LangChain-Tutior) |
|            🪄 One-click            |                            ------                            | [🎖Awesome-One-Click-Deployment](https://github.com/ConnectAI-E/Awesome-One-Click-Deployment) |



