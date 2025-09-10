package enums


type SubscriptionStatus string

const (
	SubscriptionActive    SubscriptionStatus = "active"       
	SubscriptionCancelled SubscriptionStatus = "cancelled" 
	SubscriptionExpired   SubscriptionStatus = "expired"   
)