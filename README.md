# freshgo

Freshworks Golang APIs

## To do

- [x] [Freshchat](https://developers.freshchat.com/)
    - [x] User
    - [x] Agent
    - [x] Conversation
- [ ] Freshdesk

## Install

```sh
go get github.com/toanppp/freshgo
```

## Example

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/toanppp/freshgo/freshchat"
)

func main() {
	f := freshchat.New("url", "accessToken")

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
				ActorID:   "userID",
			},
		},
		ChannelID: "channelID",
		Users: []freshchat.User{
			{
				ID: "userID",
			},
		},
	})
	if err != nil {
		log.Fatalf("an error occurred: %v", err)
	}

	fmt.Printf("%+v\n", resp)
}
```
