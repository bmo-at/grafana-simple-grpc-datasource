package client

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type BackendAPIDatasourceSettings struct {
	ID                          string `json:"-"`
	Endpoint                    string `json:"endpoint"`
	APIKey                      string `json:"-"`
	ApiKeyAuthenticationEnabled bool   `json:"apikey_authentication_enabled"`
}

func (s *BackendAPIDatasourceSettings) Load(config backend.DataSourceInstanceSettings) error {
	if config.JSONData != nil && len(config.JSONData) > 1 {
		if err := json.Unmarshal(config.JSONData, s); err != nil {
			return fmt.Errorf("could not unmarshal DatasourceSettings json: %w", err)
		}
	}
	s.ID = config.UID
	s.APIKey = config.DecryptedSecureJSONData["apiKey"]
	return nil
}
