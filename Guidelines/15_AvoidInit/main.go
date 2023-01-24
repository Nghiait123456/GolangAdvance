package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"path"
)

/**
Avoid init() where possible. When init() is unavoidable or desirable, code should attempt to:

Be completely deterministic, regardless of program environment or invocation.
Avoid depending on the ordering or side-effects of other init() functions. While init() ordering is well-known, code can change, and thus relationships between init() functions can make code brittle and error-prone.
Avoid accessing or manipulating global or environment state, such as machine information, environment variables, working directory, program arguments/inputs, etc.
Avoid I/O, including both filesystem, network, and system calls.
Code that cannot satisfy these requirements likely belongs as a helper to be called as part of main() (or elsewhere in a program's lifecycle), or be written as part of main() itself. In particular, libraries that are intended to be used by other programs should take special care to be completely deterministic and not perform "init magic".


Considering the above, some situations in which init() may be preferable or necessary might include:

Complex expressions that cannot be represented as single assignments.
Pluggable hooks, such as database/sql dialects, encoding type registries, etc.
Optimizations to Google Cloud Functions and other forms of deterministic precomputation.

*/

//this is bad example
type FooBad struct {
	// ...
}

var _defaultFoo FooBad

func init() {
	_defaultFoo = FooBad{
		// ...
	}
}

type ConfigBad struct {
	// ...
}

var _config ConfigBad

func init() {
	// Bad: based on current directory
	cwd, _ := os.Getwd()

	// Bad: I/O
	raw, _ := os.ReadFile(
		path.Join(cwd, "config", "config.yaml"),
	)

	yaml.Unmarshal(raw, &_config)
}

// this is good example
type Foo struct {
}

//var _defaultFooGood = Foo{
//	// ...
//}

// or, better, for testability:

var _defaultFooGood = defaultFoo()

func defaultFoo() Foo {
	return Foo{
		// ...
	}
}

type ConfigGood struct {
	// ...
}

func loadConfig() ConfigGood {
	cwd, _ := os.Getwd()
	// handle err

	raw, _ := os.ReadFile(
		path.Join(cwd, "config", "config.yaml"),
	)
	// handle err

	var config ConfigGood
	yaml.Unmarshal(raw, &config)

	return config
}
