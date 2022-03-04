package types

import "math/big"

type GrantInfo struct {
	Title        string
	AdminAddress string
}

type ProjectInfo struct {
	Active          bool
	Id              int64
	Title           string
	Slug            string
	Description     string
	ReferUrl        string
	Logo            string
	AdminAddress    string
	TokenAddress    string
	TokenSymbol     string
	ContractAddress string
	Network         string
}

type DonationInfo struct {
	Donor          string
	AdminAddress   string
	TokenAddress   string
	Amount         string
	Symbol         string
	FormatedAmount *big.Int
	Decimals       int64
	Timestamp      string
	TxHash         string
	Approach       string
}
