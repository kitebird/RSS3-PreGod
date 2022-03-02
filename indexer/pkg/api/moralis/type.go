package moralis

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type MoralisChainType string

const (
	Unknown MoralisChainType = "unknown"

	ETH     MoralisChainType = "eth"
	BSC     MoralisChainType = "bsc"
	Polygon MoralisChainType = "polygon"
	AVAX    MoralisChainType = "avalanche"
	Fantom  MoralisChainType = "fantom"
)

func GetChainType(network constants.NetworkName) MoralisChainType {
	switch network {
	case constants.NetworkName_Ethereum:
		return ETH
	case constants.NetworkName_BNB:
		return BSC
	case constants.NetworkName_Polygon:
		return Polygon
	case constants.NetworkName_Avalanche:
		return AVAX
	case constants.NetworkName_Fantom:
		return Fantom
	default:
		return Unknown
	}
}

func (mt MoralisChainType) GetNFTItemTypeID() constants.ItemTypeID {
	switch mt {
	case "ETH":
		return constants.ItemType_Ethereum_Nft
	case "BSC":
		return constants.ItemType_Bsc_Nft
	case "Polygon":
		return constants.ItemType_Polygon_Nft
	case "AVAX":
		return constants.ItemType_Avax_Nft
	case "Fantom":
		return constants.ItemType_Fantom_Nft
	default:
		return constants.ItemType_Unknown
	}
}
