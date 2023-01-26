package __ImportAliasing

//Import aliasing must be used if the package name does not match the last element of the import path
//import (
//	"net/http"
//
//	client "example.com/client-go"
//	trace "example.com/trace/v2"
//)

// this is bad name
//import (
//	"fmt"
//	"os"
//
//
//	nettrace "golang.net/x/trace"
//)

//In all other scenarios, import aliases should be avoided unless there is a direct conflict between imports.
// this is good
//import (
//	"fmt"
//	"os"
//	"runtime/trace"
//
//	nettrace "golang.net/x/trace"
//)
