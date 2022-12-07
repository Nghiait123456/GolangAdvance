package bad

type Payment struct {
	paymentMethod string
}

func (p *Payment) GetNow() bool {
	if p.paymentMethod == "Visa" {
		p := VisaPayNow{
			paymentMethod: p.paymentMethod,
		}

		return p.PayNowVisa()
	}

	if p.paymentMethod == "Paypal" {
		p := PaypalPayNow{
			paymentMethod: p.paymentMethod,
		}

		return p.PayNowPayPal()
	}

	return false
}

type PaypalPayNow struct {
	paymentMethod string
}

func (pp *PaypalPayNow) PayNowPayPal() bool {
	// todo action paynow paypal
	return true
}

type VisaPayNow struct {
	paymentMethod string
}

func (pp *VisaPayNow) PayNowVisa() bool {
	// todo action paynow paypal
	return true
}

type MasterCardPayNow struct {
	paymentMethod string
}

func (pp *MasterCardPayNow) PayNowMasterCard() bool {
	// todo action paynow paypal
	return true
}
