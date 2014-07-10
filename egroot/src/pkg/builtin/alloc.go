package builtin

import "unsafe"


// Alloc is used for dynamic memory allocation. Its deffinition will be changed
// in the future to accept pointer to some TypeInfo struct instead of size and
// align.
var Alloc func(n int, size, align uintptr) unsafe.Pointer