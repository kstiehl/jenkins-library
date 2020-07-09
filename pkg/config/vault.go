package config

import (
	"path"

	"github.com/SAP/jenkins-library/pkg/vault"
	"github.com/hashicorp/vault/api"
)

// vaultClient interface for mocking
type vaultClient interface {
	GetKvSecret(string) (map[string]string, error)
}

func getVaultClientFromConfig(config StepConfig) (vaultClient, error) {
	address, addressOk := config.Config["vaultAddress"].(string)
	basePath, basePathOk := config.Config["vaultRootPath"].(string)
	token, tokenOk := config.Config["vaultToken"].(string)

	// if vault isn't used it's not an error
	if !addressOk || address == "" || !basePathOk || basePath == "" || !tokenOk || token == "" {
		return nil, nil
	}

	// namespaces are only available in vault enterprise so using them should be optional
	namespace := config.Config["vaultNamespace"].(string)

	client, err := vault.NewClient(&api.Config{Address: address}, token, namespace)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func getVaultConfig(client vaultClient, config StepConfig, params []StepParameters) (map[string]interface{}, error) {
	vaultConfig := map[string]interface{}{}
	for _, param := range params {

		// we don't overwrite secrets that have already been set in any way
		if _, ok := config.Config[param.Name].(string); ok {
			continue
		}
		for _, ref := range param.GetReferences("vaultSecret") {
			// it should be possible to configure the root path were the secret is stored
			basePath := config.Config[ref.Name].(string)
			secret, err := client.GetKvSecret(path.Join(basePath, ref.Path))
			if err != nil {
				return nil, err
			}
			if secret == nil {
				continue
			}

			field := secret[param.Name]
			if field != "" {
				vaultConfig[param.Name] = field
				break
			}
		}
	}
	return vaultConfig, nil
}
