please view: https://github.com/Nghiait123456/GolangAdvance/tree/master/Proverbs#dont_panic


Code running in production must avoid panics. Panics are a major source of cascading failures. If an error occurs, the function must return an error and allow the caller to decide how to handle it. Panic/recover is not an error handling strategy. A program must panic only when something irrecoverable happens such as a nil dereference. An exception to this is program initialization: bad things at program startup that should abort the program may cause panic.
