package model

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/kamva/mgm/v3"
)

type Attachment struct {
	Content    string   `json:"content" bson:"content"`
	Address    []string `json:"address" bson:"address"`
	MimeType   string   `json:"mime_type" bson:"mime_type"`
	Type       string   `json:"type" bson:"type"`
	SizeInByte int      `json:"size_in_bytes" bson:"size_in_bytes"`
}

type Object struct {
	mgm.DefaultModel `bson:",inline"`

	Uid string `json:"uid" bson:"uid"` // Index: (Uid, ItemType)

	Tags        constants.ItemTags `json:"tags" bson:"tags"`
	Authors     []string           `json:"authors" bson:"authors"`
	Title       string             `json:"title" bson:"title"`
	Summary     string             `json:"summary" bson:"summary"`
	Attachments []Attachment       `json:"attachments" bson:"attachments"`
}

func NewAttachment(content string, address []string, mimetype string, t string, size_in_bytes int) *Attachment {
	return &Attachment{
		Content:    content,
		Address:    address,
		MimeType:   mimetype,
		Type:       t,
		SizeInByte: size_in_bytes,
	}
}

func NewObject(
	authors []string,
	uid string,
	title string,
	summary string,
	tags constants.ItemTags,
	attachments []Attachment,
) *Object {
	return &Object{
		Authors:     authors,
		Uid:         uid,
		Title:       title,
		Summary:     summary,
		Tags:        tags,
		Attachments: attachments,
	}
}
