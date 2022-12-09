- [Preview](#Preview)
- [Encapsulation](#Encapsulation)
- [Abstraction](#Abstraction)

## Preview <a name="Preview"></a>

Golang is born as an extremely powerful and preeminent concurrency language in today's world, bettering a long-standing
concurrency approach, a big idea about using server resources. As you know, golang itself is not pure OOP, but I can
develop fully OOP guarantee code in go. A large, complex, feature-rich project is completely deployable and easy to
maintain in golang with the powerful tools it provides. </br>

## Encapsulation <a name="Encapsulation"></a>

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

## Abstraction <a name="Abstraction"></a>

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