package __ImportGroupOrdering

// this is bad
//import (
//	"fmt"
//	"os"
//	"go.uber.org/atomic"
//	"golang.org/x/sync/errgroup"
//)

// this is good
//import (
//	"fmt"
//	"os"
//
//	"go.uber.org/atomic"
//	"golang.org/x/sync/errgroup"
//)

/**
There should be two import groups:

Standard library
Everything else
This is the grouping applied by goimports by default.
*/
