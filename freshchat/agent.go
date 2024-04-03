package freshchat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Agent struct {
	ID                 string          `json:"id"`
	CreatedTime        time.Time       `json:"created_time"`
	AgentCapacity      int             `json:"agent_capacity"`
	AgentStatus        AgentStatus     `json:"agent_status"`
	AvailabilityStatus string          `json:"availability_status"`
	Avatar             Avatar          `json:"avatar"`
	Biography          string          `json:"biography"`
	Email              string          `json:"email"`
	FirstName          string          `json:"first_name"`
	FreshidGroupDs     any             `json:"freshid_group_ids"`
	FreshidUuid        string          `json:"freshid_uuid"`
	Groups             []string        `json:"groups"`
	IsDeactivated      bool            `json:"is_deactivated"`
	IsDeleted          bool            `json:"is_deleted"`
	LastName           string          `json:"last_name"`
	LicenseType        string          `json:"license_type"`
	Locale             string          `json:"locale"`
	LoginStatus        bool            `json:"login_status"`
	OrgContactID       string          `json:"org_contact_id"`
	RoleID             string          `json:"role_id"`
	RoleName           string          `json:"role_name"`
	RoutingType        string          `json:"routing_type"`
	SkillID            string          `json:"skill_id"`
	SocialProfiles     []SocialProfile `json:"social_profiles"`
	Timezone           string          `json:"timezone"`
}

type AgentStatus struct {
	Name string `json:"name"`
}

type SocialProfile struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (f *freshchat) GetAgentInfo(ctx context.Context, agentID string) (Agent, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/%s", f.url, fmt.Sprintf(pathAgent, agentID)))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return Agent{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}
	req.Header.Add("Authorization", f.accessToken)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return Agent{}, fmt.Errorf("httpClient.Do: %w", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		p, _ := io.ReadAll(resp.Body)
		return Agent{}, fmt.Errorf("request failed: %s", string(p))
	}

	p, err := io.ReadAll(resp.Body)
	if err != nil {
		return Agent{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output Agent
	if err := json.Unmarshal(p, &output); err != nil {
		return Agent{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}
