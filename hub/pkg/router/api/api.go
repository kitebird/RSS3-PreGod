package api

import (
	instance "github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router/api/instance/handler"
	item "github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router/api/item/handler"
)

var GetInstance = instance.GetInstance

var GetItem = item.GetItem
var GetItemPagedList = item.GetItemPagedList

// var GetItemList = instance.GetItemList
// var GetLinkList = instance.GetLinkList
// var GetBacklinkList = instance.GetBacklinkList
