package freshchat_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/toanppp/freshgo/freshchat"
)

func TestFreshchat_CreateConversation(t *testing.T) {
	f := freshchat.New(url, accessToken)
	resp, err := f.CreateConversation(context.Background(), freshchat.Conversation{
		Status: freshchat.ConversationStatusNew,
		Messages: []freshchat.Message{
			{
				MessageParts: []freshchat.MessagePart{
					{
						Text: freshchat.Text{
							Content: "Hello World",
						},
					},
				},
				ActorType: freshchat.ActorTypeUser,
				ActorID:   userID,
			},
		},
		ChannelID: channelID,
		Users: []freshchat.User{
			{
				ID: userID,
			},
		},
	})
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestFreshchat_ListMessages(t *testing.T) {
	f := freshchat.New(url, accessToken)
	resp, err := f.ListMessages(context.Background(), conversationID, freshchat.ListReq{})
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}
	if len(resp.Messages) == 0 {
		t.Fatalf("empty message: %+v", resp)
	}
	fmt.Printf("%+v\n", resp)
}

func TestFreshchat_SendMessage(t *testing.T) {
	f := freshchat.New(url, accessToken)
	resp, err := f.SendMessage(context.Background(), conversationID, freshchat.Message{
		MessageParts: []freshchat.MessagePart{
			{
				Text: freshchat.Text{
					Content: "Second message",
				},
			},
		},
		ActorType: freshchat.ActorTypeUser,
		ActorID:   userID,
	})
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestFreshchat_UpdateConversation(t *testing.T) {
	f := freshchat.New(url, accessToken)
	resp, err := f.UpdateConversation(context.Background(), conversationID, freshchat.Conversation{
		Status:    freshchat.ConversationStatusResolved,
		ChannelID: channelID,
	})
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}
