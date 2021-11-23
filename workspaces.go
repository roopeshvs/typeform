package gotypeform

import (
	"bytes"
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type Workspaces struct {
	Workspaces []Workspace `json:"items"`
	PageCount  int         `json:"page_count"`
	TotalItems int         `json:"total_items"`
}

type Workspace struct {
	AccountID string          `json:"account_id"`
	Default   bool            `json:"default"`
	Forms     *WorkspaceForms `json:"forms,omitempty"`
	ID        string          `json:"id"`
	Members   *[]Members      `json:"members,omitempty"`
	Name      string          `json:"name"`
	Self      Self            `json:"self"`
	Shared    bool            `json:"shared"`
}
type WorkspaceForms struct {
	Count int    `json:"count"`
	Href  string `json:"href"`
}
type Members struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Role            string `json:"role"`
	AccountMemberID string `json:"account_member_id"`
}
type Self struct {
	Href string `json:"href"`
}
type WorkspaceName struct {
	Name string `json:"name"`
}
type WorkspaceQueryParams struct {
	Search   string `url:"search,omitempty"`
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
}

func (t *Typeform) GetAllWorkspaces(params WorkspaceQueryParams) (*Workspaces, error) {
	queryParams, _ := query.Values(params)
	contents, err := t.buildAndExecRequest("GET", workspacesUrl+"?"+queryParams.Encode(), nil)
	if err != nil {
		return nil, err
	}
	var workspaces *Workspaces
	err = json.Unmarshal(contents, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, err
}

func (t *Typeform) GetAllWorkspacesInAccount(id string, params WorkspaceQueryParams) (*Workspaces, error) {
	queryParams, _ := query.Values(params)
	contents, err := t.buildAndExecRequest("GET", accountsUrl+"/"+id+workspacesUrl+"?"+queryParams.Encode(), nil)
	if err != nil {
		return nil, err
	}
	var workspaces *Workspaces
	err = json.Unmarshal(contents, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, err
}

func (t *Typeform) GetWorkspace(id string) (*Workspace, error) {
	contents, err := t.buildAndExecRequest("GET", workspacesUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	var workspace *Workspace
	err = json.Unmarshal(contents, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, err
}

func (t *Typeform) UpdateWorkspace(id string, updates []UpdateRequestBody) error {
	json_data, err := json.Marshal(updates)
	if err != nil {
		return err
	}
	_, err = t.buildAndExecRequest("PATCH", workspacesUrl+"/"+id, bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	return err
}

func (t *Typeform) CreateWorkspace(name string) (*Workspace, error) {
	requestBody := &WorkspaceName{
		Name: name,
	}
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	contents, err := t.buildAndExecRequest("POST", workspacesUrl, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	var workspace *Workspace
	err = json.Unmarshal(contents, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (t *Typeform) DeleteWorkspace(id string) error {
	_, err := t.buildAndExecRequest("DELETE", workspacesUrl+"/"+id, nil)
	return err
}
