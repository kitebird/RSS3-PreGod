package uuid

import (
	"bytes"
	"strconv"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/google/uuid"
)

var NAMESPACE = uuid.Must(uuid.Parse("6ba7b815-9dad-11d1-80b4-00c04fd430c8"))

func NewAccountUUID(
	identity string,
	platformID constants.PlatformID,
	createdAt string, // timestamp in milliseconds
) uuid.UUID {
	var buf bytes.Buffer

	buf.WriteString(strconv.Itoa(int(constants.PrefixIDAccount)))
	buf.WriteString(identity)
	buf.WriteString(strconv.Itoa(int(platformID)))
	buf.WriteString(createdAt)

	uuid := genUUID(buf.Bytes())

	return uuid
}

func NewNoteUUID(
	ownerIdentity string,
	ownerPlatformID constants.PlatformID,
	noteNetworkID constants.NetworkID,
	createdAt string, // timestamp in milliseconds
) uuid.UUID {
	var buf bytes.Buffer

	buf.WriteString(strconv.Itoa(int(constants.PrefixIDNote)))
	buf.WriteString(ownerIdentity)
	buf.WriteString(strconv.Itoa(int(noteNetworkID)))
	buf.WriteString(createdAt)

	uuid := genUUID(buf.Bytes())

	return uuid
}

func NewAssetUUID(
	ownerIdentity string,
	ownerPlatformID constants.PlatformID,
	assetNetworkID constants.NetworkID,
	createdAt string, // timestamp in milliseconds
) uuid.UUID {
	var buf bytes.Buffer

	buf.WriteString(strconv.Itoa(int(constants.PrefixIDAsset)))
	buf.WriteString(ownerIdentity)
	buf.WriteString(strconv.Itoa(int(assetNetworkID)))
	buf.WriteString(createdAt)

	uuid := genUUID(buf.Bytes())

	return uuid
}

func NewLinkUUID(
	linkTypeID constants.LinkTypeID,
	subjectIdentity string,
	subjectPrefixID constants.PrefixID,
	subjectSuffixID string, // TODO: define type for suffixID
	objectIdentity string,
	objectPrefixID constants.PrefixID,
	objectSuffixID string,
) uuid.UUID {
	var buf bytes.Buffer

	buf.WriteString(strconv.Itoa(int(constants.PrefixIDLink)))
	buf.WriteString(strconv.Itoa(int(linkTypeID)))
	buf.WriteString(subjectIdentity)
	buf.WriteString(strconv.Itoa(int(subjectPrefixID)))
	buf.WriteString(subjectSuffixID)
	buf.WriteString(objectIdentity)
	buf.WriteString(strconv.Itoa(int(objectPrefixID)))
	buf.WriteString(objectSuffixID)

	uuid := genUUID(buf.Bytes())

	return uuid
}

func genUUID(bytes []byte) uuid.UUID {
	uuid := uuid.NewSHA1(NAMESPACE, bytes)

	return uuid
}
