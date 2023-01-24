Verify interface compliance at compile time where appropriate. This includes:

1) Exported types that are required to implement specific interfaces as part of their API contract
2) Exported or unexported types that are part of a collection of types implementing the same interface
3) Other cases where violating an interface would break users


We will explain it in code example: