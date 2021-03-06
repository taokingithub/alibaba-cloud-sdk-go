package codeup

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ResultItem is a nested struct in codeup response
type ResultItem struct {
	Name                  string     `json:"Name" xml:"Name"`
	ExternUserId          string     `json:"ExternUserId" xml:"ExternUserId"`
	OrganizationRole      string     `json:"OrganizationRole" xml:"OrganizationRole"`
	UpdatedAt             string     `json:"UpdatedAt" xml:"UpdatedAt"`
	EnableSslVerification bool       `json:"EnableSslVerification" xml:"EnableSslVerification"`
	VisibilityLevel       int        `json:"VisibilityLevel" xml:"VisibilityLevel"`
	AvatarUrl             string     `json:"AvatarUrl" xml:"AvatarUrl"`
	BranchName            string     `json:"BranchName" xml:"BranchName"`
	WebUrl                string     `json:"WebUrl" xml:"WebUrl"`
	MergeRequestsEvents   bool       `json:"MergeRequestsEvents" xml:"MergeRequestsEvents"`
	PushEvents            bool       `json:"PushEvents" xml:"PushEvents"`
	CreatedAt             string     `json:"CreatedAt" xml:"CreatedAt"`
	NamespaceId           int64      `json:"NamespaceId" xml:"NamespaceId"`
	OrganizationId        string     `json:"OrganizationId" xml:"OrganizationId"`
	Description           string     `json:"Description" xml:"Description"`
	NameWithNamespace     string     `json:"NameWithNamespace" xml:"NameWithNamespace"`
	SshCloneUrl           string     `json:"SshCloneUrl" xml:"SshCloneUrl"`
	ProtectedBranch       bool       `json:"ProtectedBranch" xml:"ProtectedBranch"`
	Path                  string     `json:"Path" xml:"Path"`
	NoteEvents            bool       `json:"NoteEvents" xml:"NoteEvents"`
	Email                 string     `json:"Email" xml:"Email"`
	TagPushEvents         bool       `json:"TagPushEvents" xml:"TagPushEvents"`
	AccessLevel           int        `json:"AccessLevel" xml:"AccessLevel"`
	LastActivityAt        string     `json:"LastActivityAt" xml:"LastActivityAt"`
	PathWithNamespace     string     `json:"PathWithNamespace" xml:"PathWithNamespace"`
	Id                    int64      `json:"Id" xml:"Id"`
	OrganizationName      string     `json:"OrganizationName" xml:"OrganizationName"`
	LastTestResult        string     `json:"LastTestResult" xml:"LastTestResult"`
	SecretToken           string     `json:"SecretToken" xml:"SecretToken"`
	Url                   string     `json:"Url" xml:"Url"`
	Archive               bool       `json:"Archive" xml:"Archive"`
	State                 string     `json:"State" xml:"State"`
	HttpCloneUrl          string     `json:"HttpCloneUrl" xml:"HttpCloneUrl"`
	ProjectId             int64      `json:"ProjectId" xml:"ProjectId"`
	CreatorId             int64      `json:"CreatorId" xml:"CreatorId"`
	ImportStatus          string     `json:"ImportStatus" xml:"ImportStatus"`
	CommitInfo            CommitInfo `json:"CommitInfo" xml:"CommitInfo"`
}
