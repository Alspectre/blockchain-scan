package rpcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonRpc(server string, method string, params interface{}, target interface{}) error {
	body := &ParamsRpc{
		ID:      1,
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Failed to marshal object :")
		fmt.Println(body)
		return err
	}

	requestClient, err := http.Post(server, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	defer requestClient.Body.Close()

	err = json.NewDecoder(requestClient.Body).Decode(&target)
	if err != nil {
		return err
	}

	return nil
}
