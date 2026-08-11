# ai-llm-client

流式 AI 大模型客户端封装（基于 `sashabaranov/go-openai`）。
适用于 OpenAI 及任意 **OpenAI 兼容**网关（部分 Anthropic/Ollama 代理）。本 pattern **没有** Anthropic 原生 SDK。

## Agent 使用说明

1. **复制** `llm.go` 到 `<产品根>/<app-id>/internal/ai/`。
2. **安装依赖**: 
   `go get github.com/sashabaranov/go-openai`
3. **在 `main.go` 中初始化**:
   ```go
   import "your_app/internal/ai"

   llmClient := ai.NewClient(ai.Config{
       BaseURL: os.Getenv("LLM_BASE_URL"), // 例如中转服务地址，不填则默认 openai
       APIKey:  os.Getenv("LLM_API_KEY"),
       Model:   "gpt-4o",
   })
   ```
4. **配合 Gin 提供打字机流式接口 (SSE)**:
   ```go
   r.GET("/api/chat", func(c *gin.Context) {
       c.Writer.Header().Set("Content-Type", "text/event-stream")
       c.Writer.Header().Set("Cache-Control", "no-cache")
       c.Writer.Header().Set("Connection", "keep-alive")

       respChan, errChan := llmClient.GenerateStream(c.Request.Context(), "你是一个聪明的助手", "讲个笑话")
       
       for {
           select {
           case chunk, ok := <-respChan:
               if !ok {
                   return // 传输完毕
               }
               c.SSEvent("message", chunk)
               c.Writer.Flush()
           case err := <-errChan:
               c.SSEvent("error", err.Error())
               return
           }
       }
   })
   ```
