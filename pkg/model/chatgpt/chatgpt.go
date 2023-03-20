package chatgpt

import (
	"context"

	"portal/pkg/model"

	"github.com/sashabaranov/go-openai"
	"github.com/zeromicro/go-zero/core/logx"
)

type GptClient struct {
	Config  model.Config
	Client  *openai.Client
	Session map[string][]openai.ChatCompletionMessage
}

func New(conf model.Config) *GptClient {
	return &GptClient{
		Config:  conf,
		Client:  openai.NewClient(conf.ApiKey),
		Session: make(map[string][]openai.ChatCompletionMessage),
	}
}

func (c *GptClient) Reply(userId, query string) (reply string, err error) {
	newQuery := c.buildSessionQuery(userId, query)

	return c.ReplyText(userId, newQuery)
}

func (c *GptClient) ReplyText(userId string, messages []openai.ChatCompletionMessage) (reply string, err error) {
	resp, err := c.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            openai.GPT3Dot5Turbo,
			Messages:         messages,
			MaxTokens:        c.Config.MaxTokens,
			Temperature:      0.9,
			TopP:             1,
			FrequencyPenalty: 0,
			PresencePenalty:  0,
		},
	)
	if err != nil {
		return
	}

	reply = resp.Choices[0].Message.Content
	usedToken := resp.Usage.TotalTokens
	logx.Debugf("[CHATGPT] response: %v", resp)
	logx.Infof("[CHATGPT] reply %s", reply)
	if len(reply) != 0 {
		c.saveSession(userId, reply, usedToken)
	}
	return
}

func (c *GptClient) buildSessionQuery(userId, query string) []openai.ChatCompletionMessage {
	var sessions []openai.ChatCompletionMessage
	sessions, ok := c.Session[userId]
	if !ok {
		session := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: c.Config.Description,
		}
		sessions = append(sessions, session)
	}

	session := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: query,
	}
	sessions = append(sessions, session)
	c.Session[userId] = sessions
	return sessions
}

func (c *GptClient) saveSession(userId, answer string, usedTokens int) {
	maxTokens := c.Config.MaxTokens
	if maxTokens <= 0 || maxTokens > 4000 {
		maxTokens = 1000
	}

	sessions, ok := c.Session[userId]
	if ok {
		session := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: answer,
		}
		sessions = append(sessions, session)
	}

	if usedTokens > maxTokens && len(sessions) >= 3 {
		sessions = sessions[len(sessions)-1:]
	}
}

func (c *GptClient) ClearSession(userId string) {
	c.Session[userId] = c.Session[userId][:0]
}
