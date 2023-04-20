- [Preview](#preview)
- [Encapsulation](#encapsulation)
- [Abstraction](#abstraction)
- [Inheritance](#inheritance)
- [Polymorphism](#polymorphism)

## Preview <a name="preview"></a>

Golang is born as an extremely powerful and preeminent concurrency language in today'sBad world, bettering a
long-standing
concurrency approach, a big idea about using server resources. As you know, golang itself is not pure OOP, but I can
develop fully OOP guarantee code in go. A large, complex, feature-rich project is completely deployable and easy to
maintain in golang with the powerful tools it provides. </br>

## Encapsulation <a name="encapsulation"></a>

Encapsulation is an operation that ensures data is used by the correct object it allows. In OOP language, public,
protect, private,.. scopes are very familiar. With golang, this is approximately at the package level. </br>

In golang, there are two states of data, exported and unexported. Exported, you can access the package data from another
package, with unexported you can't do that, you can only access the package from within the package itself. </br>

All exported states will have an uppercase initial, and unexported states will have a lowercase first letter. </br>

```
type Customer struct {
	id   int // is unexported
	name string // is unexported
}


// is exported
func (c *Customer) GetID() int {
	return c.id
}

// is unexported
func (c *Customer) getName() string {
	return c.name
}
```

## Abstraction <a name="abstraction"></a>

This is a common property everywhere, every object, machine,... around you. A certain tool will have 2 main parts, the
user interface (user manual), and the structure (how to, how it works). From a user perspective, I usually only care
about the usage part. For example: I ride a motorbike, i care how to steer, how to enter and decelerate, how to shift
gears, brake... I will not care about how cylinders, transmissions... work. Its internals are always working, but I
don't need to care about that to operate it. </br>

In software, the same is true. A software module should be designed with a clear separation between the user interface (
third parties using it will be interested), and its interior, how it works, operation... A software will be a collection
of many such parts joined together. </br>

In Go, Abstraction can be powerfully and simply implemented through interfaces. Define the interfaces of a module, users
will easily use it through flexible interface design. </br>

```
type Customer struct {
	id   int // is unexported
	name string // is unexported
}

type CustomerInterface interface {
    FindNewFriend() // this is Abstraction, other sturct only call and use it 
    ShowLocation()  // this is Abstraction, other sturct only call and use it 
}

// is exported
func (c *Customer) FindNewFriend()  {
	//todo process....
}

// is unexported
func (c *Customer) ShowLocation()  {
	//todo process....
}
```

## Inheritance <a name="inheritance"></a>

Why is there no type inheritance? </br>

Object-oriented programming, at least in the best-known languages, involves too much discussion of the relationships
between types, relationships that often could be derived automatically. Go takes a different approach. </br>

Rather than requiring the programmer to declare ahead of time that two types are related, in Go a type automatically
satisfies any interface that specifies a subset of its methods. Besides reducing the bookkeeping, this approach has real
advantages. Types can satisfy many interfaces at once, without the complexities of traditional multiple inheritance.
Interfaces can be very lightweight—an interface with one or even zero methods can express a useful concept. Interfaces
can be added after the fact if a new idea comes along or for testing—without annotating the original types. Because
there are no explicit relationships between types and interfaces, there is no type hierarchy to manage or discuss. </br>

It's possible to use these ideas to construct something analogous to type-safe Unix pipes. For instance, see how
fmt.Fprintf enables formatted printing to any output, not just a file, or how the bufio package can be completely
separate from file I/O, or how the image packages generate compressed image files. All these ideas stem from a single
interface (io.Writer) representing a single method (Write). And that's only scratching the surface. Go's interfaces have
a profound influence on how programs are structured. </br>

It takes some getting used to but this implicit style of type dependency is one of the most productive things about Go. </br>

This link: https://go.dev/doc/faq#inheritance

## Polymorphism <a name="polymorphism"></a>

In OOP, there is often a class hierarchy that includes, for example, a superclass A, and a group of subclasses that
directly inherit the superclass: B, C, D. Polymorphism in OOP allows the same name, but will have different behavior in
different subclasses (B, C, D,... ), or different subclasses may also have different behavior from parent class A. More
broadly, they are completely It is possible to replace the superclass A, but in that case, you won't know its behavior,
until you actually define the binding object as B, C, or D. This example is an example. Classic example and extremely
popular in OOP. </br>

In golang, polymorphism is also powerful and flexible, it relies on interfaces. Small and powerful designed interface
combined with embedded struct also creates strong OOP polymorphism in go. </br>

```
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
```