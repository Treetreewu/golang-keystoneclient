// Copyright 2017 EasyStack, Inc.

package model

type authIdentity struct {
	Methods  []string            `json:"methods"`
	Password *PasswordCredential `json:"password,omitempty"`
	Token    *TokenCredential    `json:"token,omitempty"`
}

type auth struct {
	Identity authIdentity `json:"identity"`
	Scope    Scope        `json:"scope,omitempty"`
}

type TokenRequestData struct {
	Auth auth `json:"auth"`
}

type IDOrName struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type domain struct {
	IDOrName
}

type Scope interface{}

type ExplicitUnScoped = string

type DomainScope struct {
	Scope  `json:",omitempty,inline"`
	Domain domain `json:"domain"`
}

func NewDomainScope(id *string, name *string) *DomainScope {
	return &DomainScope{
		Domain: domain{
			IDOrName{id, name},
		},
	}
}

type system struct {
	All bool `json:"all"`
}

type SystemScope struct {
	Scope  `json:",omitempty,inline"`
	System system `json:"system"`
}

func NewSystemScope() *SystemScope {
	return &SystemScope{
		System: system{true},
	}
}

type project struct {
	IDOrName
	Domain domain `json:"domain,omitempty"`
}

type ProjectScope struct {
	Scope   `json:",omitempty,inline"`
	Project project `json:"project"`
}

func NewProjectScope(
	projectID *string, ProjectName *string,
	domainID *string, domainName *string,
) *ProjectScope {
	return &ProjectScope{
		Project: project{
			IDOrName: IDOrName{projectID, ProjectName},
			Domain: domain{
				IDOrName{domainID, domainName},
			},
		},
	}
}

type Credential interface{}

type user struct {
	IDOrName
	Email    *string `json:"email,omitempty"`
	Password string  `json:"password"`
	Domain   domain  `json:"domain"`
}

type PasswordCredential struct {
	Credential `json:",omitempty,inline"`
	User       user `json:"user"`
}

// NewPasswordCredential creates a password credential.
func NewPasswordCredential(
	userID *string, userName *string, password string,
	domainID *string, domainName *string,
) *PasswordCredential {
	return &PasswordCredential{
		User: user{
			IDOrName: IDOrName{userID, userName},
			Password: password,
			Domain: domain{
				IDOrName{domainID, domainName},
			},
		},
	}
}

// NewPasswordCredentialByEmail creates a password credential by email.
// This feature is supported by easystack cloud only.
func NewPasswordCredentialByEmail(
	email string, password string,
	domainID *string, domainName *string,
) *PasswordCredential {
	return &PasswordCredential{
		User: user{
			Email:    &email,
			Password: password,
			Domain: domain{
				IDOrName{domainID, domainName},
			},
		},
	}
}

type TokenCredential struct {
	Credential `json:",omitempty,inline"`
	ID         string `json:"id"`
}

func NewTokenCredential(id string) *TokenCredential {
	return &TokenCredential{ID: id}
}
