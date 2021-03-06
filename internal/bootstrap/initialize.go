package bootstrap

import (
	"fmt"
	"strings"

	vault "github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

func checkInit(client *vault.Client) (bool, error) {
	init, err := client.Sys().InitStatus()
	if err != nil {
		return false, err
	}
	return init, nil
}

func operatorInit(client *vault.Client) (*string, *[]string, error) {

	initReq := &vault.InitRequest{
		SecretShares:    vaultKeyShares,
		SecretThreshold: vaultKeyThreshold,
	}
	initResp, err := client.Sys().Init(initReq)
	if err != nil {
		return nil, nil, err
	}
	log.Info("Vault successfully initialized")
	return &initResp.RootToken, &initResp.Keys, nil
}

// log tokens to K8s log if you don't want to save it in a secret
func logTokens(rootToken *string, unsealKeys *[]string) {
	tokenLog := fmt.Sprintf("Root Token: %s", *rootToken)
	unsealKeysLog := fmt.Sprintf("Unseal Key(s): %s", strings.Join(*unsealKeys, ";"))
	log.Info(tokenLog)
	log.Info(unsealKeysLog)
}
