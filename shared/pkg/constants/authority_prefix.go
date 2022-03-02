package constants

type Prefix string

const (
	Prefix_RSS3     Prefix = "rss3://"
	Prefix_Account  Prefix = "account"
	Prefix_Instance Prefix = "instance"
	Prefix_Note     Prefix = "note"
	Prefix_Asset    Prefix = "asset"
)

func IsValidPrefix(prefix string) bool {
	switch Prefix(prefix) {
	case Prefix_Account, Prefix_Instance, Prefix_Note, Prefix_Asset:
		return true
	}

	return false
}
