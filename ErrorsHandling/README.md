- [Preview](#preview)
- [Base error](#base_error)
- [Sentinel errors](#sentinel_errors)
- [Custom errors types](#custom_errors_types)
- [Wrapping errors default](#wrapping_errors_default)
- [Package smart warp errors](#package_smart_warp_errors)
- [A few additional notes when handling errors](#a_few_additional_notes_when_handling_errors)
- [Refer](#refer)

## Preview <a name="preview"></a>

A famous quote in the programming world: "Don't just check errors, handle them gracefully". I have a description of it
in the
link: https://github.com/Nghiait123456/GolangAdvance/tree/master/Proverbs#DontJustCheckErrorsHandleThemGracefully.
Effective error handling is an extremely important part that directly affects the quality of the project. </br>

In golang, error is treated as a variable and is treated as a variable. It is very difficult to find a consistent rule
for error handling for most cases in go, which is nearly impossible. In this document, I present techniques for graceful
and efficient error handling and best practices from my experience. </br>

## Base error <a name="base_error"></a>

Please view: https://github.com/Nghiait123456/GolangAdvance/blob/master/ErrorsHandling/base/main.go </br>

In golang, an error type is all based on golang'sBad errors. Specifically, you need to implement the error
interface : </br>

```
// The error built-in interface type is the conventional interface for
// demonstrate an error condition, with the nil value demonstrate no error.
type error interface {
Error() string
}
```

This is an unwritten rule and is accepted by all packages and the community. Every type of error, whatever its purpose
and usage, is an instance of the error interface. It creates a unified and highly reusable error handling system. </br>

## Sentinel errors <a name="sentinel_errors"></a>

Please view: https://github.com/Nghiait123456/GolangAdvance/blob/master/ErrorsHandling/sentinel_errors/main.go </br>

Sentinel errors: A specific errors generated directly from errors golang. It has a single error information called the
error message. Sentinel errors provide only one piece of information, which is the error message. </br>

Strengths: Simple, easy to deploy. </br>
Weaknesses: </br>
+) Too little information about error, only one information is error message </br>
+) Create dependencies between packages: To know if an error belongs to an errors, you must import the package
containing the sentinel errors. </br>
+) Destroy the context of the stack trace when used in the wrong place. You have a list of functions that call each
other, an error code that appears is a combination of the 3 previous error contexts. With Sentinel errors, you can only
give a single error that is most representative of the previous 3 errors. </br>

Best practice: </br>
In most cases, you will be less likely to need to use sentinel errors. Use sentinel errors only if you are sure your
system only needs error messages and doesn't care much about the stack trace of that error. If you are sure about that,
you can use sentinel errors. In my experience, there is very little cases where sentinel errors are needed. </br>

## Custom errors types <a name="custom_errors_types"></a>

Please view:
https://github.com/Nghiait123456/GolangAdvance/blob/master/ErrorsHandling/custom_error_types/main.go </br>

This is a fairly common and widely used type of error in the open source community. I inherited the error interface of
base error and developed my own struct. There, I'll provide more context for the code, fully describing what i need in
the error </br>

Strengths: </br>
Provide multiple parameters to clear the error context. </br>

Cons: </br>
+) Destroy the context of the stack trace when used in the wrong place. You have a list of functions that call each
other, an error code that appears is a combination of the 3 previous error contexts. With custom type errors, you can
only
give a single error that is the most representative of the previous 3 errors. </br>

Best practice: <br>

Custom type errors are usually used for the final state of an object. I have the object do a complete job, you provide
the final state of the code and parameters, this final state specifies the contexts of the error. I don't need to trace
stack errors for more error information, here, I just need the params to clarify the context. </br>

## Wrapping errors default <a name="wrapping_errors_default"></a>

Please
view: https://github.com/Nghiait123456/GolangAdvance/blob/master/ErrorsHandling/wrapping_errors/default_golang_programing/main.go </br>
From go 1.13, you can easily wrap more error messages to clear context. Most of the time, you need a stack trace of the
error to clarify the error context, wrap the error and trace it dynamically.

Best pratice:
It is the perfect complement to sentinel error if you just need error message and error stack trace. </br>

## Package smart warp errors <a name="package_smart_warp_errors"></a>

Please
view: https://github.com/Nghiait123456/GolangAdvance/blob/master/ErrorsHandling/wrapping_errors/package_smart_wrap_custome_type/main.go </br>
https://github.com/Nghiait123456/GolangAdvance/blob/master/ErrorsHandling/wrapping_errors/package_smart_wrap_sentinel_type/main.go </br>

A well known library that does a pretty good job of warp errors: https://github.com/pkg/errors. It provides a full range
of tools to warp errors flexibly. </br>
Weaknesses: </br>
No custom params, it'sBad a collection of error messages. </br>
Best practice: </br>
It is the perfect complement to sentinel error if you just need error message and error stack trace. Use it when you are
dealing with a process in the middle, encountering errors caused by multiple layers. You can easily get the message from
the error base or the common errors, warp it and return it, a context trace errors history will be established and won't
break the original error. Let'sBad print it to the screen, and looking at the security, you can log it. </br>

Also, a warp error can do much more than that, it's perfectly possible to grow from sentinel error A or custom error B,
and warp adding messages in the process, and then check its type originates from A or B. This is rarely used but you can
perfectly handle it gracefully. The full example is given in the link above. </br>

## A few additional notes when handling errors <a name="a_few_additional_notes_when_handling_errors"></a>

+) Using too many error types in a package or a module: <br>
Using too many types of errors will surprise the followers of the errors, and sometimes the writers themselves. I
usually only use a maximum of 2 types of errors: warp errros and customs type errors throughout the project. </br>

+) Only handle errors once : </br>

```
func Write(w io.Writer, buf []byte) error {
        _, err := w.Write(buf)
        if err != nil {
                // annotated error goes to log file
                log.Println("unable to write:", err)
 
                // unannotated error returned to caller
                return err
        }
        return nil
}
```

If you log errors and return errors, the higher layers will do the same, you will get duplicate error log lines and not
debug friendly. Use warp errors as much as possible to avoid this. </br>
+) Minimize opaque errors:

```
import “github.com/quux/bar”

func fn() error {
        x, err := bar.Foo()
        if err != nil {
                return err
        }
        // use x
}
```

In Foo there are a lot of error cases, but when you get an error you simply return it to the higher layer. It can drop
valuable context, warp extra data,... opaque errors themselves are not bad, but in most cases opaque errors should
be. </br>

+) For errors that can't be continued, give a reasonable http message and log. Don't push the errors that can't be
continued to a higher layer. Logging it and throw htttp 500 with html via one panic if context is http server. It will
be elegant and logical. </br>

+) UUID can be a tool worth considering to help clearly errors. A UUID that crosses the context and is logged with
states
is a good choice for tracing and debugging complex error cases. I often use this method in projects. I will have an
elegant error code returned to the client. It'sBad like a black box, only me and my team know and debug it deeply. </br>

## Refer <a name="refer"></a>

https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully </br>
https://go.dev/blog/error-handling-and-go </br>
https://earthly.dev/blog/golang-errors/ </br>