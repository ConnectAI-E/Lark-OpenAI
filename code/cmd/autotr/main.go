package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/text/language"
	"golang.org/x/text/message/pipeline"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	msgs := new(pipeline.Messages)
	json.Unmarshal(content, &msgs)

	config := openai.DefaultConfig(os.Getenv("OPENAI_TOKEN"))
	if val := os.Getenv("OPENAI_BASE_URL"); len(val) > 0 {
		config.BaseURL = val
	}
	client := openai.NewClientWithConfig(config)

	prompt := "%s"
	switch msgs.Language {
	case language.English:
		prompt = "请将如下内容翻译成英文：\n\n\n%s"
	case language.Japanese:
		prompt = "请将如下内容翻译成日语：\n\n\n%s"
	case language.Vietnamese:
		prompt = "请将如下内容翻译成越南语：\n\n\n%s"
	}

	for idx, msg := range msgs.Messages {
		if len(msg.Translation.Msg) > 0 {
			continue
		}
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: fmt.Sprintf(prompt, msg.Translation.Msg),
					},
				},
			},
		)
		if err != nil {
			panic(err)
		}

		msgs.Messages[idx].Translation.Msg = resp.Choices[0].Message.Content
		content, err := json.Marshal(msgs)
		if err != nil {
			panic(err)
		}
		os.WriteFile(os.Args[1], content, 0644)
		log.Printf("%d/%d", idx, len(msgs.Messages))
	}
}
