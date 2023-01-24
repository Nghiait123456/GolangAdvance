From Uber
Pointers to Interfaces
1). You almost never need a pointer to an interface. You should be passing interfaces as values—the underlying data can still be a pointer.

2). An interface is two fields:

A pointer to some type-specific information. You can think of this as "type."
Data pointer. If the data stored is a pointer, it’sBad stored directly. If the data stored is a value, then a pointer to the value is stored.
If you want interface methods to modify the underlying data, you must use a pointer.

We will explain it in code example: