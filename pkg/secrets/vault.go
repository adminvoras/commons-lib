package secrets

import (
	"context"
	"fmt"
	"time"

	"github.com/adminvoras/commons-lib/pkg/errors"
	"github.com/hashicorp/vault-client-go"
)

type Secret interface {
	Get(key string) (string, error)
}

type secret struct {
	client  *vault.Client
	project string
	scope   string
}

func NewSecrets(token, project, scope string) (Secret, error) {
	if token == "" {
		return nil, errors.New(nil, "token cannot be nil")
	}

	if project == "" {
		return nil, errors.New(nil, "project cannot be nil")
	}

	if scope == "" {
		return nil, errors.New(nil, "scope cannot be nil")
	}

	client, err := vault.New(
		vault.WithAddress("http://149.50.139.210:8200"),
		vault.WithRequestTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}

	if err = client.SetToken(token); err != nil {
		return nil, err
	}

	return &secret{
		client:  client,
		project: project,
		scope:   scope,
	}, nil
}

func (s secret) Get(key string) (string, error) {
	resp, err := s.client.Read(context.Background(), fmt.Sprintf("/%s/%s", s.project, s.scope))
	if err != nil {
		return "", errors.New(err, "error reading secret")
	}

	data, ok := resp.Data[key]
	if !ok {
		return "", errors.New(nil, "not map interface")
	}

	return fmt.Sprintf("%s", data), nil
}
