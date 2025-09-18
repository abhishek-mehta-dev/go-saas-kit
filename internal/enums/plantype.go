package enums

type PlanType string


const (
	PlanStandard PlanType = "standard"
	PlanPremium  PlanType = "premium"
)


type BillingCycle string

const (
	Monthly BillingCycle = "monthly"
	Yearly BillingCycle = "yearly"
)

