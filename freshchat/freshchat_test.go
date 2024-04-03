package freshchat_test

import "os"

var (
	url         = os.Getenv("FRESHCHAT_URL")
	accessToken = os.Getenv("FRESHCHAT_ACCESS_TOKEN")

	userID         = os.Getenv("FRESHCHAT_USER_ID")
	agentID        = os.Getenv("FRESHCHAT_AGENT_ID")
	channelID      = os.Getenv("FRESHCHAT_CHANNEL_ID")
	conversationID = os.Getenv("FRESHCHAT_CONVERSATION_ID")
)
