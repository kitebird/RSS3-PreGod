package types

import "fmt"

// MoralisNFTItem store all indexed NFTs from moralis api.
type MoralisNFTItem struct {
	TokenAddress      string `json:"token_address"`
	TokenId           string `json:"token_id"`
	BLockNumberMinted string `json:"block_number_minted"`
	OwnerOf           string `json:"owner_of"`
	BlockNumber       string `json:"block_number"`
	Amount            string `json:"amount"`
	ContractType      string `json:"contract_type"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	TokenURI          string `json:"token_uri"`
	MetaData          string `json:"metadata"`
	SyncedAt          string `json:"synced_at"`
	IsValid           int64  `json:"is_valid"`
	Syncing           int64  `json:"syncing"`
	Frozen            int64  `json:"frozen"`
}

type MoralisNFTResult struct {
	Total    int64            `json:"total"`
	Page     int64            `json:"page"`
	PageSize int64            `json:"page_size"`
	Result   []MoralisNFTItem `json:"result"`
	Status   string           `json:"status"`
}

func (i MoralisNFTItem) String() string {
	return fmt.Sprintf(`TokenAddress: %s, TokenId: %s, OwnerOf: %s, TokenURI: %s`,
		i.TokenAddress, i.TokenId, i.OwnerOf, i.TokenURI)
}

// MoralisNFTTransferItem store the transfers of NFTS.
type MoralisNFTTransferItem struct {
	BlockNumber      string `json:"block_number"`
	BlockTimestamp   string `json:"block_timestamp"`
	BlockHash        string `json:"block_hash"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionIndex int64  `json:"transaction_index"`
	LogIndex         int64  `json:"log_index"`
	Value            string `json:"value"`
	ContractType     string `json:"contract_type"`
	TransactionType  string `json:"transaction_type"`
	TokenAddress     string `json:"token_address"`
	TokenId          string `json:"token_id"`
	FromAddress      string `json:"from_address"`
	ToAddress        string `json:"to_address"`
	Amount           string `json:"amount"`
	Verified         int64  `json:"verified"`
	Operator         string `json:"operator"`
}

type MoralisNFTTransferResult struct {
	Total       int64                    `json:"total"`
	Page        int64                    `json:"page"`
	PageSize    int64                    `json:"page_size"`
	Result      []MoralisNFTTransferItem `json:"result"`
	Cursor      string                   `json:"cursor"`
	BlockExists bool                     `json:"block_exists"`
}

func (i MoralisNFTTransferItem) String() string {
	return fmt.Sprintf(`From: %s, To: %s, TokenAddress: %s, ContractType: %s, TokenId: %s`,
		i.FromAddress, i.ToAddress, i.TokenAddress, i.ContractType, i.TokenId)
}

type MoralisGetLogsItem struct {
	TransactionHash string `json:"transaction_hash"`
	Address         string `json:"address"`
	BlockTimestamp  string `json:"block_timestamp"`
	BlockNumber     string `json:"block_number"`
	BlockHash       string `json:"block_hash"`
	Data            string `json:"data"`
	Topic0          string `json:"topic0"`
	Topic1          string `json:"topic1"`
	Topic2          string `json:"topic2"`
	Topic3          string `json:"topic3"`
}

func (i MoralisGetLogsItem) String() string {
	return fmt.Sprintf(`TransactionHash: %s, Address: %s, Data: %s, Topic0: %s, Topic1: %s, Topic2:%s, Topic3: %s`,
		i.TransactionHash, i.Address, i.Data, i.Topic0, i.Topic1, i.Topic2, i.Topic3)
}

type MoralisGetLogsResult struct {
	Total    int64                `json:"total"`
	Page     int64                `json:"page"`
	PageSize int64                `json:"page_size"`
	Result   []MoralisGetLogsItem `json:"result"`
}
