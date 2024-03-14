package vaultconfig

import (
	"bytes"
	"fmt"
	"goblock/utils"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type VaultInterface struct {
	Initialized bool `omitempty,json:"initialized"`
	Sealed      bool `omitempty,json:"sealed"`
	Shares      int  `omitempty,json:"shares"`
	Threshold   int  `omitempty,json:"threshold"`
}

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

func Setup() {
	init := VaultExec("vault status -format yaml")
	convert := convert(init)
	vaultSecretsPath := "vault_secrets.yml"

	var unseal_key map[string]interface{}

	if !convert.Initialized {
		fmt.Println("============= Initialization START ===============")
		vaultInit := VaultExec("vault operator init -format yaml --recovery-shares=3 --recovery-threshold=2")

		err := ioutil.WriteFile(vaultSecretsPath, vaultInit.Bytes(), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		err = yaml.Unmarshal(vaultInit.Bytes(), &unseal_key)
		if err != nil {
			fmt.Println("Error parsing YAML:", err)
			return
		}
		fmt.Println("============== Initialization END ===============")
	} else {
		vaultSecrets, err := ioutil.ReadFile(vaultSecretsPath)
		if err != nil {
			fmt.Println("Vault keys are missing")
			return
		}
		err = yaml.Unmarshal(vaultSecrets, &unseal_key)
		if err != nil {
			fmt.Println("Error parsing YAML:", err)
			return
		}

	}

	vaultRootToken := unseal_key["root_token"]
	unsealKeys := unseal_key["unseal_keys_b64"]

	if convert.Sealed {
		fmt.Println("============= Unsealing START ===============")
		fmt.Println(vaultRootToken)
		fmt.Println(unsealKeys)
		fmt.Println("============== Unsealing END ===============")
	}
}

func VaultExec(command string) bytes.Buffer {
	cmd := exec.Command("docker-compose", "exec", "-T", "vault", "sh", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return out
	}

	return out
}

func convert(source bytes.Buffer) *VaultInterface {
	var status VaultInterface

	err := yaml.Unmarshal(source.Bytes(), &status)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	fmt.Println("Docker Compose command executed successfully")

	return &status
}
