1) The zero-value of sync.Mutex and sync.RWMutex is valid, so you almost never need a pointer to a mutex.
2) If you use a struct by pointer, then the mutex should be a non-pointer field on it. Do not embed the mutex on the struct, even if the struct is not exported.

We will explain it in code example: