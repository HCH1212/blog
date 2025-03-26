package service

import (
	"context"
	"github.com/HCH1212/blog/backend/ai"
	"time"
)

func ChatService(ctx context.Context, content string) (string, error) {
	// 使用模版创建messages
	messages, err := ai.CreateMessagesFromTemplate(content)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	// 创建llm
	//llm, err := ai.CreateOllamaChatModel(ctx)
	//if err != nil {
	//	return "", err
	//}
	llm, err := ai.CreateOpenAIChatModel(ctx)
	if err != nil {
		return "", err
	}

	// 生成回复
	res, err := ai.Generate(ctx, llm, messages)
	if err != nil {
		return "", err
	}

	// 使用deepseek时去掉思考过程
	//// 找到 <think> 标签的结束位置
	//start := strings.Index(res.Content, "</think>")
	//if start == -1 {
	//	fmt.Println("未找到 </think> 标签")
	//}
	//
	//// 截取 <think> 标签之后的内容
	//contentAfterThink := res.Content[start+len("</think>"):]
	//
	//// 去除多余的空格和换行符
	//contentAfterThink = strings.TrimSpace(contentAfterThink)
	//
	//return contentAfterThink, nil
	return res.Content, nil
}
