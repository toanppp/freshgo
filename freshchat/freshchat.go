package freshchat

import (
	"context"
	"fmt"
	"net/http"
)

const (
	pathUsers         = "users"
	pathConversations = "conversations"
	pathConversation  = "conversations/%s"
	pathMessages      = "conversations/%s/messages"
	pathAgents        = "agents"
	pathAgent         = "agents/%s"
)

type Freshchat interface {
	CreateUser(ctx context.Context, input User) (User, error)
	ListUsers(ctx context.Context, input UsersReq) (UsersResp, error)

	CreateConversation(ctx context.Context, input Conversation) (Conversation, error)
	UpdateConversation(ctx context.Context, conversationID string, input Conversation) (Conversation, error)
	ListMessages(ctx context.Context, conversationID string, input ListReq) (MessagesResp, error)
	SendMessage(ctx context.Context, conversationID string, input Message) (Message, error)

	GetAgentInfo(ctx context.Context, agentID string) (Agent, error)
}

type freshchat struct {
	url         string
	accessToken string
	httpClient  *http.Client
}

func New(url, accessToken string, opts ...Option) Freshchat {
	f := &freshchat{
		url:         url,
		accessToken: fmt.Sprintf("Bearer %s", accessToken),
		httpClient:  http.DefaultClient,
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

type Option func(f *freshchat)

func OptionHTTPClient(client *http.Client) func(*freshchat) {
	return func(z *freshchat) {
		z.httpClient = client
	}
}
