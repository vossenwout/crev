package logging

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// DisplayUsageAndExit displays the usage information and exits the program.
func DisplayUsageAndExit() {
	log.SetFlags(0)
	printUsage()
	os.Exit(1)
}

// printUsage prints the detailed usage information.
func printUsage() {
	fmt.Println("Usage: ai-code-review [options] <root directory>")
	fmt.Println()
	fmt.Println("Options:")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("  -%s=%s\n    \t%s\n", f.Name, f.DefValue, f.Usage)
	})
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ai-code-review .")
	fmt.Println("  ai-code-review cmd/")
	fmt.Println("  ai-code-review cmd/ -concurrency=100 -ignore=.,readme,tests")
	fmt.Println()
	fmt.Println("Notes:")
	fmt.Println("  - The root directory should be the first non-flag argument.")
	fmt.Println("  - Patterns are used to exclude specific files or directories from processing.")
}
