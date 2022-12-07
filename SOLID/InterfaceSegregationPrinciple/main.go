package InterfaceSegregationPrinciple

type CardRisk struct {
}

type CardRiskInterface interface {
	AmountCardRiskInterface
	CustomerOrderRiskInterface
	ProductCardRiskInterface
	CheckRiskFrTokenCard() bool
}

type AmountCardRisk struct {
}
type AmountCardRiskInterface interface {
	CheckAmountRisk(amount uint64) bool
}

type CustomerOrderRisk struct {
}
type CustomerOrderRiskInterface interface {
	CheckCustomerOrderRisk(amount uint64) bool
}

type ProductCardRisk struct {
}
type ProductCardRiskInterface interface {
	CheckProductCardRisk(amount uint64) bool
}
