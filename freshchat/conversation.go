package freshchat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	ConversationStatusNew      = "new"
	ConversationStatusAssigned = "assigned"
	ConversationStatusResolved = "resolved"
	ConversationStatusReopened = "reopened"

	ActorTypeUser = "user"
)

type Conversation struct {
	ConversationID     string     `json:"conversation_id,omitempty"`
	ChannelID          string     `json:"channel_id,omitempty"`
	AssignedOrgAgentID string     `json:"assigned_org_agent_id,omitempty"`
	AssignedAgentID    string     `json:"assigned_agent_id,omitempty"`
	AssignedOrgGroupID string     `json:"assigned_org_group_id,omitempty"`
	AssignedGroupID    string     `json:"assigned_group_id,omitempty"`
	Messages           []Message  `json:"messages,omitempty"`
	AppID              string     `json:"app_id,omitempty"`
	Status             string     `json:"status,omitempty"`
	SkillID            int64      `json:"skill_id,omitempty"`
	Properties         Properties `json:"properties,omitempty"`
	Users              []User     `json:"users,omitempty"`
}

type Message struct {
	MessageParts     []MessagePart `json:"message_parts,omitempty"`
	AppID            string        `json:"app_id,omitempty"`
	ActorID          string        `json:"actor_id,omitempty"`
	OrgActorID       string        `json:"org_actor_id,omitempty"`
	ID               string        `json:"id,omitempty"`
	ChannelID        string        `json:"channel_id,omitempty"`
	ConversationID   string        `json:"conversation_id,omitempty"`
	InteractionID    string        `json:"interaction_id,omitempty"`
	MessageType      string        `json:"message_type,omitempty"`
	ActorType        string        `json:"actor_type,omitempty"`
	CreatedTime      string        `json:"created_time,omitempty"`
	UserID           string        `json:"user_id,omitempty"`
	RestrictResponse bool          `json:"restrictResponse,omitempty"`
	BotsPrivateNote  bool          `json:"botsPrivateNote,omitempty"`
}

type MessagePart struct {
	Text Text `json:"text,omitempty"`
}

type Text struct {
	Content string `json:"content,omitempty"`
}

type Properties struct {
	Priority string `json:"priority,omitempty"`
	CFType   string `json:"cf_type,omitempty"`
	CFRating string `json:"cf_rating,omitempty"`
}

type MessagesResp struct {
	Messages []Message `json:"messages,omitempty"`
}

func (f *freshchat) CreateConversation(ctx context.Context, input Conversation) (Conversation, error) {
	url := fmt.Sprintf("%s/%s", f.url, pathConversations)
	pReq, _ := json.Marshal(input)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(pReq))
	if err != nil {
		return Conversation{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", f.accessToken)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return Conversation{}, fmt.Errorf("httpClient.Do: %w", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		pResp, _ := io.ReadAll(resp.Body)
		return Conversation{}, fmt.Errorf("request failed: %s", string(pResp))
	}

	pResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conversation{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output Conversation
	if err := json.Unmarshal(pResp, &output); err != nil {
		return Conversation{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}

func (f *freshchat) UpdateConversation(ctx context.Context, conversationID string, input Conversation) (Conversation, error) {
	url := fmt.Sprintf("%s/%s", f.url, fmt.Sprintf(pathConversation, conversationID))
	pReq, _ := json.Marshal(input)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(pReq))
	if err != nil {
		return Conversation{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", f.accessToken)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return Conversation{}, fmt.Errorf("httpClient.Do: %w", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		pResp, _ := io.ReadAll(resp.Body)
		return Conversation{}, fmt.Errorf("request failed: %s", string(pResp))
	}

	pResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conversation{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output Conversation
	if err := json.Unmarshal(pResp, &output); err != nil {
		return Conversation{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}

func (f *freshchat) ListMessages(ctx context.Context, conversationID string, input ListReq) (MessagesResp, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/%s", f.url, fmt.Sprintf(pathMessages, conversationID)))
	u.RawQuery = input.Values().Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return MessagesResp{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}
	req.Header.Add("Authorization", f.accessToken)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return MessagesResp{}, fmt.Errorf("httpClient.Do: %w", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		p, _ := io.ReadAll(resp.Body)
		return MessagesResp{}, fmt.Errorf("request failed: %s", string(p))
	}

	p, err := io.ReadAll(resp.Body)
	if err != nil {
		return MessagesResp{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output MessagesResp
	if err := json.Unmarshal(p, &output); err != nil {
		return MessagesResp{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}

func (f *freshchat) SendMessage(ctx context.Context, conversationID string, input Message) (Message, error) {
	url := fmt.Sprintf("%s/%s", f.url, fmt.Sprintf(pathMessages, conversationID))
	pReq, _ := json.Marshal(input)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(pReq))
	if err != nil {
		return Message{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", f.accessToken)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return Message{}, fmt.Errorf("httpClient.Do: %w", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		pResp, _ := io.ReadAll(resp.Body)
		return Message{}, fmt.Errorf("request failed: %s", string(pResp))
	}

	pResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output Message
	if err := json.Unmarshal(pResp, &output); err != nil {
		return Message{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}
