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

type User struct {
	ID           string     `json:"id,omitempty"`
	CreatedTime  string     `json:"created_time,omitempty"`
	UpdatedTime  string     `json:"updated_time,omitempty"`
	Avatar       Avatar     `json:"avatar,omitempty"`
	Email        string     `json:"email,omitempty"`
	FirstName    string     `json:"first_name,omitempty"`
	LastName     string     `json:"last_name,omitempty"`
	LoginStatus  bool       `json:"login_status,omitempty"`
	OrgContactID string     `json:"org_contact_id,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	Properties   []Property `json:"properties,omitempty"`
	ReferenceID  string     `json:"reference_id,omitempty"`
	RestoreID    string     `json:"restore_id,omitempty"`
}

type Avatar struct {
	URL string `json:"url"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UsersResp struct {
	ListResp
	Users []User `json:"users"`
}

type UsersReq struct {
	ListReq
	ReferenceID string
}

func (r *UsersReq) Values() url.Values {
	values := r.ListReq.Values()
	values.Add("reference_id", r.ReferenceID)
	return values
}

func (f *freshchat) CreateUser(ctx context.Context, input User) (User, error) {
	url := fmt.Sprintf("%s/%s", f.url, pathUsers)
	pReq, _ := json.Marshal(input)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(pReq))
	if err != nil {
		return User{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", f.accessToken)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("httpClient.Do: %w", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		pResp, _ := io.ReadAll(resp.Body)
		return User{}, fmt.Errorf("request failed: %s", string(pResp))
	}

	pResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output User
	if err := json.Unmarshal(pResp, &output); err != nil {
		return User{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}

func (f *freshchat) ListUsers(ctx context.Context, input UsersReq) (UsersResp, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/%s", f.url, pathUsers))
	u.RawQuery = input.Values().Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return UsersResp{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}
	req.Header.Add("Authorization", f.accessToken)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return UsersResp{}, fmt.Errorf("httpClient.Do: %w", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		p, _ := io.ReadAll(resp.Body)
		return UsersResp{}, fmt.Errorf("request failed: %s", string(p))
	}

	p, err := io.ReadAll(resp.Body)
	if err != nil {
		return UsersResp{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output UsersResp
	if err := json.Unmarshal(p, &output); err != nil {
		return UsersResp{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}
