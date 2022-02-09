package types

import "fmt"

type NFTItem struct {
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

type NFTResult struct {
	Total    int64     `json:"total"`
	Page     int64     `json:"page"`
	PageSize int64     `json:"page_size"`
	Result   []NFTItem `json:"result"`
	Status   string    `json:"status"`
}

func (i NFTItem) String() string {
	return fmt.Sprintf(`TokenAddress: %s, TokenId: %s, OwnerOf: %s, TokenURI: %s`,
		i.TokenAddress, i.TokenId, i.OwnerOf, i.TokenURI)
}
