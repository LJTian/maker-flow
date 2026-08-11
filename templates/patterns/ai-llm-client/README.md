# ai-llm-client

Streaming AI client pattern built on `sashabaranov/go-openai`.
Works with OpenAI and any **OpenAI-compatible** gateway (some Anthropic/Ollama proxies). There is **no** native Anthropic SDK in this pattern.

## Agent Usage Instructions

1. **Copy** `llm.go` into `<product-root>/<app-id>/internal/ai/`.
2. **Install deps**: 
   `go get github.com/sashabaranov/go-openai`
3. **Initialize in `main.go`**:
   ```go
   import "your_app/internal/ai"

   llmClient := ai.NewClient(ai.Config{
       BaseURL: os.Getenv("LLM_BASE_URL"), // Optional, for custom endpoints
       APIKey:  os.Getenv("LLM_API_KEY"),
       Model:   "gpt-4o",
   })
   ```
4. **SSE Streaming Route (Gin Example)**:
   ```go
   r.GET("/api/chat", func(c *gin.Context) {
       c.Writer.Header().Set("Content-Type", "text/event-stream")
       c.Writer.Header().Set("Cache-Control", "no-cache")
       c.Writer.Header().Set("Connection", "keep-alive")

       respChan, errChan := llmClient.GenerateStream(c.Request.Context(), "You are helpful", "Tell me a joke")
       
       for {
           select {
           case chunk, ok := <-respChan:
               if !ok {
                   return // EOF
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
