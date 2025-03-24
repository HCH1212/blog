package ai

import (
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/components/model"
)

func CreateOllamaChatModel(ctx context.Context) (model.ChatModel, error) {
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		// 基础配置
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Timeout: 120 * time.Second,        // 请求超时时间

		// 模型配置
		Model: "deepseek-r1:8b", // 模型名称
	})
	if err != nil {
		return nil, err
	}
	return chatModel, nil
}
