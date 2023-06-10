<p align='center'>
    <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/50035229/421327a5-1cf1-46ad-93b1-15c9c9d36490' alt='' width='800'/>
</p>


<details align='center'>
    <summary> 📷 「Connect-AI」のすべての機能を展開するには、クリックしてください。</summary>
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

[**English**](./readme.md) · [**简体中文**](https://github.com/ConnectAI-E/Feishu-OpenAI)· [**繁體中文**](./readme_zh-hk.md) · [**日本語**](./readme_jp.md) · [**Tiếng Việt**](./readme_vi.md)

</samp>
</strong>


## 👻ロボットの機能

🗣 音声コミュニケーション：個人的な会話ができる

💬 複数のトピックでの対話：個人やグループで複数のトピックについて話し合える

🖼 テキストから画像へ：テキストを画像に変換したり、画像を検索したりできる

🛖 シナリオ設定：豊富なシナリオリストが内蔵されており、ワンクリックでAIの役割を切り替えられる

🎭 ロールプレイ：シナリオモードをサポートし、ディスカッションの楽しさと創造性を高める

🤖 AIモード：4つのAIモードが内蔵されており、AIの知識と創造性を体感できる

🔄 コンテキスト保持：対話ボックスに返信するだけで、同じトピックのディスカッションを続けられる

⏰ 自動終了：タイムアウトにより自動的に対話が終了し、ディスカッション履歴をクリアできる

📝 リッチテキストカード：リッチなテキストカードの返信をサポートし、情報をより豊かに表示できる

👍 インタラクティブなフィードバック：即時にロボットの処理結果を取得できる

🎰 残高照会：トークンの消費状況をリアルタイムに確認できる

🔙 履歴の巻き戻し：簡単に過去の対話に戻り、トピックのディスカッションを続けられる 🚧

🔒 管理者モード：内蔵された管理者モードを使用し、より安全かつ信頼性の高い利用が可能 🚧

🌐 複数トークンの負荷分散：高頻度のプロダクションコールのシナリオを最適化

↩️ 逆プロキシのサポート：異なる地域のユーザーに対してより高速かつ安定したアクセス体験を提供

📚 Feishuドキュメントとの相互作用：企業の従業員のスーパーアシスタントになる 🚧

🎥 トピックのコンテンツを瞬時にPPTに変換：プ

レゼンテーションをより簡単にする 🚧

📊 テーブル分析：Feishuの表を簡単にインポートし、データ分析の効率を向上させる 🚧

🍊 プライベートデータのトレーニング：会社の製品情報を使用してGPTを二次トレーニングし、顧客の個別の要求により適したものにする 🚧



## 🌟 ベース

- 🍏 ダイアログはOpenAI-[GPT4](https://platform.openai.com/account/api-keys)と[Lark](https://www.larksuite.com/)に基づいています。
- 🥒 [Serverless](https://github.com/serverless-devs/serverless-devs)、[local](https://dashboard.cpolar.com/login)、[Docker](https://www.docker.com/)、[バイナリパッケージ](https://github.com/Leizhenpeng/feishu-chatgpt/releases/)をサポートしています。

## 🛵 開発

<details>
    <summary>ローカル開発</summary>
<br>

```bash
git clone git@github.com:ConnectAI-E/lark-openai.git
cd Lark-OpenAI/code
```

サーバーがパブリックネットワークIPを持っていない場合は、リバースプロキシを使用できます。

Flying Bookのサーバーは中国からngrokにアクセスするのが非常に遅いため、国内のリバースプロキシサービスプロバイダを使用することをおすすめします。

- [cpolar](https://dashboard.cpolar.com/)
- [natapp](https://natapp.cn/)

```bash
# config.yamlを設定する
mv config.example.yaml config.yaml

// デプロイのテスト
go run ./
cpolar http 9000

// 本番デプロイ
nohup cpolar http 9000 -log=stdout &

// サーバーステータスを確認する
https://dashboard.cpolar.com/status

// サービスを停止する
ps -ef | grep cpolar
kill -9 PID
```

<br>

</details>

<details>
    <summary>Serverless開発</summary>
<br>

```bash
git clone git@github.com:ConnectAI/lark-openai.git
cd Lark-OpenAI/code
```

[severless](https://docs.serverless-devs.com/serverless-devs/quick_start)ツールをインストールする

```bash
# config.yamlを設定する
mv config.example.yaml config.yaml
# severless cliをインストールする
npm install @serverless-devs/s -g
```

インストールが完了したら、ローカル環境と以下のチュートリアル`severless`に従ってデプロイしてください。

- ローカル `linux`/`mac os` 環境

1. 's.yaml'内のデプロイリージョンとデプロイキーを修正します。

```
edition: 1.0.0
name: lark-openai
access: "aliyun" # カスタムキー名を修正します。

vars: # グローバル変数
region: "cn-hongkong" # クラウドファンクションをデプロイしたいリージョンを修正します。

```

2. ワンクリックデプロイ

```bash
cd ..
s deploy
```

- ロ

ーカル `windows`

1. まず、ローカルの `cmd` コマンドプロンプトツールを開き、コンピューター上のGoの環境変数の設定を確認するために `go env` を実行します。以下の変数と値が正しいことを確認してください。

```cmd
set GO111MODULE=on
set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
```

値が正しくない場合（例：コンピューター上で `set GOOS=windows` のように設定されている場合）、以下のコマンドを実行して `GOOS` の変数値を設定してください。

```cmd
go env -w GOOS=linux
```

2. `s.yaml` 内のデプロイリージョンとデプロイキーを修正します。

```
edition: 1.0.0
name: lark-openai
access: "aliyun" # カスタムキーのエイリアスを修正します。

vars: # グローバル変数
  region: "cn-hongkong" # クラウドファンクションのデプロイリージョンを修正します。

```

3. `s.yaml` 内の `pre-deploy` を修正し、2番目のステップ `run` の前のリング変数の変更部分を削除します。

```
  pre-deploy:
        - run: go mod tidy
          path: ./code
        - run: go build -o
            target/main main.go  # GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 の部分を削除する
          path: ./code

```

4. ワンクリックデプロイ

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
    <summary>Docker 開発</summary>
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

注意：

- `APP_LANG` はLarkボットの言語です。例えば,`en`,`ja`,`vi`,`zh-hk`.
- `BOT_NAME` はLarkボットの名前です。例えば、`chatGpt`です。
- `OPENAI_KEY` はOpenAIのキーです。複数のキーを持っている場合は、カンマで区切って指定します。例えば、`sk-xxx1,sk-xxx2,sk-xxx3`です。
- `HTTP_PROXY` はホストマシンのプロキシアドレスです。例えば、`http://host.docker.internal:7890`です。プロキシを使用していない場合は、これを未設定のままにすることができます。
- `API_URL` はOpenAIのAPIエンドポイントアドレスです。例えば、`https://api.openai.com`です。リバースプロキシを使用していない場合は、これを未設定のままにすることができます。

---

Azureバージョンをデプロイするためには

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
lark-openai:latest
```

注意：
- `APP_LANG` はLarkボットの言語です。例えば,`en`,`ja`,`vi`,`zh-hk`.
- `BOT_NAME` はLarkボットの名前です。例えば、`chatGpt`です。
- `AZURE_ON` はAzureを使用するかどうかを示します。`true`に設定してください。
- `AZURE_API_VERSION` はAzureのAPIバージョンです。例えば、`2023-03-15-preview`です。
- `AZURE_RESOURCE_NAME` はAzureのリソース名であり、`https://{AZURE_RESOURCE_NAME}.openai.azure.com`と似た形式です。
- `AZURE_DEPLOYMENT_NAME` はAzureのデプロイメント名であり、`https://{AZURE_RESOURCE_NAME}.openai.azure.com/deployments/{AZURE_DEPLOYMENT_NAME}/chat/completions`と似た形式です。
- `AZURE_OPENAI_TOKEN` はAzureのOpenAIトークンです。

</details>

<details>
    <summary>Docker-Compose 開発</summary>
<br>

docker-compose.yamlを編集し、環境変数をenvironmentで設定するか、volumesを使用して対応する設定ファイルをマウントし、以下のコマンドを実行します。
```bash
# イメージをビルドする
docker compose build

# サービスを起動する
docker compose up -d

# サービスを停止する
docker compose down
```

イベントコールバックアドレス：http://IP:9000/webhook/event
    
カードコールバック

アドレス：http://IP:9000/webhook/card

</details>


## 詳細な設定手順

<details align='left'>
    <summary> 📸 Larkロボットの設定のステップバイステップのスクリーンショットガイドを展開するにはクリックしてください</summary>
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
      <img src='https://github.com/ConnectAI-E/Lark-OpenAI/assets/110169811/4d7b0a16-2871-4472-bd78-78d3989c050d' alt

='' width='800'/>
    </p>
</details>


- [OpenAI](https://platform.openai.com/account/api-keys)からAPIキーを取得してください（ 🙉 以下は誰でも利用できる無料のキーです）
- [Lark](https://open.larksuit.com/)ボットを作成してください
    1. [Feishu Open Platform](https://open.feishu.cn/?lang=en-US)にアクセスし、アプリを作成し、APPIDとSecretを取得してください
    2. 「Features-Bot」に移動し、ボットを作成してください
    3. cpolar、serverless、またはRailwayから公開アドレスを取得し、Larkボットのバックエンドの「イベント購読」セクションに入力してください。例：
   - `http://xxxx.r6.cpolar.top` はcpolarによって公開された公開アドレスです。
   - `/webhook/event` は統一されたアプリケーションのルートです。
   - 最終的なコールバックアドレスは `http://xxxx.r6.cpolar.top/webhook/event` です。
   
    4. Larkボットのバックエンドの「ボット」セクションに、メッセージカードのリクエストURLを入力してください。例：
   - `http://xxxx.r6.cpolar.top` はcpolarによって公開された公開アドレスです。
   - `/webhook/card` は統一されたアプリケーションのルートです。
   - メッセージカードの最終的なリクエストURLは `http://xxxx.r6.cpolar.top/webhook/card` です。

    5. 「イベント購読」セクションで、次の3つの用語を検索してください：「Bot Join Group」、「Receive Messages」、「Messages Read」。それらの背後にあるすべての権限をチェックしてください。
   権限管理インターフェースに移動し、「Image」を検索し、「Get and upload image or file resources」をチェックしてください。
   最後に、次のコールバックイベントが追加されます。
        - im:resource（画像や他のファイルを読み込みおよびアップロード）
        - im:message
        - im:message.group_at_msg（ボットをメンションしたグループチャットのメッセージを読み取る）
        - im:message.group_at_msg:readonly（ボットをメンションしたグループメッセージを取得する）
        - im:message.p2p_msg（ボットに送信された個人メッセージを読み取る）
        - im:message.p2p_msg:readonly（ボットに送信された個人メッセージを取得する）
        - im:message:send_as_bot（アプリとしてメッセージ

を送信する）
        - im:chat:readonly（グループ情報を取得する）
        - im:chat（グループ情報を取得および更新する）

5. バージョンを公開し、企業管理者の承認を待ちます

## Connect-AI 詳細


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




