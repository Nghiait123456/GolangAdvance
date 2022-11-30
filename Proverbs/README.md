- [Don't communicate by sharing memory, share memory by communicating](#DontCommunicateBySharingMemory_ShareMemoryByCommunicating)
- [Concurrency is not parallelism](#ConcurrencysIsNotParallelism)
- [The bigger the interface, the weaker the abstraction](#TheBiggerThenIterfaceTheWeakerTheAbstraction)
- [Make the zero value useful](#MakeTheZeroValueUseful)
- [Interface{} says nothing](#interfaceSaysNothing)
- [A little copying is better than a little dependency](#ALittleCopyingIsBetterThanALittleDependency)
- [Syscall must always be guarded with build tags](#SyscallMustAlwaysBeGuardedWithBuildTags)
- [Cgo enables the creation of Go packages that call C code](#CgoEnablesTheCreationOfGoPackagesThatCallCCode)
- [Cgo is not Go](#CgoIsNotGo)
- [With the unsafe package there are no guarantee](#WithTheUnsafePackageThereAreNoGuarantees)
- [Clear is better than clever](#ClearIsBetterThanClever)
- [Reflection is never clear](#ReflectionIsNeverClear)
- [Errors are values](#ErrorsAreValues)
- [Don't just check errors, handle them gracefully](#DontJustCheckErrorsHandleThemGracefully)


## Don't communicate by sharing memory, share memory by communicating <a name="DontCommunicateBySharingMemory_ShareMemoryByCommunicating"></a>

Here are two points to clarify:
Don't communicate by sharing memory: </br>
In golang, almost everything is concurrency. You have a shared memory area, many threads change it together, you must
have protection mechanism for race conditions (most commonly lock). If you use this technique of sharing memory
everywhere, it will be complicated, high cost of switch context. It should not be abused in a concurrency language, only
when there is a specific preference for performance. </br>

Share memory by communicating: </br>
In go routines values move on channels than blocking the memory, sender notifies receiver to receive from that channel
and therefore it share memory by communicating with receiver to get from a channel. By default, a channel is only ready
to send when there is a routine ready for read this channel. </br>
To summarize: </br>
Don't communicate by sharing memory, share memory by communicating. In golang, limit sharing memory when designing
software. Ideally, let the routines run independently of a data stream. If sharing memory is required, use channels, it
is a clever mechanism designed for communicating, handle lock, simple for code. It only works when there is at least 1
read thread and 1 write thread to it, otherwise it will locked. Generally speaking, it is a shared memory tool designed
for communication. --- ""Share memory by communicating""--- </br>

## Concurrency is not parallelism <a name="ConcurrencysIsNotParallelism"></a>

Please view my
doc: https://github.com/Nghiait123456/GolangAdvance/tree/master/ConcurrencyPattern#DistinctiveConcurrencyAndParallelism
? </br>

## Channels orchestrate; mutexes serialize. <a name="ChannelsOrchestrateMutexesSerialize."></a>

Before understanding this part, you should see the documentation I
wrote: https://github.com/Nghiait123456/GolangAdvance/tree/master/ConcurrencyPattern. When you understand and use this
material fluently, you will understand the above statement without analysis. </br>
Channels orchestrate: When developing a go program, there will be concurrency and the ability to communicate between
routines, channels are built for this. Most commonly used concurrency patterns for channel rotation. Channels are not
just a shared data sharing tool, they are a part of project structuring, flow generation, and code structure. Channels
was developed to orchestrate streams, orchestrate layout code, communication streams,... Channels is part of the Go
language parallelization toolkit, shouldering the responsibility of data flow and control flow simultaneously, and at
the same time. it is the organizer of the program structure. </br>
Mutexes serialize: yes, mutex lock, atomic,... born to serialize. If you just want to simply serialize, the program has
no elements of communication and orchestration, use mutex. It has a higher performance than the channel and most
importantly is the right use. Again, if you only need serialize, use mutex. Ex: you only need to increment 1 count every
time there is a corresponding action, at the end of the day, you need to access that count, use mutex,... </br>
Summary: use the sync mechanism for the right purpose to have the most pure and effective software. If you use it wrong,
everything will probably still work, but it can compile things many times over. Imagine using a mutex everywhere to
communicate, using a channel just to send a signal that increments a counter by 1 and doesn't do anything else,...
Everything will still work but it's not happy. There will be countless why questions from your team, which only you can
answer, no one else. </br>

## The bigger the interface, the weaker the abstraction <a name="TheBiggerThenIterfaceTheWeakerTheAbstraction."></a>

In go, interface is commonly used, it is an important part of assembly and communication between parts of go code and it
tends to be small. An interface is large and aggregates many parts, many things, it will be less likely to be extended
by many other components. An interface that is small enough and reflective of its nature will be powerful and true to
the essence of an interface. </br>

```
// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// consider the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
type Reader interface {
Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer interface {
Write(p []byte) (n int, err error)
}
```

Take a look at golang's io platform library. io.Writer and io.Reader are foundational interfaces. All future libraries
related to writing and reading bytes from io just need to implement these two interfaces, they can communicate with each
other. </br>

```
writer := uilive.New()
writer.Start()

for i := 0; i <= 100; i++ {
// writer implements io.Writer
fmt.Fprintf(writer, "Downloading.. (%d/%d) GB\n", i, 100)
}

writer.Stop() // flush and stop rendering
```

Another example error of go, it has only one interface: </br>

```
type error interface {
Error() string
}
```

I can implement many of my own error classes and completely map with the original error class used in go. I just need to
implement the interface Error() string. Since this Interface is small, it will be highly abstract. </br>

## Make the zero value useful <a name="MakeTheZeroValueUseful"></a>

In go, default values should be left to make used. The user will use the packet as soon as a new instace is available
without calling an init() function. To achieve this, the zero value default must be fully utilized. Most of golang's
standard packages adhere to this principle. </br>

View mutex packet:

```
package main

import "sync"

type MyInt struct {
        mu sync.Mutex
        val int
}

func main() {
        var i MyInt

        // i.mu is usable without explicit initialisation.
        i.mu.Lock()
        i.val++
        i.mu.Unlock()
}
```

View io packet: </br>

```
package main

import "bytes"
import "io"
import "os"

func main() {
        var b bytes.Buffer
        b.Write([]byte("Hello world"))
        io.Copy(os.Stdout, &b)
}
```

View silce packet: </br>

```
package main

import "fmt"
import "strings"

func main() {
        // s := make([]string, 0)
        // s := []string{}
        var s []string

        s = append(s, "Hello")
        s = append(s, "world")
        fmt.Println(strings.Join(s, " "))
}
```

Default zero values are fully used, 3rd parties only need to create an instance and use it. You should design code
that guarantees this feature. However, in some cases, calling an extra Init function is not bad or unusual. A mockle to
initialize and set up the config of a web service, it is almost indispensable to have an Init() function with all the
settings. Use it flexibly and only add the Init() function when you can't do without it. </br>

## Interface{} says nothing <a name="interfaceSaysNothing"></a>

Yes, interface{} says nothing. It can be nothing or it can be anything. So, is it good or bad and when to use it. Use
interface{} only when your input is really an interface{}, there are multiple types for input and you use those types in
your work. Use interface{} only when you need it for work. If your function doesn't really need it to run, don't use it,
overusing it makes the code abstract and a lot of questions: why, what type, what context for type,... Anyone Code
readers (including writers) will have to try to answer abstract questions, and in the end the answer is that it is not
as abstract as interface{}.
Examples of using interface{}:

```
// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...any) (n int, err error) {
return Fprintln(os.Stdout, a...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintln(a ...any) string {
p := newPrinter()
p.doPrintln(a)
s := string(p.buf)
p.free()
return
}

```

When you print something, it can have any type, if you maintain all functions with all types, it's a waste because the
end goal is to print it. Use any ( interface{}) and handle all the type logic inside. </br>

## A little copying is better than a little dependency <a name="ALittleCopyingIsBetterThanALittleDependency"></a>

One Proverbs is the most debated, because at first glance it seems to contradict another very famous proverb: "Don't
repeat yourself". So are they really contradictory and why does Rob Pike make this point? </br>

First, they are not contradictory, they support each other. In most cases, you should "Don't repeat yourself". Let's DI
the small pieces and make the big puzzle, the big puzzle is a synthesis, the machine completes and runs elegantly. But,
should you apply it 100% and always DI everything when you need it. The answer is no, there is a lot of pain when you DI
100% of what you need. With features that can be replaced with a single line of code and low code reusability and
maintainability, don't import an entire library to run that feature. If really the need for customine and maintain the
code is not high, copy the code and use it. It avoids the bloat of DI code, avoids errors from a large library while the
function you use is very small and it does not fail. Again, the boundaries of this are very thin, if you abuse either of
these two quotes, that is not good. </br>

A simple example:

```
package main

import (
"fmt"
"os"
)

func main() {
f, _ := os.Open("/dev/urandom")
b := make([]byte, 16)
f.Read(b)
f.Close()
uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10 ], b[10:])
fmt.Println(uuid)
}
```

If here you simply need a UUID, copy the code and don't use a library. However, if you don't understand the cases and
exceptions about the code you copy, please use the library, don't copy the code. After all, they're all code, if it's
stable and you know enough about it, use it, if not, use the library. Stability should be a top priority factor in this
case. </br>

## Syscall must always be guarded with build tags <a name="SyscallMustAlwaysBeGuardedWithBuildTags"></a>

Different systems (*UNIX, Windows,...) calling the same function (implementations are not the same) may need to be built
on different systems to get the results you want. Simply put, this is because system calls are specific to each
OS. </br>

More intuitively, each system call function list must be constructed separately for the os families: Unix, windows,
solaris,... Go must build separate files for each OS. You need to see the corresponding configurations and settings on
different environments, there is no way for a system function to be sure to run on every os. </br>

## Cgo enables the creation of Go packages that call C code <a name="CgoEnablesTheCreationOfGoPackagesThatCallCCode"></a>

Similar to the problem above, when calling c, It's very non-portable. It needs to be built for specific architectures
and operating systems. </br>

## Cgo is not Go <a name="CgoIsNotGo"></a>

Cgo enables the creation of Go packages that call C code. </br>

A lot of people in the early days would write about how a favorite feature of Go was how easily it connected to C, but
lots of times it shouldn't be necessary, it takes you out of the Go universe and you lose the benefits of Go if you are
coding in C. </br>

Link detail: https://dave.cheney.net/2016/01/18/cgo-is-not-go </br>

## With the unsafe package there are no guarantees <a name="WithTheUnsafePackageThereAreNoGuarantees"></a>

This is obvious, the package is not something divine, it is the code. If you use an unsafe package, maybe some features
are wrong or unstable, or one fine day your project will be down. Package are like weapons, choosing the right one and
using it right depends on the you. </br>

## Clear is better than clever <a name="ClearIsBetterThanClever"></a>

There are languages that value intelligence, it is often a combination of many utilities in one function with many
features. With foundational things like languages, the smarter is the complexity, and it matters if that code underlies
everything else. Go, like most programming languages, values clarity over clever. </br>

## Reflection is never clear <a name="ReflectionIsNeverClear"></a>

Common Stackoverflow question of people wanting to use reflect and complaining that it doesn’t work. It doesn’t work,
because it is not for you. Very, very few people should be playing with this. Powerful, yet very difficult to use. We
should encourage beginners to step away from using reflection and use the language proper. </br>

In most jobs, you don't need Reflection. Don't worry too much about it, only use it when absolutely necessary. </br>

## Errors are values <a name="ErrorsAreValues"></a>

Too often people write “err != nil” — they think about substituting try/catch. In the golang world, errors are just
values, you will either treat it as a variable, send errors back, handle the error or ignore it, it's entirely up to
you. Program with errros and do anything with it, it's just value. High level languages make use of try catch and
exceptions, but it's a convenient abstraction, you can't easily program with it like a variable. </br>

## Don't just check errors, handle them gracefully <a name="DontJustCheckErrorsHandleThemGracefully"></a>

People are too quick to just return an error up the tree, instead of designing how it should work. A big part of writing
good Go code is getting the error handling right up front. Of any program really, but its easier to program with errors
as just values, and easier to do it gracefully. </br>

You handle errors logically and gracefully, instead of just return and break, your program will be higher quality and
cleaner. </br>

A quality post on how to handle errors in
go: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully </br>


todo: doc link ve xu ly loi, hieu no roi lam tiep