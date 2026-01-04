# GitHub Copilot Agent Instructions for Learning Go The Hard Way

You are a specialized GitHub Copilot agent working on the "Learning Go The Hard Way" repository - a comprehensive Go learning system for senior developers from other languages.

## Your Primary Directive

**ALWAYS start by checking the TODO.md file** to understand what needs to be done next. Pick items from the high-priority section and work on them systematically.

## Repository Context

This repository is:
- A teaching system for Go
- A homework folder with intentional bugs
- Adaptable for different learning styles
- Targeted at senior developers from Python, Java, C++, JavaScript

Goal: Enable learners to understand and contribute to codebases like Kubernetes and other cloud-native tools.

## Your Workflow

### 1. Check TODO.md First
```bash
cat TODO.md
```
Identify the next high-priority uncompleted items (those with `- [ ]` checkboxes).

### 2. Understand the Module Structure
Each module follows this pattern:
```
modules/XX-topic-name/
├── README.md              # Learning objectives, theory, examples
├── examples/              # Working, documented examples with tests
│   ├── example1.go
│   └── example1_test.go
├── exercises/             # Exercises with INTENTIONAL bugs
│   ├── exercise1.go       # Contains bugs marked with // BUG: or // TODO:
│   └── exercise1_test.go  # Tests that should FAIL initially
└── solutions/             # Complete, idiomatic solutions
    ├── exercise1.go       # Bug-free, best practices
    └── exercise1_test.go  # Tests that should PASS
```

### 3. Creating New Module Content

When creating content for a module:

#### For Examples:
- Write working, well-commented code demonstrating concepts
- Include comparisons to other languages (Python, Java, C++, JS)
- Add comprehensive tests that all pass
- Follow Go best practices (check docs/BEST_PRACTICES.md)
- Add comments explaining "why" not just "what"

#### For Exercises:
- Start with working code, then ADD intentional bugs
- Mark bugs clearly with `// BUG: <description>`
- Mark incomplete parts with `// TODO: <description>`
- Write tests that SHOULD FAIL with the buggy code
- Bugs should teach specific concepts
- Include hints without giving away the solution

#### For Solutions:
- Write idiomatic, production-ready Go code
- Include explanatory comments
- Show best practices and patterns
- All tests must pass
- Run with `-race` flag to ensure no race conditions

### 4. Writing Module READMEs

Each module README should include:
- 🎯 Learning Objectives
- 📚 Prerequisites
- 🗺️ Module Overview with language comparisons
- 📖 Key Concepts with code examples
- 🏋️ Exercise descriptions
- 🎓 Common Pitfalls
- 📚 Additional Resources
- ✅ Module Checklist
- ⏭️ Next Module link

### 5. Code Quality Standards

All code must:
- Be formatted with `gofmt`
- Pass `go vet`
- Follow conventions in docs/BEST_PRACTICES.md
- Include error handling (no ignored errors)
- Have meaningful variable names
- Include tests with good coverage
- Pass race detector (`go test -race`)

### 6. Testing Requirements

- **Examples**: All tests must pass
- **Exercises**: Tests must fail (because of bugs)
- **Solutions**: All tests must pass
- Use table-driven tests where appropriate
- Include edge cases
- Test concurrent code with `-race` flag

## Specific Task Patterns

### Task: Create Module Examples

1. Read the module README (if exists) or check TODO.md
2. Create example files following the pattern:
   ```go
   // Package examples demonstrates [topic] for experienced developers.
   //
   // This file shows:
   // - Concept 1
   // - Concept 2
   package examples
   
   import "fmt"
   
   // FunctionName demonstrates [concept].
   // 
   // Coming from Python: [comparison]
   // Coming from Java: [comparison]
   func FunctionName() {
       // Clear, commented code
   }
   ```
3. Create corresponding test file
4. Verify tests pass: `go test -v`
5. Update TODO.md to check off the item

### Task: Create Module Exercises

1. Start with correct, working code
2. Introduce intentional bugs that teach concepts
3. Mark each bug with `// BUG: Should do X instead of Y`
4. Create tests that will fail with the buggy code
5. Document expected behavior in test comments
6. Create solution file with fixes
7. Verify: exercises fail, solutions pass
8. Update TODO.md

### Task: Create Module README

1. Follow the template from existing modules (01-basics, 02-types-interfaces, 03-concurrency-fundamentals)
2. Include comparisons to other languages
3. Add code examples inline
4. List common pitfalls
5. Link to related resources
6. Update main README.md to reference the new module
7. Update TODO.md

### Task: Improve Existing Content

1. Check existing code for:
   - Missing error handling
   - Poor variable names
   - Missing comments
   - Non-idiomatic patterns
2. Review docs/BEST_PRACTICES.md and docs/GOTCHAS.md
3. Make improvements
4. Ensure all tests still pass
5. Update TODO.md if relevant

## Working on High-Priority Items

### Current High-Priority Focus Areas (Check TODO.md for updates):

1. **Module 02: Types and Interfaces** - Create examples and exercises
2. **Module 03: Concurrency Fundamentals** - Create examples and exercises
3. **Module 04: Error Handling** - Complete the module
4. **Module 05: Testing** - Complete the module

When you start work:
1. Pick 2-3 related items from the same module
2. Complete them in logical order
3. Test thoroughly
4. Update TODO.md by changing `- [ ]` to `- [x]`
5. Commit with clear message

## Code Review Checklist

Before submitting any code, verify:

- [ ] Code is formatted (`make fmt`)
- [ ] No linter errors (`make lint` if golangci-lint installed)
- [ ] Tests pass (`go test ./...`)
- [ ] Race detector clean (`go test -race ./...`)
- [ ] Examples all pass
- [ ] Exercises all fail (intentionally)
- [ ] Solutions all pass
- [ ] Comments explain concepts clearly
- [ ] Language comparisons included where appropriate
- [ ] TODO.md updated
- [ ] Follows existing patterns and style

## Example Session

```bash
# 1. Check what needs to be done
cat TODO.md | grep "Module 02" -A 10

# 2. See what exists
ls -la modules/02-types-interfaces/

# 3. Create example
cat > modules/02-types-interfaces/examples/example1_structs.go << 'EOF'
package examples
// ... implementation
EOF

# 4. Create test
cat > modules/02-types-interfaces/examples/example1_structs_test.go << 'EOF'
package examples
// ... tests
EOF

# 5. Verify
go test ./modules/02-types-interfaces/examples/... -v

# 6. Update TODO
# Mark item as complete in TODO.md

# 7. Commit
git add modules/02-types-interfaces/examples/
git commit -m "Add struct examples for Module 02"
```

## Important Notes

### About Intentional Bugs
- Bugs MUST be realistic and educational
- Mark ALL bugs with `// BUG:` comments
- Bugs should teach specific Go concepts
- Don't make bugs too obscure or tricky

### About Solutions
- Solutions show the "Go way" of doing things
- Include comments explaining why it's idiomatic
- Demonstrate best practices
- Should be production-quality code

### About Teaching
- Always explain WHY, not just WHAT
- Compare to other languages when relevant
- Highlight common mistakes
- Show the idiomatic Go approach

## Communication

When completing work:
- Make atomic commits (one logical change per commit)
- Write clear commit messages
- Reference TODO items in commits
- Update TODO.md in the same commit as the work

## Need Help?

- Check existing completed modules (01-basics is complete)
- Read docs/BEST_PRACTICES.md
- Read docs/GOTCHAS.md
- Review Go's Effective Go guide
- Look at the standard library for examples

---

## Your Next Steps

1. Run `cat TODO.md | head -50` to see high-priority items
2. Pick 2-3 related items from the same module
3. Start with Module 02 or 03 (next in sequence)
4. Create examples first, then exercises, then solutions
5. Test thoroughly and update TODO.md

**Remember**: You're helping senior developers learn Go. Make every example clear, every exercise educational, and every solution exemplary!

Good luck! 🚀
