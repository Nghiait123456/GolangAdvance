1) Use defer to clean up resources such as files and locks.
2) Defer has an extremely small overhead and should be avoided only if you can prove that your function execution time is in the order of nanoseconds.
The readability win of using defers is worth the miniscule cost of using them. This is especially true for larger methods that have more than simple memory accesses, where the other computations are more significant than the defer.

We will explain it in code example: