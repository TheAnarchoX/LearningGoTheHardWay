# Quick Start Guide

Welcome to Learning Go The Hard Way! This guide will get you up and running in 5 minutes.

## Installation

### 1. Install Go

**macOS:**
```bash
brew install go
```

**Linux:**
```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

**Windows:**
Download from https://go.dev/dl/ and run the installer.

### 2. Verify Installation
```bash
go version  # Should show go1.21 or later
```

### 3. Clone the Repository
```bash
git clone https://github.com/TheAnarchoX/LearningGoTheHardWay.git
cd LearningGoTheHardWay
```

### 4. Install Development Tools
```bash
make install-tools
```

## Your First 5 Minutes

### Step 1: Run the Examples (2 minutes)
```bash
# See all working examples
make test-module MODULE=01-basics
```

You should see output like:
```
=== RUN   TestVariables
Zero values - name: '', age: 0, isActive: false
...
PASS
```

### Step 2: Try to Fix a Bug (2 minutes)

Open `modules/01-basics/exercises/exercise1_fix_bugs.go`:

```go
// BUG: This function has a logic error.
func CalculateSum(a, b int) int {
    return a - b // BUG: Should be addition, not subtraction
}
```

Fix the bug by changing `-` to `+`:
```go
func CalculateSum(a, b int) int {
    return a + b  // Fixed!
}
```

Test your fix:
```bash
cd modules/01-basics/exercises
go test -v -run TestCalculateSum
```

Should now show:
```
=== RUN   TestCalculateSum
--- PASS: TestCalculateSum (0.00s)
```

### Step 3: Compare with Solution (1 minute)
```bash
diff exercises/exercise1_fix_bugs.go solutions/exercise1_fix_bugs.go
```

See what the idiomatic solution looks like!

## Learning Paths

### Path 1: Linear (Recommended for Go Beginners)
Start with Module 01 and work sequentially:
1. Read the module README
2. Study the examples
3. Fix the exercises
4. Compare with solutions
5. Move to next module

```bash
cd modules/01-basics
cat README.md
```

### Path 2: Topic-Based (For Experienced Developers)
Jump to topics you want to learn:

**Want to learn concurrency?**
```bash
cd modules/03-concurrency-fundamentals
cat README.md
```

**Want to understand interfaces?**
```bash
cd modules/02-types-interfaces
cat README.md
```

### Path 3: Debug-Driven (For Hands-On Learners)
See all failing tests and start fixing:
```bash
make test  # See what's broken
cd modules/01-basics/exercises
go test -v  # Focus on one module
```

## Daily Workflow

### Morning: Learn New Concepts
```bash
cd modules/XX-topic
cat README.md  # Read theory
go test ./examples/... -v  # Run examples
```

### Afternoon: Practice
```bash
cd modules/XX-topic/exercises
go test -v  # See what fails
# Fix bugs...
go test -v  # Verify fixes
```

### Evening: Review
```bash
diff exercises/ solutions/  # Compare implementations
cat ../README.md  # Review concepts
```

## Common Commands

```bash
# Run all tests
make test

# Run tests for one module
make test-module MODULE=01-basics

# Format your code
make fmt

# Build examples
make build

# Generate coverage report
make coverage

# Clean build artifacts
make clean
```

## IDE Setup

### VS Code (Recommended)
1. Install Go extension
2. Open workspace: `code LearningGoTheHardWay`
3. Install recommended extensions when prompted

### GoLand
1. Open project directory
2. Enable Go modules support
3. Run tests from the IDE

### Vim/Neovim
```bash
# Install vim-go
git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go
# Use :GoTest, :GoBuild, etc.
```

## Getting Help

### In This Repository
- Check module README files
- Read `docs/BEST_PRACTICES.md`
- Read `docs/GOTCHAS.md`
- Review solution code

### External Resources
- [A Tour of Go](https://go.dev/tour/) - Interactive tutorial
- [Go Playground](https://go.dev/play/) - Try code online
- [Go by Example](https://gobyexample.com/) - Code examples
- [Effective Go](https://go.dev/doc/effective_go) - Style guide

### Community
- Open an issue if you find a real bug
- Start a discussion for questions
- Check existing issues first

## Tips for Success

1. **Do the exercises yourself first** - Don't peek at solutions!
2. **Run tests frequently** - Get immediate feedback
3. **Read error messages carefully** - Go's errors are helpful
4. **Use the race detector** - `go test -race`
5. **Format your code** - `make fmt` before committing
6. **Read other people's code** - Solutions show best practices

## Next Steps

After the quick start:

1. **Complete Module 01** - Foundation is crucial
2. **Read BEST_PRACTICES.md** - Learn idiomatic Go
3. **Read GOTCHAS.md** - Avoid common mistakes
4. **Move to Module 02** - Build on the foundation
5. **Join the community** - Share your progress!

## Troubleshooting

### Tests won't run
```bash
# Ensure Go is installed
go version

# Download dependencies
go mod download

# Try running a simple test
go test ./modules/01-basics/examples/...
```

### Linter not found
```bash
make install-tools
```

### Can't understand an error
1. Read the error message fully
2. Check the line number it references
3. Look at the test to see what's expected
4. Compare with the solution

---

**Ready to master Go?** Start with `cd modules/01-basics` and let's go! 🚀
