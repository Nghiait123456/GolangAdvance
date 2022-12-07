todo:
https://s8sg.medium.com/solid-principle-in-go-e1a624290346
https://dave.cheney.net/2016/08/20/solid-go-design
https://web.archive.org/web/20150906155800/http://www.objectmentor.com/resources/articles/Principles_and_Patterns.pdf

- [Preview](#Preview)
- [Single Responsibility Principle](#SingleResponsibilityPrinciple)
- [Open/Closed Principle](#OpenClosedPrinciple)
- [Liskov Substitution Principle](#LiskovSubstitutionPrinciple)

## Preview <a name="Preview"></a>

In 2002 Robert Martin published his book, Agile Software Development, Principles, Patterns, and Practices. In it he
describes five principles of reusable software design, which he called the SOLID principles, after the first letters in
their names. </br>

Single Responsibility Principle </br>
Open/Closed Principle </br>
Liskov Substitution Principle </br>
Interface Segregation Principle </br>
Dependency Inversion Principle </br>

These principles quickly became well known and popular in the software industry. These rules are language independent,
they are the most general principles of code layout. In this document, I dissect and learn how to use it in go to write
good code. </br>

## Single Responsibility Principle <a name="SingleResponsibilityPrinciple"></a>

“Do one thing and do it well” — McIlroy (Unix philosophy) </br>

In short, a class or module should only do one thing. Any reason to maintain its code should stem from only one
purpose. </br>

In practice, to ensure this point, a lot of relevant knowledge is required. You need to understand at a basic level the
domain you are working with, and separate the domain in a nice way. The separation of classes and objects will be based
on your understanding of that domain, which is always associated with the reality of the domain. This is a job that
needs practice and practice continuously, every day. There's no other way. </br>

Please
view: https://github.com/Nghiait123456/GolangAdvance/blob/master/SOLID/SingleResponsibilityPrinciple/main.go. </br>

In this code, BalanceCalculatorNotGood specializes in calculating a partner's balance. For convenience reasons, I added
a function CheckRiskPartner() with the mindset that checking risk partner is quite simple and without much logic, I can
also take advantage of the paymentMethod variable of the struct BalanceCalculatorNotGood. This is a pretty dangerous
mistake and violates the Single responsibility principle. Imagine one fine day, a partner requires dozens of strict risk
control methods. And then, you would develop that code in the struct BalanceCalculatorNotGood. A pot of related and
unrelated codes will grow. When many of your classes or modules suffer from this, the code becomes confusing, difficult
to maintain, and there is always a potential for bugs. </br>

Here, add a struct PartnerRisk that handles and checks partner-related risks. When the ParterRisk code needs to grow, it
will be free with the BalanceCalculator, and it is completely divisible depending on the domain business. </br>

## Open/Closed Principle <a name="OpenClosedPrinciple"></a>

![OpenClosedPrinciple_1.png](img%2FOpenClosedPrinciple_1.png) </br>
![OpenClosedPrinciple_2_1.png](img%2FOpenClosedPrinciple_2_1.png) </br>
![OpenClosedPrinciple_2_2.png](img%2FOpenClosedPrinciple_2_2.png) </br>

“A module should be open for extensions, but closed for modification” — Robert C. Martin

Code expansion and development is always happening. This principle in short, must ensure that extending a module does
not affect dependent and related modules, as well as conversely, adding and modifying dependent modules will not affect
the original module. I make this clear with an example. </br>

Please view bad
example: https://github.com/Nghiait123456/GolangAdvance/blob/master/SOLID/OpenClosedPrinciple/bad/main.go </br>
Please view good example:
https://github.com/Nghiait123456/GolangAdvance/blob/master/SOLID/OpenClosedPrinciple/good/main.go </br>
A classic example of this is the abstract factory patterns. I have a payment factory with lots of payment methods,
including paynow() method. In the bad code, my anonymous factory handles the paynow() by stuffing all the necessary
paynow enable partners:

```
func (p *Payment) PayNow() bool {
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
```

One fine day, I needed to add a payment with my master card, and in addition to adding the MasterCard provider code, I
had to modify the code func (p *Payment) PayNow() bool of the struct Payment. This seriously violates the closed for
modification principle. Obviously adding a MasterCard won't involve struct Payment, but I will always have to fix the
payment code if a new provider arrives. </br>

Another problem that is often hidden, extending the struct Payment will also affect the MasterCard code or any other
code, it is risky because the code is tied together on the same struct Payment. </br>

What's scarier is that this problem will always exist for structs that extend from struct Payment. One fine day, I need
to extend struct Payment , all above problems will appear in new struct. Imagine this problem arises in many places,
it's scary. </br>

Looking at the good code, I solved this problem by abstract factory, and just put the related things together.

```
type Payment struct {
    partnerCode string
    paymentMethod string
    paymentNow PaymentNowInterface
}

type PaymentInterface interface {
    PaymentNow() bool
}

func (p *Payment) PaymentNow() bool {
    return p.paymentNow.PayNow()
}
```

Payment will only care about the PaymentNow() bool, and PaymentNow() will simply call the PayNow() interface. Newly
added providers must implement the PayNow() interface if they want to be valid for Payment. Adding the MasterCard
provider will no longer have to modify the Payment struct type, it ensures the principle of closed for
modification. </br>

```
type PaymentNowInterface interface {
    PayNow() bool
}
```

All the new modules for PayNow simply implements the PaymentNowInterface interface without worrying about the Payment
struct. In golang, this is done easily and simply through interfaces and embedding code.  </br>

Finally, I need a place to aggregate the materials for my factory, this place needs to be independent because the
purpose it was born with is to be independent of Payment Struct from providers. </br>

```
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
```

Good practice of this principle requires a deep understanding of it and an understanding of the module you are
coding. </br>

## Liskov Substitution Principle <a name="LiskovSubstitutionPrinciple"></a>
![LiskovSubstitutionPrinciple.png](img%2FLiskovSubstitutionPrinciple.png) </br>



