package constants

type LinkTypeID int32

const (
	LinkType_Unknown LinkTypeID = 0

	LinkType_Following LinkTypeID = 1
	LinkType_Comment  LinkTypeID = 2
	LinkType_Like   LinkTypeID = 3
	LinkType_Collection   LinkTypeID = 4
)

var LinkTypeMap = map[LinkTypeID]string{
	LinkType_Unknown: "unknown",

	LinkType_Following:   "following",
	LinkType_Comment:   "comment",
	LinkType_Like:   "like",
	LinkType_Collection:   "collection",
}

// Converts LinkTypeID to string.
func (id LinkTypeID) String() string {
	return LinkTypeMap[id]
}

// Converts string to LinkTypeID.
func StringToLinkTypeID(LinkType string) LinkTypeID {
	for k, v := range LinkTypeMap {
		if v == LinkType {
			return k
		}
	}

	return LinkType_Unknown
}
