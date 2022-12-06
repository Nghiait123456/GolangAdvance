package SingleResponsibilityPrinciple

type BalanceCalculatorNotGood struct {
	balance     uint64
	partnerCode string
}

func (b *BalanceCalculatorNotGood) PlusBalance(plus uint64) {
	b.balance = b.balance + plus
}

func (b *BalanceCalculatorNotGood) ReduceBalance(reduce uint64) {
	b.balance = b.balance - reduce
}

func (b *BalanceCalculatorNotGood) CheckRiskPartner(partnerCode string) bool {
	// todo implement code check risk Partner
	return true
}

type BalanceCalculatorGood struct {
	balance     uint64
	partnerCode string
}

func (b *BalanceCalculatorGood) PlusBalance(plus uint64) {
	b.balance = b.balance + plus
}

func (b *BalanceCalculatorGood) ReduceBalance(reduce uint64) {
	b.balance = b.balance - reduce
}

type PartnerRisk struct {
	partnerCode string
}

func (p *PartnerRisk) CheckRiskPartner(partnerCode string) bool {
	// todo actions check risk
	return true
}
