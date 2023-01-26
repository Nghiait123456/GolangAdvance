package EmbeddingInStructs

/**
Embedded types should be at the top of the field list of a struct, and there must be an empty line separating embedded fields from regular fields.

Embedding should provide tangible benefit, like adding or augmenting functionality in a semantically-appropriate way. It should do this with zero adverse user-facing effects (see also: Avoid Embedding Types in Public Structs).

Exception: Mutexes should not be embedded, even on unexported types. See also: Zero-value Mutexes are Valid.

Embedding should not:

Be purely cosmetic or convenience-oriented.
Make outer types more difficult to construct or use.
Affect outer types' zero values. If the outer type has a useful zero value, it should still have a useful zero value after embedding the inner type.
Expose unrelated functions or fields from the outer type as a side-effect of embedding the inner type.
Expose unexported types.
Affect outer types' copy semantics.
Change the outer type's API or type semantics.
Embed a non-canonical form of the inner type.
Expose implementation details of the outer type.
Allow users to observe or control type internals.
Change the general behavior of inner functions through wrapping in a way that would reasonably surprise users.
Simply put, embed consciously and intentionally. A good litmus test is, "would all of these exported inner methods/fields be added directly to the outer type"; if the answer is "some" or "no", don't embed the inner type - use a field instead.
*/

// this is bad
//type A struct {
//	// Bad: A.Lock() and A.Unlock() are
//	//      now available, provide no
//	//      functional benefit, and allow
//	//      users to control details about
//	//      the internals of A.
//	sync.Mutex
//}
//
//type Book struct {
//	// Bad: pointer changes zero value usefulness
//	io.ReadWriter
//
//	// other fields
//}
//
//// later
//
//var b Book
//b.Read(...)  // panic: nil pointer
//b.String()   // panic: nil pointer
//b.Write(...) // panic: nil pointer
//
//
//type Client struct {
//	sync.Mutex
//	sync.WaitGroup
//	bytes.Buffer
//	url.URL
//}

//this is good
//type countingWriteCloser struct {
//	// Good: Write() is provided at this
//	//       outer layer for a specific
//	//       purpose, and delegates work
//	//       to the inner type's Write().
//	io.WriteCloser
//
//	count int
//}
//
//func (w *countingWriteCloser) Write(bs []byte) (int, error) {
//	w.count += len(bs)
//	return w.WriteCloser.Write(bs)
//}
//
//
//type Book struct {
//	// Good: has useful zero value
//	bytes.Buffer
//
//	// other fields
//}
//
//// later
//
//var b Book
//b.Read(...)  // ok
//b.String()   // ok
//b.Write(...) // ok
//
//
//type Client struct {
//	mtx sync.Mutex
//	wg  sync.WaitGroup
//	buf bytes.Buffer
//	url url.URL
//}
