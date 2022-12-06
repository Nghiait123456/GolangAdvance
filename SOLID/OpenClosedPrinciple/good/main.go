package bad

type Payment struct {
	partnerCode   string
	paymentMethod string
	paymentNow    PaymentNowInterface
}

type PaymentInterface interface {
	PaymentNow() bool
}

func (p *Payment) PaymentNow() bool {
	return p.paymentNow.PayNow()
}

func NewPayment(partnerCode string) PaymentInterface {
	switch partnerCode {
	case "EX_1":
		{
			return &Payment{
				partnerCode: partnerCode,
			}
		}
	case "EX_2":
		{
			return &Payment{
				partnerCode: partnerCode,
			}
		}

	default:
		panic("partnerCode not valid")
	}

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
