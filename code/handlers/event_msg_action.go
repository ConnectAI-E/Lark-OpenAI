package handlers

import (
	"lark-openai/services/openai"
	"lark-openai/utils"
	"time"
)

type MessageAction struct { /*消息*/
}

func (*MessageAction) Execute(a *ActionInfo) bool {
	msg := a.handler.sessionCache.GetMsg(*a.info.sessionId)
	if !hasSystemRole(msg) {
		msg = append(msg, openai.Messages{
			Role: "system", Content: "You are ChatGPT, " +
				"a large language model trained by OpenAI. " +
				"Answer in user's language as concisely as" +
				" possible. Knowledge cutoff: 20230601 " +
				"Current date" + time.Now().Format("20060102"),
		})
	}
	msg = append(msg, openai.Messages{
		Role: "user", Content: a.info.qParsed,
	})
	// get ai mode as temperature
	aiMode := a.handler.sessionCache.GetAIMode(*a.info.sessionId)
	completions, err := a.handler.gpt.Completions(msg, aiMode)
	if err != nil {
		replyMsg(*a.ctx, utils.I18n.Sprintf(
			"🤖️：消息机器人摆烂了，请稍后再试～\n错误信息: %v", err), a.info.msgId)
		return false
	}
	msg = append(msg, completions)
	a.handler.sessionCache.SetMsg(*a.info.sessionId, msg)
	//if new topic
	if len(msg) == 2 {
		//fmt.Println("new topic", msg[1].Content)
		sendNewTopicCard(*a.ctx, a.info.sessionId, a.info.msgId,
			completions.Content)
		return false
	}
	err = replyMsg(*a.ctx, completions.Content, a.info.msgId)
	if err != nil {
		replyMsg(*a.ctx, utils.I18n.Sprintf(
			"🤖️：消息机器人摆烂了，请稍后再试～\n错误信息: %v", err), a.info.msgId)
		return false
	}
	return true
}

func hasSystemRole(msg []openai.Messages) bool {
	for _, m := range msg {
		if m.Role == "system" {
			return true
		}
	}
	return false
}
