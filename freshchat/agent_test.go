package freshchat_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/toanppp/freshgo/freshchat"
)

func TestFreshchat_GetAgentInfo(t *testing.T) {
	f := freshchat.New(url, accessToken)
	resp, err := f.GetAgentInfo(context.Background(), agentID)
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}
	if resp.ID != agentID {
		t.Errorf("invalid agent_id: %+v", resp)
	}
	fmt.Printf("%+v\n", resp)
}
