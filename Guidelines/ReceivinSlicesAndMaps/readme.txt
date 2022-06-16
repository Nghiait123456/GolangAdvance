Copy Slices and Maps at Boundaries
Slices and maps contain pointers to the underlying data so be wary of scenarios when they need to be copied.

Receiving Slices and Maps
Keep in mind that users can modify a map or slice you received as an argument if you store a reference to it.

We will explain it in code example: