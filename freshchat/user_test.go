package freshchat_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/toanppp/freshgo/freshchat"
)

func TestFreshchat_CreateUser(t *testing.T) {
	f := freshchat.New(url, accessToken)
	resp, err := f.CreateUser(context.Background(), freshchat.User{
		ReferenceID: "100115360",
		FirstName:   "Toan",
		LastName:    "Pham",
		Email:       "toan.pham3@tiki.vn",
	})
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}
	if resp.ID == "" {
		t.Errorf("invalid id: %+v", resp)
	}
	fmt.Printf("%+v\n", resp)
}

func TestFreshchat_ListUsers(t *testing.T) {
	f := freshchat.New(url, accessToken)
	resp, err := f.ListUsers(context.Background(), freshchat.UsersReq{
		ListReq: freshchat.ListReq{
			Page:         1,
			ItemsPerPage: 1,
		},
		ReferenceID: "100115360",
	})
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}
	if len(resp.Users) == 0 {
		t.Fatalf("empty users: %+v", resp)
	}
	if resp.Users[0].ReferenceID != "100115360" {
		t.Errorf("invalid reference_id: %+v", resp)
	}
	fmt.Printf("%+v\n", resp)
}
