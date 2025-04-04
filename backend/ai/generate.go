package ai

import (
	"context"
	"log"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func Generate(ctx context.Context, llm model.ChatModel, in []*schema.Message) (*schema.Message, error) {
	result, err := llm.Generate(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func stream(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message] {
	result, err := llm.Stream(ctx, in)
	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}
