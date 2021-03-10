// Package version provides the version of this tool. The variable `Version`
// can be populated at compile time using `ldflags`.
package version

var (
	// Version is provided by ldflags at compile time
	Version = "(devel)"
)
