- [Don't communicate by sharing memory, share memory by communicating](#DontCommunicateBySharingMemory_ShareMemoryByCommunicating)
- [Concurrency is not parallelism](#ConcurrencysIsNotParallelism)
- [The bigger the interface, the weaker the abstraction](#TheBiggerThenIterfaceTheWeakerTheAbstraction)


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

link tham khao
https://gregosuri.com/2015/12/04/go-proverbs-illustrated/