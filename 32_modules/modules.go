/*
=============================================================================
                           📦 GO MODULES TUTORIAL
=============================================================================

📚 CORE CONCEPT:
Go modules are the official dependency management system. They define
a collection of packages with versioning, dependencies, and metadata.

🔑 KEY FEATURES:
• Dependency management and versioning
• Reproducible builds
• Module proxy and checksum database
• Semantic versioning support

💡 REAL-WORLD ANALOGY:
Module = Recipe Book
- go.mod = Table of contents with ingredient list
- go.sum = Ingredient verification checksums
- Dependencies = Required ingredients from other cookbooks
- Versions = Edition numbers of cookbooks

🎯 WHY USE MODULES?
• Reliable dependency management
• Version control and compatibility
• Reproducible builds across environments
• Simplified project setup

=============================================================================
*/

package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("📦 GO MODULES TUTORIAL")
	fmt.Println("======================")

	// 🎯 DEMO 1: Module Information
	fmt.Println("\n🎯 DEMO 1: Current Module Info")
	fmt.Println("==============================")

	// Get current working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("❌ Error getting working directory: %v\n", err)
	} else {
		fmt.Printf("📁 Current directory: %s\n", pwd)
	}

	// Go environment information
	fmt.Printf("🔧 Go version: %s\n", runtime.Version())
	fmt.Printf("🔧 Go root: %s\n", runtime.GOROOT())
	
	// Build context
	ctx := build.Default
	fmt.Printf("🔧 GOPATH: %s\n", ctx.GOPATH)
	fmt.Printf("🔧 GOOS: %s\n", ctx.GOOS)
	fmt.Printf("🔧 GOARCH: %s\n", ctx.GOARCH)

	// 🎯 DEMO 2: Module Commands Overview
	fmt.Println("\n🎯 DEMO 2: Module Commands")
	fmt.Println("==========================")

	fmt.Println("📝 Essential Go Module Commands:")
	fmt.Println()

	commands := []struct {
		command     string
		description string
	}{
		{"go mod init <module-name>", "Initialize a new module"},
		{"go mod tidy", "Add missing and remove unused modules"},
		{"go mod download", "Download modules to local cache"},
		{"go mod verify", "Verify dependencies have expected content"},
		{"go mod why <package>", "Explain why packages are needed"},
		{"go mod graph", "Print module requirement graph"},
		{"go mod edit", "Edit go.mod from tools or scripts"},
		{"go get <package>", "Add dependency and update go.mod"},
		{"go get -u", "Update dependencies to latest versions"},
		{"go list -m all", "List all modules in build"},
	}

	for _, cmd := range commands {
		fmt.Printf("  %-30s - %s\n", cmd.command, cmd.description)
	}

	// 🎯 DEMO 3: Module File Structure
	fmt.Println("\n🎯 DEMO 3: Module File Structure")
	fmt.Println("================================")

	fmt.Println("📁 Typical Go module structure:")
	fmt.Println(`
myproject/
├── go.mod          # Module definition and dependencies
├── go.sum          # Dependency checksums for verification
├── main.go         # Main application entry point
├── README.md       # Project documentation
├── internal/       # Private packages (not importable)
│   └── config/
│       └── config.go
├── pkg/           # Public packages (importable by others)
│   └── utils/
│       └── utils.go
├── cmd/           # Application entry points
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
└── vendor/        # Vendored dependencies (optional)
    └── ...`)

	// 🎯 DEMO 4: go.mod File Example
	fmt.Println("\n🎯 DEMO 4: go.mod File Structure")
	fmt.Println("===============================")

	fmt.Println("📄 Example go.mod file:")
	fmt.Println(`
module github.com/username/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/stretchr/testify v1.8.4
    golang.org/x/crypto v0.12.0
)

require (
    // Indirect dependencies (automatically managed)
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    // ... more indirect dependencies
)

replace (
    // Replace directive for local development
    github.com/username/mylib => ../mylib
)

exclude (
    // Exclude specific versions
    github.com/problematic/package v1.2.3
)`)

	// 🎯 DEMO 5: Semantic Versioning
	fmt.Println("\n🎯 DEMO 5: Semantic Versioning")
	fmt.Println("==============================")

	fmt.Println("📋 Semantic Versioning (SemVer) format: MAJOR.MINOR.PATCH")
	fmt.Println()

	versions := []struct {
		version     string
		description string
	}{
		{"v1.2.3", "Stable release: Major=1, Minor=2, Patch=3"},
		{"v0.1.0", "Pre-release: Major=0 indicates unstable API"},
		{"v2.0.0", "Breaking change: Major version increment"},
		{"v1.3.0", "New feature: Minor version increment"},
		{"v1.2.4", "Bug fix: Patch version increment"},
		{"v1.2.3-beta.1", "Pre-release version with identifier"},
		{"v1.2.3+build.1", "Build metadata (ignored by Go modules)"},
	}

	for _, v := range versions {
		fmt.Printf("  %-15s - %s\n", v.version, v.description)
	}

	// 🎯 DEMO 6: Version Selection
	fmt.Println("\n🎯 DEMO 6: Version Selection")
	fmt.Println("============================")

	fmt.Println("📋 Go module version selection rules:")
	fmt.Println()

	rules := []struct {
		pattern     string
		description string
	}{
		{"go get package@latest", "Get latest stable version"},
		{"go get package@v1.2.3", "Get specific version"},
		{"go get package@v1", "Get latest v1.x.x version"},
		{"go get package@master", "Get latest commit from master branch"},
		{"go get package@commit", "Get specific commit hash"},
		{"go get package@none", "Remove dependency"},
		{"go get -u package", "Update to latest compatible version"},
		{"go get -u=patch package", "Update to latest patch version only"},
	}

	for _, rule := range rules {
		fmt.Printf("  %-25s - %s\n", rule.pattern, rule.description)
	}

	// 🎯 DEMO 7: Module Proxy
	fmt.Println("\n🎯 DEMO 7: Module Proxy")
	fmt.Println("=======================")

	fmt.Println("🌐 Go Module Proxy benefits:")
	fmt.Println("  • Faster downloads through caching")
	fmt.Println("  • Availability even if source is down")
	fmt.Println("  • Immutable versions for reproducibility")
	fmt.Println("  • Security through checksum verification")
	fmt.Println()

	fmt.Println("🔧 Environment variables:")
	fmt.Println("  GOPROXY=https://proxy.golang.org,direct")
	fmt.Println("  GOSUMDB=sum.golang.org")
	fmt.Println("  GOPRIVATE=github.com/mycompany/*")

	// 🎯 DEMO 8: Best Practices
	fmt.Println("\n🎯 DEMO 8: Module Best Practices")
	fmt.Println("================================")

	fmt.Println("✅ Module best practices:")
	fmt.Println("  • Use semantic versioning consistently")
	fmt.Println("  • Run 'go mod tidy' regularly")
	fmt.Println("  • Commit go.sum file to version control")
	fmt.Println("  • Use replace directive for local development")
	fmt.Println("  • Keep dependencies up to date")
	fmt.Println("  • Use internal/ for private packages")
	fmt.Println("  • Document breaking changes clearly")
	fmt.Println("  • Test with different Go versions")

	// 🎯 DEMO 9: Troubleshooting
	fmt.Println("\n🎯 DEMO 9: Common Issues & Solutions")
	fmt.Println("===================================")

	issues := []struct {
		problem  string
		solution string
	}{
		{
			"Module not found",
			"Check module path, run 'go mod tidy'",
		},
		{
			"Version conflicts",
			"Use 'go mod why' to understand dependencies",
		},
		{
			"Checksum mismatch",
			"Run 'go clean -modcache' and retry",
		},
		{
			"Private repo access",
			"Set GOPRIVATE environment variable",
		},
		{
			"Outdated dependencies",
			"Run 'go get -u' to update",
		},
	}

	for _, issue := range issues {
		fmt.Printf("  Problem: %s\n", issue.problem)
		fmt.Printf("  Solution: %s\n\n", issue.solution)
	}

	// Check if we're in a module
	fmt.Println("🔍 Module Detection:")
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("  ✅ This directory contains a go.mod file")
	} else {
		fmt.Println("  ❌ No go.mod file found in current directory")
		fmt.Println("  💡 Run 'go mod init <module-name>' to create one")
	}

	fmt.Println("\n✨ All module demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📦 MODULE STRUCTURE:
┌─────────────────────────────────────────────────────────────────────────┐
│ // go.mod file structure                                                │
│ module github.com/username/project                                      │
│                                                                         │
│ go 1.21                                                                 │
│                                                                         │
│ require (                                                               │
│     github.com/gin-gonic/gin v1.9.1                                     │
│     github.com/stretchr/testify v1.8.4                                  │
│ )                                                                       │
│                                                                         │
│ require (                                                               │
│     // Indirect dependencies                                            │
│     github.com/bytedance/sonic v1.9.1 // indirect                       │
│ )                                                                       │
│                                                                         │
│ replace github.com/old/package => github.com/new/package v1.0.0         │
│ exclude github.com/bad/package v1.2.3                                   │
└─────────────────────────────────────────────────────────────────────────┘

🔧 MODULE COMMANDS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Initialize new module                                                │
│ go mod init github.com/username/project                                 │
│                                                                         │
│ // Add dependency                                                       │
│ go get github.com/gin-gonic/gin                                         │
│                                                                         │
│ // Update dependencies                                                  │
│ go get -u                              // Update all                    │
│ go get -u=patch                        // Update patches only           │
│ go get package@latest                  // Update specific package       │
│                                                                         │
│ // Clean up dependencies                                                │
│ go mod tidy                            // Add missing, remove unused     │
│                                                                         │
│ // Verify dependencies                                                  │
│ go mod verify                          // Check checksums               │
│ go mod download                        // Download to cache             │
│                                                                         │
│ // Inspect dependencies                                                 │
│ go list -m all                         // List all modules              │
│ go mod why package                     // Why is package needed         │
│ go mod graph                           // Dependency graph              │
└─────────────────────────────────────────────────────────────────────────┘

📋 VERSION PATTERNS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│   Pattern       │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ v1.2.3          │ Exact version                                           │
│ v1.2            │ Latest patch in v1.2.x                                  │
│ v1              │ Latest minor in v1.x.x                                  │
│ latest          │ Latest stable version                                   │
│ master          │ Latest commit on master branch                          │
│ commit-hash     │ Specific commit                                         │
│ none            │ Remove dependency                                       │
└─────────────────┴─────────────────────────────────────────────────────────┘

🌐 MODULE PROXY:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Environment variables                                                │
│ GOPROXY=https://proxy.golang.org,direct                                 │
│ GOSUMDB=sum.golang.org                                                  │
│ GOPRIVATE=github.com/mycompany/*                                        │
│ GONOPROXY=github.com/secret/*                                           │
│ GONOSUMDB=github.com/private/*                                          │
│                                                                         │
│ // Proxy benefits:                                                      │
│ • Faster downloads                                                      │
│ • Better availability                                                   │
│ • Immutable versions                                                    │
│ • Security verification                                                 │
└─────────────────────────────────────────────────────────────────────────┘

🔄 MODULE WORKFLOW:
┌─────────────────────────────────────────────────────────────────────────┐
│ 1. Initialize module:                                                   │
│    go mod init github.com/username/project                              │
│                                                                         │
│ 2. Add dependencies:                                                    │
│    go get github.com/gin-gonic/gin                                      │
│                                                                         │
│ 3. Write code using dependencies                                        │
│                                                                         │
│ 4. Clean up:                                                            │
│    go mod tidy                                                          │
│                                                                         │
│ 5. Verify:                                                              │
│    go mod verify                                                        │
│                                                                         │
│ 6. Build/test:                                                          │
│    go build                                                             │
│    go test ./...                                                        │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Use semantic versioning for releases
• Keep go.mod and go.sum in version control
• Run go mod tidy before committing
• Use replace directive for local development
• Set GOPRIVATE for private repositories
• Update dependencies regularly
• Test with multiple Go versions

🚨 COMMON ISSUES:
❌ Not running go mod tidy after changes
❌ Ignoring go.sum file in version control
❌ Using replace in production builds
❌ Not understanding semantic versioning
❌ Mixing GOPATH and module modes

🎯 MIGRATION FROM GOPATH:
1. Move code outside GOPATH
2. Run go mod init in project root
3. Run go mod tidy to convert dependencies
4. Remove vendor/ directory (optional)
5. Update CI/CD scripts
6. Test thoroughly

=============================================================================
*/