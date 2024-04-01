package rpcclient

import (
	"encoding/json"
	"fmt"
	"goblock/db/models"
	"goblock/utils"
	"strconv"
)

type BlockHeight struct {
	ID      int64  `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

var (
	reponseRPC  ResponseRPC
	transaction Transaction
)

func InitConfig(bc []models.BlockchainCurrency) *ClientType {
	setting := &ClientType{}
	var params []models.BlockchainCurrency
	for i, v := range bc {
		contracts := optionCurrencies(v)
		if contracts != nil && contracts.GasLimit != "" {
			params[i] = v
		} else {
			setting.Eth = v
		}
	}

	setting.Erc20 = params

	return setting
}

func (service *ClientType) LatestBlockNumber(server string) (int64, error) {
	var params []interface{}
	var blockHeght BlockHeight
	var heightString string

	err := JsonRpc(server, "eth_blockNumber", params, &reponseRPC)
	if err != nil {
		fmt.Println("Failed to marshal object :")
		fmt.Println(err)
		return 0, err
	}

	recordJSON, err := json.Marshal(reponseRPC)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(recordJSON, &blockHeght)
	if err != nil {
		return 0, err
	}

	heightString = blockHeght.Result
	return utils.ConvertFromHex(heightString), nil
}

func (service *ClientType) FetchBlock(server string, height int) (*[]TransactionDetail, error) {
	var params []interface{}
	var response ResponseTransaction
	heightOnHex := utils.ConvertToHex(height)
	params = append(params, fmt.Sprintf("0x%s", heightOnHex))
	params = append(params, true)

	err := JsonRpc(server, "eth_getBlockByNumber", params, &response)
	if err != nil {
		fmt.Printf("Failed to fetch rpc : %s", server)
		fmt.Println(err)
		return nil, err
	}

	recordJSON, err := json.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(recordJSON, &transaction)
	if err != nil {
		return nil, err
	}

	transactionParse, err := service.parsingTransactionDetail(server, transaction)
	if err != nil {
		return nil, err
	}

	return transactionParse, nil
}

func (service *ClientType) parsingTransactionDetail(server string, transactions Transaction) (*[]TransactionDetail, error) {
	var transactionDetail []TransactionDetail

	recordJson, err := json.Marshal(transactions.Transactions)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(recordJson, &transactionDetail)
	if err != nil {
		return nil, err
	}

	for i, v := range transactionDetail {
		if v.Input != "0x" {
			if len(v.Input) < 75 {
				continue
			}

			transactionDetail[i].BlockNumber = strconv.FormatInt(utils.ConvertFromHex(v.BlockNumber), 10)
			transactionDetail[i].Gas = strconv.FormatInt(utils.ConvertFromHex(v.Gas), 10)
			transactionDetail[i].GasPrice = strconv.FormatInt(utils.ConvertFromHex(v.GasPrice), 10)
			transactionDetail[i].Nonce = strconv.FormatInt(utils.ConvertFromHex(v.Nonce), 10)
			transactionDetail[i].ContractAddress = v.To
			transactionDetail[i].To = fmt.Sprintf("0x%s", v.Input[34:74])
			transactionDetail[i].Status = "pending"
			transactionDetail[i].Currency = ""
		} else {
			transactionDetail[i].BlockNumber = strconv.FormatInt(utils.ConvertFromHex(v.BlockNumber), 10)
			transactionDetail[i].Gas = strconv.FormatInt(utils.ConvertFromHex(v.Gas), 10)
			transactionDetail[i].GasPrice = strconv.FormatInt(utils.ConvertFromHex(v.GasPrice), 10)
			transactionDetail[i].Nonce = strconv.FormatInt(utils.ConvertFromHex(v.Nonce), 10)
			transactionDetail[i].Status = "pending"
			transactionDetail[i].Currency = service.Eth.CurrencyId
		}
	}

	return &transactionDetail, nil
}

func (service *ClientType) FetchTransactionReceipt(tx_id string, server string) (*ReceiptTransaction, error) {
	var params []interface{}
	var responseData ResponseRPC
	var resultTransaction ReceiptTransaction
	params = append(params, tx_id)
	err := JsonRpc(server, "eth_getTransactionReceipt", params, &responseData)
	if err != nil {
		return nil, err
	}

	recordJSON, err := json.Marshal(responseData.Result)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(recordJSON, &resultTransaction)
	if err != nil {
		return nil, err
	}

	return &resultTransaction, nil
}

func optionCurrencies(currencies models.BlockchainCurrency) *models.Option {
	var optionData models.Option
	recordJSON, err := json.Marshal(currencies.Options)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(recordJSON, &optionData)
	if err != nil {
		return nil
	}

	return &optionData
}
