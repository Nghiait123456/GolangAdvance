package bad

type Payment struct {
	paymentMethod string
	paymentNow    PaymentNowInterface
}

type PaymentInterface interface {
	PaymentNow() bool
}

func (p *Payment) PaymentNow() bool {
	return p.paymentNow.PayNow()
}

/////////////////////////////////////////////////////////////////////////////

type PaymentNowInterface interface {
	PayNow() bool
}

type PaypalPayNow struct {
	paymentMethod string
}

func (pp *PaypalPayNow) PayNow() bool {
	// todo action paynow paypal
	return true
}

type VisaPayNow struct {
	paymentMethod string
}

func (pp *VisaPayNow) PayNow() bool {
	// todo action paynow paypal
	return true
}

type MasterCardPayNow struct {
	paymentMethod string
}

func (pp *MasterCardPayNow) PayNow() bool {
	// todo action paynow paypal
	return true
}
