package constants

type PlatformNameID int32

const (
	PlatformName_Unknown PlatformNameID = 0

	PlatformName_Rss3   PlatformNameID = 1 // not yet so far
	PlatformName_Evm    PlatformNameID = 2
	PlatformName_Solana PlatformNameID = 3
	PlatformName_Flow   PlatformNameID = 4

	PlatformName_Twitter PlatformNameID = 5
	PlatformName_Misskey PlatformNameID = 6
	PlatformName_Jike    PlatformNameID = 7
)

var SignablePlatformNameMap = map[PlatformNameID]string{
	PlatformName_Rss3:   "rss3",
	PlatformName_Evm:    "evm",
	PlatformName_Solana: "solana",
	PlatformName_Flow:   "flow",
}

var UnsignablePlatformNameMap = map[PlatformNameID]string{
	PlatformName_Unknown: "unknown",

	PlatformName_Twitter: "twitter",
	PlatformName_Misskey: "misskey",
	PlatformName_Jike:    "jike",
}

// Converts PlatformName ID to string.
func (id PlatformNameID) String() string {
	if v, ok := SignablePlatformNameMap[id]; ok {
		return v
	}

	if v, ok := UnsignablePlatformNameMap[id]; ok {
		return v
	}

	return "unknown"
}

// Converts string to PlatformName ID.
func StringToPlatformNameID(platformName string) PlatformNameID {
	for k, v := range SignablePlatformNameMap {
		if v == platformName {
			return k
		}
	}

	for k, v := range UnsignablePlatformNameMap {
		if v == platformName {
			return k
		}
	}

	return PlatformName_Unknown
}
