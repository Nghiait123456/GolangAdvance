todo:
https://s8sg.medium.com/solid-principle-in-go-e1a624290346
https://dave.cheney.net/2016/08/20/solid-go-design
https://web.archive.org/web/20150906155800/http://www.objectmentor.com/resources/articles/Principles_and_Patterns.pdf

- [Preview](#Preview)
- [Single Responsibility Principle](#SingleResponsibilityPrinciple)
- [Open/Closed Principle](#OpenClosedPrinciple)
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
