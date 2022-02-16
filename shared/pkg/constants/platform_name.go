package constants

type PlatformNameID int32
type PlatformName string

const (
	PlatformNameID_Unknown PlatformNameID = 0

	PlatformNameID_Rss3   PlatformNameID = 1 // not yet so far
	PlatformNameID_Evm    PlatformNameID = 2
	PlatformNameID_Solana PlatformNameID = 3
	PlatformNameID_Flow   PlatformNameID = 4

	PlatformNameID_Twitter PlatformNameID = 5
	PlatformNameID_Misskey PlatformNameID = 6
	PlatformNameID_Jike    PlatformNameID = 7
)

const (
	PlatformName_Unknown PlatformName = "unknown"

	PlatformName_Rss3   PlatformName = "rss3"
	PlatformName_Evm    PlatformName = "evm"
	PlatformName_Solana PlatformName = "solana"
	PlatformName_Flow   PlatformName = "flow"

	PlatformName_Twitter PlatformName = "twitter"
	PlatformName_Misskey PlatformName = "misskey"
	PlatformName_Jike    PlatformName = "jike"
)

var signablePlatformNameMap = map[PlatformNameID]PlatformName{
	PlatformNameID_Rss3:   PlatformName_Rss3,
	PlatformNameID_Evm:    PlatformName_Evm,
	PlatformNameID_Solana: PlatformName_Solana,
	PlatformNameID_Flow:   PlatformName_Flow,
}

var unsignablePlatformNameMap = map[PlatformNameID]PlatformName{
	PlatformNameID_Unknown: PlatformName_Unknown,

	PlatformNameID_Twitter: PlatformName_Twitter,
	PlatformNameID_Misskey: PlatformName_Misskey,
	PlatformNameID_Jike:    PlatformName_Jike,
}

// Converts PlatformName ID to string.
func (id PlatformNameID) String() PlatformName {
	if v, ok := signablePlatformNameMap[id]; ok {
		return v
	}

	if v, ok := unsignablePlatformNameMap[id]; ok {
		return v
	}

	return PlatformName_Unknown
}

// Checks if the platform is signable.
func (id PlatformNameID) IsSignable() bool {
	_, ok := signablePlatformNameMap[id]

	return ok
}

// Checks if a platform name is valid.
func IsValidPlatformName(platformName string) bool {
	for _, v := range signablePlatformNameMap {
		if string(v) == platformName {
			return true
		}
	}

	for _, v := range unsignablePlatformNameMap {
		if string(v) == platformName {
			return true
		}
	}

	return false
}

// Converts string to PlatformName ID.
func StringToPlatformNameID(platformName string) PlatformNameID {
	for k, v := range signablePlatformNameMap {
		if string(v) == platformName {
			return k
		}
	}

	for k, v := range unsignablePlatformNameMap {
		if string(v) == platformName {
			return k
		}
	}

	return PlatformNameID_Unknown
}
