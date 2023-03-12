- [Preview](#preview)
- [Garbage Collection Advance](#garbage-collection-advance)
- [GMOP](#gmop)

## Preview <a name="preview"></a>

Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines. It
also includes the low-level type information used by the reflect package; see reflect's documentation for the
programmable interface to the run-time type system. The most general description of the go runtime design but it is
rather difficult to access and not written for learning, it is for the writers of this
language: https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw /edit# </br>

## Garbage Collection Advance <a name="garbage-collection-advance"></a>

Garbage collection is an important component in the go runtime, it automatically cleans the heap on the go runtime. I
describe it from basic to advanced in this link: https://www.youtube.com/watch?v=5dw6_HrU_zU&t=43s. This is a
conversation vs Who has a detailed scenario, I have tried to describe from basic to advanced so that it can be easily
reached by newcomers. </br>

## GMOP <a name="gmop"></a>

The most important and fundamental component of the go runtime is the GMOP. In this link, I will describe from basic to
advanced, so that people who don't know go runtime can understand: https://www.youtube.com/watch?v=Lu7doNezsM8 </br>