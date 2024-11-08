package secrets

import (
	"context"
	"fmt"
	"time"

	"github.com/adminvoras/commons-lib/pkg/errors"
	"github.com/hashicorp/vault-client-go"
)

type Secret interface {
	Get(ctx context.Context, key, project, scope string) (string, error)
}

type secret struct {
	client *vault.Client
}

func NewSecrets(token string) (Secret, error) {
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
		client: client,
	}, nil
}

func (s secret) Get(ctx context.Context, key, project, scope string) (string, error) {
	resp, err := s.client.Read(ctx, fmt.Sprintf("/%s/%s", project, scope))
	if err != nil {
		return "", errors.New(err, "error reading secret")
	}

	data, ok := resp.Data[key]
	if !ok {
		return "", errors.New(nil, "not map interface")
	}

	return fmt.Sprintf("%s", data), nil
}
