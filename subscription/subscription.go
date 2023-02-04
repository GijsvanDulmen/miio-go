package subscription

import (
	"github.com/GijsvanDulmen/miio-go/subscription/common"
	"github.com/GijsvanDulmen/miio-go/subscription/target"
)

func NewTarget() common.SubscriptionTarget {
	return target.NewTarget()
}

type SubscriptionTarget = common.SubscriptionTarget
type Subscription = common.Subscription
