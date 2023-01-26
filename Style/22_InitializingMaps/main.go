package InitializingMaps

/**
Prefer make(..) for empty maps, and maps populated programmatically. This makes map initialization visually distinct from declaration, and it makes it easy to add size hints later if available.
*/

// this is bad
//var (
//	// m1 is safe to read and write;
//	// m2 will panic on writes.
//	m1 = map[T1]T2{}
//	m2 map[T1]T2
//)

// this is good
//var (
//	// m1 is safe to read and write;
//	// m2 will panic on writes.
//	m1 = make(map[T1]T2)
//	m2 map[T1]T2
//)

/**
Where possible, provide capacity hints when initializing maps with make(). See Specifying Map Capacity Hints for more information.

On the other hand, if the map holds a fixed list of elements, use map literals to initialize the map.
*/

// this is bad
//m := make(map[T1]T2, 3)
//m[k1] = v1
//m[k2] = v2
//m[k3] = v3

// this is good
//m := map[T1]T2{
//k1: v1,
//k2: v2,
//k3: v3,
//}

/**
The basic rule of thumb is to use map literals when adding a fixed set of elements at initialization time, otherwise use make (and specify a size hint if available).
*/
