// Copyright 2017 EasyStack, Inc.

package model

type TokenRequestData struct {
	Auth struct {
		Identity struct {
			Methods  []string            `json:"methods"`
			Password *PasswordCredential `json:"password,omitempty"`
			Token    *TokenCredential    `json:"token,omitempty"`
		} `json:"identity"`
		Scope Scope `json:"scope,omitempty"`
	} `json:"auth"`
}

type IDOrName struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Scope interface{}

type ExplicitUnScoped = string

type DomainScope struct {
	Domain struct {
		IDOrName
	} `json:"domain"`
}

func NewDomainScope(id *string, name *string) *DomainScope {
	return &DomainScope{
		Domain: struct{ IDOrName }{
			IDOrName: IDOrName{
				ID:   id,
				Name: name,
			},
		},
	}
}

type SystemScope struct {
	System struct {
		All bool `json:"all"`
	} `json:"system"`
}

func NewSystemScope() *SystemScope {
	return &SystemScope{System: struct {
		All bool `json:"all"`
	}(struct{ All bool }{All: true})}
}

type ProjectScope struct {
	Project struct {
		IDOrName
		Domain struct {
			IDOrName
		} `json:"domain,omitempty"`
	} `json:"project"`
}

func NewProjectScope(
	projectID *string, ProjectName *string,
	domainID *string, domainName *string,
) *ProjectScope {
	return &ProjectScope{
		Project: struct {
			IDOrName
			Domain struct{ IDOrName } `json:"domain,omitempty"`
		}(struct {
			IDOrName
			Domain struct{ IDOrName }
		}{
			IDOrName: IDOrName{
				ID:   projectID,
				Name: ProjectName,
			},
			Domain: struct{ IDOrName }{
				IDOrName: IDOrName{
					ID:   domainID,
					Name: domainName,
				},
			},
		},
		),
	}
}

type Credential interface{}

type PasswordCredential struct {
	User struct {
		IDOrName
		Password string `json:"password"`
		Domain   struct {
			IDOrName
		} `json:"domain"`
	} `json:"user"`
}

func NewPasswordCredential(
	userID *string, userName *string, password string,
	domainID *string, domainName *string,
) *PasswordCredential {
	return &PasswordCredential{
		User: struct {
			IDOrName
			Password string             `json:"password"`
			Domain   struct{ IDOrName } `json:"domain"`
		}(struct {
			IDOrName
			Password string
			Domain   struct{ IDOrName }
		}{
			IDOrName: IDOrName{
				ID:   userID,
				Name: userName,
			}, Password: password, Domain: struct{ IDOrName }{
				IDOrName: IDOrName{
					ID:   domainID,
					Name: domainName,
				},
			},
		}),
	}
}

type TokenCredential struct {
	ID string `json:"id"`
}

func NewTokenCredential(id string) *TokenCredential {
	return &TokenCredential{ID: id}
}
