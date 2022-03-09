package mq

const (
	Namespace = "rss3-pre-god-ns"
	GroupName = "rss3-pre-god-group"
)

type TopicName string

const (
	TopicNameUpdateAssetDetail TopicName = "asset-detail"
	TopicNameGiveAssetDetail   TopicName = "give-asset-detail"
)
