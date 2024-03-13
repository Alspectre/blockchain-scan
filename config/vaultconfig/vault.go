package vaultconfig

import (
	"fmt"
	"goblock/utils"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)

func InitVault() (*api.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	address := os.Getenv("VAULT_ADDRESS")

	client, err := api.NewClient(&api.Config{
		Address: address,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func DecryptValue(client *api.Client, value string) (string, error) {
	// Perform decryption using Vault's transit secret engine
	secret, err := client.Logical().Write("transit/decrypt/backendexchange_blockchains_server", map[string]interface{}{
		"ciphertext": value,
	})
	if err != nil {
		return "", err
	}
	if secret == nil || secret.Data == nil {
		return "", fmt.Errorf("decryption failed")
	}

	decodeString, err := utils.Base64Decode(secret.Data["plaintext"].(string))
	if err != nil {
		return "", fmt.Errorf(decodeString)
	}
	return decodeString, nil
}

func ReadSecret(client *api.Client, path string) (map[string]interface{}, error) {
	secret, err := client.Logical().Read(path)
	if err != nil {
		return nil, err
	}
	if secret == nil || secret.Data == nil {
		return nil, fmt.Errorf("secret not found at %s", path)
	}
	return secret.Data, nil
}
