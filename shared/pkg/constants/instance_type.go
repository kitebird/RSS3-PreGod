package constants

type InstanceTypeID int32

const (
	InstanceType_Unknown InstanceTypeID = 0

	InstanceType_Account InstanceTypeID = 1
	InstanceType_Object  InstanceTypeID = 2
	InstanceType_Asset   InstanceTypeID = 3
	InstanceType_Note    InstanceTypeID = 4
)

var InstanceTypeMap = map[InstanceTypeID]string{
	InstanceType_Unknown: "unknown",

	InstanceType_Account: "account",
	InstanceType_Object:  "object",
	InstanceType_Asset:   "asset",
	InstanceType_Note:    "note",
}

// Converts InstanceTypeID to string.
func (id InstanceTypeID) String() string {
	return InstanceTypeMap[id]
}

// Converts string to InstanceTypeID.
func StringToInstanceTypeID(instanceType string) InstanceTypeID {
	for k, v := range InstanceTypeMap {
		if v == instanceType {
			return k
		}
	}

	return InstanceType_Unknown
}
