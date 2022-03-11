package uuid

import (
	"bytes"
	"strconv"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/google/uuid"
)

var space = uuid.NameSpaceOID

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

	uuid := uuid.NewSHA1(space, buf.Bytes())

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

	uuid := uuid.NewSHA1(space, buf.Bytes())

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

	uuid := uuid.NewSHA1(space, buf.Bytes())

	return uuid
}

func NewLinkUUID(
	ownerIdentity string,
	ownerPlatformID constants.PlatformID,
	linkNetworkID constants.NetworkID,
	createdAt string, // timestamp in milliseconds
) uuid.UUID {
	var buf bytes.Buffer

	buf.WriteString(strconv.Itoa(int(constants.PrefixIDLink)))
	buf.WriteString(ownerIdentity)
	buf.WriteString(strconv.Itoa(int(linkNetworkID)))
	buf.WriteString(createdAt)

	uuid := genUUID(buf.Bytes())

	return uuid
}

func genUUID(bytes []byte) uuid.UUID {
	uuid := uuid.NewSHA1(space, bytes)

	return uuid
}
