package rpcclient

import "goblock/db/models"

type ClientType struct {
	Erc20 []models.BlockchainCurrency
	Eth   models.BlockchainCurrency
}

type ParamsRpc struct {
	ID      int         `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type ResponseRPC struct {
	ID      int64       `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

type ResponseTransaction struct {
	ID      int64       `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

type Transaction struct {
	BaseFeePerGas    string              `omitempty,json:"baseFeePerGas"`
	Difficulty       string              `omitempty,json:"difficulty"`
	Epoch            string              `omitempty,json:"epoch"`
	ExtraData        string              `omitempty,json:"extraData"`
	GasLimit         string              `omitempty,json:"gasLimit"`
	GasUsed          string              `omitempty,json:"gasUsed"`
	Hash             string              `omitempty,json:"hash"`
	LogsBloom        string              `omitempty,json:"logsBloom"`
	Miner            string              `omitempty,json:"miner"`
	MixHash          string              `omitempty,json:"mixHash"`
	Nonce            string              `omitempty,json:"nonce"`
	Number           string              `omitempty,json:"number"`
	ParentHash       string              `omitempty,json:"parentHash"`
	ReceiptsRoot     string              `omitempty,json:"receiptsRoot"`
	Sha3Uncles       string              `omitempty,json:"sha3Uncles"`
	Size             string              `omitempty,json:"size"`
	StateRoot        string              `omitempty,json:"stateRoot"`
	Timestamp        string              `omitempty,json:"timestamp"`
	TimestampNano    string              `omitempty,json:"timestampNano"`
	TotalDifficulty  string              `omitempty,json:"totalDifficulty"`
	Transactions     []TransactionDetail `omitempty,json:"transactions"`
	TransactionsRoot string              `omitempty,json:"transactionRoot"`
}

type TransactionDetail struct {
	Currency         string `omitempty,json:"currency"`
	BlockHash        string `omitempty,json:"blockHash"`
	BlockNumber      string `omitempty,json:"blockNumber"`
	From             string `omitempty,json:"from"`
	Gas              string `omitempty,json:"gas"`
	GasPrice         string `omitempty,json:"gasPrice"`
	Hash             string `omitempty,json:"hash"`
	Nonce            string `omitempty,json:"nonce"`
	To               string `omitempty,json:"to"`
	TransactionIndex string `omitempty,json:"transactionIndex"`
	Type             string `omitempty,json:"type"`
	Status           string `omitempoty,json:"status"`
	Value            string `omitempty,json:"value"`
	ContractAddress  string `omitempty,json:"contractAddress"`
	Input            string `omitempty,json:"input"`
}

type ReceiptTransaction struct {
	BlockHash         string                  `omitempty,json:"blockHash"`
	BlockNumber       string                  `omitempty,json:"blockNumber"`
	ContractAddress   string                  `omitempty,json:"contractAddress"`
	CumulativeGasUsed string                  `omitempty,json:"cumulativeGasUsed"`
	EffectiveGasPrice string                  `omitempty,json:"effectiveGasPrice"`
	From              string                  `omitempty,json:"from"`
	GasUsed           string                  `omitempty,json:"gasUsed"`
	Logs              []ReceiptTransactionLog `omitempty,json:"logs"`
	LogsBloom         string                  `omitempty,json:"logsBloom"`
	Status            string                  `omitempty,json:"status"`
	To                string                  `omitempty,json:"to"`
	TransactionHash   string                  `omitempty,json:"transactionHash"`
	TransactionIndex  string                  `omitempty,json:"transactionIndex"`
	Type              string                  `omitempty,json:"type"`
}

type ReceiptTransactionLog struct {
	Address          string   `omitempty,json:"address"`
	Topics           []string `omitempty,json:"topics"`
	Data             string   `omitempty,json:"data"`
	BlockNumber      string   `omitempty,json:"blockNumber"`
	TransactionHash  string   `omitempty,json:"transactionHash"`
	TransactionIndex string   `omitempty,json:"transactionIndex"`
	BlockHash        string   `omitempty,json:"blockHash"`
	LogIndex         string   `omitempty,json:"logIndex"`
	Removed          bool     `omitempty,json:"removed"`
}
