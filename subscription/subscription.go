package subscription

import (
	"github.com/danilarff86/miio-go/subscription/common"
	"github.com/danilarff86/miio-go/subscription/target"
)

func NewTarget() common.SubscriptionTarget {
	return target.NewTarget()
}

type SubscriptionTarget = common.SubscriptionTarget
type Subscription = common.Subscription
