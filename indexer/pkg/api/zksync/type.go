package zksync

import "fmt"

type Token struct {
	Id       int64  `json:"id"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
	Kind     string `json:"kind"`
	IsNFT    bool   `json:"is_nft"`
}

func (t Token) String() string {
	return fmt.Sprintf(`Id: %d, Address: %s, Symbol: %s, Decimals: %d, Kind: %s, IsNFT: %v`,
		t.Id, t.Address, t.Symbol, t.Decimals, t.Kind, t.IsNFT)
}

type Op struct {
	To        string `json:"to"`
	Fee       string `json:"fee"`
	From      string `json:"from"`
	Type      string `json:"type"`
	Nonce     int64  `json:"nonce"`
	Token     int64  `json:"token"`
	Amount    string `json:"amount"`
	AccountId int64  `json:"accountId"` // nolint:tagliatelle // accountId is returned by zksync api
}

func (o Op) String() string {
	return fmt.Sprintf(`From: %s, To: %s, Type: %s, Token: %d, Amount: %s`,
		o.From, o.To, o.Type, o.Token, o.Amount)
}

type Transaction struct {
	TxHash      string `json:"tx_hash"`
	BlockNumber int64  `json:"block_number"`
	Op          Op     `json:"op"`
	Success     bool   `json:"success"`
	FailReason  string `json:"fail_reason"`
	CreatedAt   string `json:"created_at"`
	BatchId     string `json:"batch_id"`
}
