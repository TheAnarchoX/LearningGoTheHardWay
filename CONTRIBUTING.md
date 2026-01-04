# Contributing to Learning Go The Hard Way

Thank you for your interest in contributing! This guide will help you understand how to contribute effectively.

## 🎯 Philosophy

This repository is designed to be:
1. **Interactive** - Learning through doing and fixing
2. **Intentionally buggy** - Some bugs are features for learning
3. **Progressive** - Building knowledge from basics to advanced
4. **Practical** - Focused on real-world applications

## 🐛 Reporting Issues

### Real Bugs (Please Report!)
- Typos in documentation or comments
- Incorrect explanations or outdated information
- Broken test infrastructure
- Missing dependencies
- Build/tooling issues

### Intentional Bugs (Don't Report!)
These are marked with:
- `// BUG:` comments
- `// TODO:` comments
- `// EXERCISE:` comments
- Failing tests in exercise directories

If unsure, check if there's a corresponding solution file that fixes it.

## 📝 Adding New Content

### Adding a New Module

1. Follow the existing module structure:
```
modules/XX-topic-name/
├── README.md           # Module overview and learning objectives
├── examples/           # Working, documented examples
│   ├── example1.go
│   └── example1_test.go
├── exercises/          # Exercises with intentional bugs
│   ├── exercise1.go
│   ├── exercise1_test.go  # These tests should fail initially
│   └── HINTS.md       # Optional hints
└── solutions/          # Complete, idiomatic solutions
    ├── exercise1.go
    └── exercise1_test.go  # These tests should pass
```

2. Module README should include:
   - Learning objectives
   - Prerequisites
   - Comparison to other languages (Python, Java, C++, JS)
   - Key concepts
   - Common pitfalls
   - Additional resources

3. Examples should:
   - Be runnable and well-documented
   - Include tests
   - Follow Go best practices
   - Include comments explaining "why" not just "what"

4. Exercises should:
   - Have clear objectives
   - Include intentional bugs or incomplete implementations
   - Have failing tests that pass when fixed
   - Include hints for common pitfalls

5. Solutions should:
   - Be idiomatic Go code
   - Include explanatory comments
   - Pass all tests
   - Show best practices

### Adding Exercises to Existing Modules

1. Create the exercise file with intentional bugs
2. Mark bugs clearly with `// BUG:` or `// TODO:` comments
3. Write failing tests
4. Create a solution file with correct implementation
5. Ensure solution tests pass
6. Update module README with the new exercise

### Code Style

All code must:
- Be formatted with `gofmt`
- Pass `golangci-lint` checks
- Include meaningful comments
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use meaningful variable names
- Include error handling

### Documentation Style

- Use clear, concise language
- Include code examples
- Compare to patterns in other languages when relevant
- Explain the "why" behind Go's approach
- Link to official Go documentation

## 🔍 Review Process

1. **Submit a Pull Request** with:
   - Clear description of changes
   - Why the change is beneficial
   - Which module(s) it affects
   - Whether it's new content or a fix

2. **Ensure CI passes**:
   - All tests pass
   - Code is formatted
   - Linters pass

3. **Address feedback** from reviewers

## 🧪 Testing Your Contributions

Before submitting:

```bash
# Format code
make fmt

# Run linters
make lint

# Test your module
make test-module MODULE=XX-topic-name

# Test everything
make test

# Build examples
make build
```

## 💡 Ideas for Contributions

### Content
- New modules on advanced topics
- More real-world examples
- Additional exercises
- Better explanations for complex concepts
- Comparisons to more languages

### Infrastructure
- Improved tooling
- Better test coverage
- Documentation improvements
- CI/CD enhancements

### Community
- Video tutorials
- Blog posts explaining concepts
- Translations
- Study group materials

## 🤔 Questions?

- Open an issue for questions
- Start a discussion for ideas
- Check existing issues and PRs first

## 📜 Code of Conduct

Be respectful, constructive, and helpful. We're all learning!

## 🙏 Recognition

Contributors will be acknowledged in:
- README.md
- Module-specific documentation
- Release notes

Thank you for helping others learn Go! 🚀
