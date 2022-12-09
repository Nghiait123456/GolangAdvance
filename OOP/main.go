package main

type PayNowInterface interface {
	PayNow() bool
}

type Payment struct {
	PayNow PayNowInterface
}

type VisaPayNow struct {
}

func (v *VisaPayNow) PayNow() bool {
	//todo action
	return true
}

type MasterCardPayNow struct {
}

func (v *MasterCardPayNow) PayNow() bool {
	//todo action
	return true
}

func main() {
	v := VisaPayNow{}
	m := MasterCardPayNow{}

	pV := Payment{
		PayNow: &v,
	}

	pM := Payment{
		PayNow: &m,
	}

	pV.PayNow.PayNow()
	pM.PayNow.PayNow()
}
