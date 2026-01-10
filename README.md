# Go Programming - Complete Reference Guide

> A comprehensive, interview-ready guide to Go programming with real-world examples, practice questions, and best practices.

**Author:** Kazi Foyez Ahmed  
**Last Updated:** January 2026  
**Go Version:** 1.22+

---

## 📚 Table of Contents

### 1. [Introduction to Go](./sections/1-introduction.md)
- [1.1 What is Go?](./sections/1-introduction.md#11-what-is-go)
- [1.2 Why Go?](./sections/1-introduction.md#12-why-go)
- [1.3 Go vs Other Languages](./sections/1-introduction.md#13-go-vs-other-languages)
- [1.4 Setup & Installation](./sections/1-introduction.md#14-setup--installation)
- [1.5 First Go Program](./sections/1-introduction.md#15-first-go-program)
- [1.6 Essential Commands](./sections/1-introduction.md#16-essential-commands)

### 2. [Fundamentals](./sections/2-fundamentals.md)
- [2.1 Package Structure](./sections/2-fundamentals.md#21-package-structure)
- [2.2 Variables & Constants](./sections/2-fundamentals.md#22-variables--constants)
- [2.3 Data Types](./sections/2-fundamentals.md#23-data-types)
- [2.4 Type Conversion](./sections/2-fundamentals.md#24-type-conversion)
- [2.5 Operators](./sections/2-fundamentals.md#25-operators)

### 3. [Strings & I/O](./sections/3-strings-io.md)
- [3.1 String Basics](./sections/3-strings-io.md#31-string-basics)
- [3.2 String Operations](./sections/3-strings-io.md#32-string-operations)
- [3.3 Input/Output (fmt, log)](./sections/3-strings-io.md#33-inputoutput)
- [3.4 File Operations](./sections/3-strings-io.md#34-file-operations)

### 4. [Control Flow](./sections/4-control-flow.md)
- [4.1 if-else Statements](./sections/4-control-flow.md#41-if-else-statements)
- [4.2 switch Statements](./sections/4-control-flow.md#42-switch-statements)
- [4.3 Loops (for, range)](./sections/4-control-flow.md#43-loops)
- [4.4 defer Statement](./sections/4-control-flow.md#44-defer-statement)

### 5. [Functions](./sections/5-functions.md)
- [5.1 Function Basics](./sections/5-functions.md#51-function-basics)
- [5.2 Multiple Return Values](./sections/5-functions.md#52-multiple-return-values)
- [5.3 Variadic Functions](./sections/5-functions.md#53-variadic-functions)
- [5.4 Closures & Anonymous Functions](./sections/5-functions.md#54-closures--anonymous-functions)
- [5.5 Methods](./sections/5-functions.md#55-methods)

### 6. [Data Structures](./sections/6-data-structures.md)
- [6.1 Arrays](./sections/6-data-structures.md#61-arrays)
- [6.2 Slices](./sections/6-data-structures.md#62-slices)
- [6.3 Maps](./sections/6-data-structures.md#63-maps)
- [6.4 Structs](./sections/6-data-structures.md#64-structs)
- [6.5 Pointers](./sections/6-data-structures.md#65-pointers)

### 7. [Interfaces & Polymorphism](./sections/7-interfaces.md)
- [7.1 Interface Basics](./sections/7-interfaces.md#71-interface-basics)
- [7.2 Implicit Implementation](./sections/7-interfaces.md#72-implicit-implementation)
- [7.3 Empty Interface](./sections/7-interfaces.md#73-empty-interface)
- [7.4 Type Assertions & Type Switch](./sections/7-interfaces.md#74-type-assertions--type-switch)
- [7.5 Interface Composition](./sections/7-interfaces.md#75-interface-composition)
- [7.6 Interface Best Practices](./sections/7-interfaces.md#76-interface-best-practices)

### 8. [Error Handling](./sections/8-error-handling.md)
- [8.1 Error Basics](./sections/8-error-handling.md#81-error-basics)
- [8.2 Error Wrapping](./sections/8-error-handling.md#82-error-wrapping)
- [8.3 Custom Errors](./sections/8-error-handling.md#83-custom-errors)
- [8.4 Panic & Recover](./sections/8-error-handling.md#84-panic--recover)
- [8.5 Best Practices](./sections/8-error-handling.md#85-best-practices)

### 9. [Concurrency](./sections/9-concurrency.md)
- [9.1 Goroutines](./sections/9-concurrency.md#91-goroutines)
- [9.2 Channels](./sections/9-concurrency.md#92-channels)
- [9.3 Select Statement](./sections/9-concurrency.md#93-select-statement)
- [9.4 Synchronization (WaitGroup, Mutex)](./sections/9-concurrency.md#94-synchronization)
- [9.5 Context](./sections/9-concurrency.md#95-context)
- [9.6 Concurrency Patterns](./sections/9-concurrency.md#96-concurrency-patterns)

### 10. [Generics (Go 1.18+)](./sections/10-generics.md)
- [10.1 Introduction to Generics](./sections/10-generics.md#101-introduction-to-generics)
- [10.2 Type Parameters](./sections/10-generics.md#102-type-parameters)
- [10.3 Type Constraints](./sections/10-generics.md#103-type-constraints)
- [10.4 Generic Functions](./sections/10-generics.md#104-generic-functions)
- [10.5 Generic Types](./sections/10-generics.md#105-generic-types)
- [10.6 Generic Interfaces](./sections/10-generics.md#106-generic-interfaces)
- [10.7 Best Practices](./sections/10-generics.md#107-best-practices)

### 11. [Testing](./sections/11-testing.md)
- [11.1 Unit Testing](./sections/11-testing.md#111-unit-testing)
- [11.2 Table-Driven Tests](./sections/11-testing.md#112-table-driven-tests)
- [11.3 Benchmarking](./sections/11-testing.md#113-benchmarking)
- [11.4 Mocking & Interfaces](./sections/11-testing.md#114-mocking--interfaces)
- [11.5 Test Coverage](./sections/11-testing.md#115-test-coverage)
- [11.6 TDD with Go](./sections/11-testing.md#116-tdd-with-go)

### 12. [Web Development](./sections/12-web-development.md)
- [12.1 HTTP Server Basics](./sections/12-web-development.md#121-http-server-basics)
- [12.2 Routing](./sections/12-web-development.md#122-routing)
- [12.3 Middleware](./sections/12-web-development.md#123-middleware)
- [12.4 JSON Handling](./sections/12-web-development.md#124-json-handling)
- [12.5 REST API Example](./sections/12-web-development.md#125-rest-api-example)
- [12.6 File Serving](./sections/12-web-development.md#126-file-serving)

### 13. [Best Practices](./sections/13-best-practices.md)
- [13.1 Google Style Guide](./sections/13-best-practices.md#131-google-style-guide)
- [13.2 Code Organization](./sections/13-best-practices.md#132-code-organization)
- [13.3 Naming Conventions](./sections/13-best-practices.md#133-naming-conventions)
- [13.4 Error Handling Patterns](./sections/13-best-practices.md#134-error-handling-patterns)
- [13.5 Performance Tips](./sections/13-best-practices.md#135-performance-tips)
- [13.6 Common Pitfalls](./sections/13-best-practices.md#136-common-pitfalls)

### 14. [Interview Preparation](./sections/14-interview-questions.md)
- [14.1 Basics & Fundamentals (Q1-Q10)](./sections/14-interview-questions.md#141-basics--fundamentals)
- [14.2 Data Structures (Q11-Q15)](./sections/14-interview-questions.md#142-data-structures)
- [14.3 Concurrency (Q16-Q20)](./sections/14-interview-questions.md#143-concurrency)
- [14.4 Error Handling (Q21-Q22)](./sections/14-interview-questions.md#144-error-handling)
- [14.5 Interfaces (Q23-Q25)](./sections/14-interview-questions.md#145-interfaces--polymorphism)
- [14.6 Performance (Q26-Q28)](./sections/14-interview-questions.md#146-performance--optimization)
- [14.7 Testing (Q29-Q30)](./sections/14-interview-questions.md#147-testing)
- [14.8 Web Servers (Q31-Q32)](./sections/14-interview-questions.md#148-web-servers)
- [14.9 Best Practices (Q33)](./sections/14-interview-questions.md#149-best-practices)
- [14.10 Coding Challenges (Q34-Q35)](./sections/14-interview-questions.md#1410-coding-challenges)
- [14.11 Generics (Q36-Q40)](./sections/14-interview-questions.md#1411-generics)

### 15. [Resources](./sections/15-resources.md)
- [15.1 Official Documentation](./sections/15-resources.md#151-official-documentation)
- [15.2 Books](./sections/15-resources.md#152-books)
- [15.3 Video Tutorials](./sections/15-resources.md#153-video-tutorials)
- [15.4 Practice Platforms](./sections/15-resources.md#154-practice-platforms)
- [15.5 Community](./sections/15-resources.md#155-community)

---

## 🎯 How to Use This Guide

### For Beginners
1. Start with **[Introduction](./sections/1-introduction.md)** to understand what Go is and why it matters
2. Follow **[Fundamentals](./sections/2-fundamentals.md)** through **[Control Flow](./sections/4-control-flow.md)** in order
3. Practice with the exercises at the end of each section
4. Build small projects to reinforce learning

### For Interview Preparation
1. Review **[Interview Preparation](./sections/14-interview-questions.md)** for common questions
2. Focus on **[Concurrency](./sections/9-concurrency.md)** and **[Interfaces](./sections/7-interfaces.md)** (frequently tested)
3. Practice coding challenges in section 14.10
4. Review **[Best Practices](./sections/13-best-practices.md)** for code quality questions

### For Experienced Developers
1. Jump to specific topics using the navigation above
2. Use as a quick reference during development
3. Review **[Generics](./sections/10-generics.md)** for Go 1.18+ features
4. Consult **[Best Practices](./sections/13-best-practices.md)** for production code

### For Web Developers
1. Review **[Web Development](./sections/12-web-development.md)** for HTTP servers
2. Study **[Concurrency](./sections/9-concurrency.md)** for handling concurrent requests
3. Check **[Context](./sections/9-concurrency.md#95-context)** for request management
4. Review error handling for robust APIs

---

## 📝 Features

### ✅ Real-World Examples
Every concept includes practical, production-ready code examples from domains like:
- E-commerce systems
- Banking applications
- User management
- API development

### ✅ Practice Questions
Each section includes:
- **Fill-in-the-blanks** - Test vocabulary and syntax
- **True/False** - Verify understanding
- **Multiple Choice** - Assess concept knowledge
- **Code Challenges** - Apply skills practically

### ✅ Interview-Ready
- 40+ common interview questions with detailed answers
- Coding challenges with complete solutions
- Best practices aligned with industry standards
- Real interview scenarios and explanations

### ✅ Visual Aids
- Diagrams explaining complex concepts
- Memory layouts for data structures
- Flow charts for algorithms
- Comparison tables

### ✅ Complete Coverage
- Go keywords (all 25)
- Standard library essentials
- Concurrency patterns
- Web development
- Testing strategies
- Performance optimization

---

## 🚀 Quick Start

```bash
# Clone or download this repository
git clone https://github.com/foyez/go

# Navigate to any chapter
# Each file is self-contained and can be read independently

# Start with the introduction
cat 1-introduction.md

# Or jump to a specific topic
cat 9-concurrency.md
```

---

## 📖 Navigation Tips

- **Sequential Learning**: Follow chapters 1-13 in order for comprehensive learning
- **Topic-Based**: Jump to specific topics using the table of contents above
- **Search**: Use Ctrl+F (Cmd+F on Mac) to search within files
- **Cross-References**: Each section links to related topics
- **Practice**: Complete practice questions before moving to the next chapter

---

## 🎓 Learning Path

### Week 1: Foundations
- Day 1-2: Introduction, Setup, Fundamentals
- Day 3-4: Strings, I/O, Control Flow
- Day 5-7: Functions, Practice Projects

### Week 2: Data & OOP
- Day 1-3: Data Structures (Arrays, Slices, Maps)
- Day 4-5: Structs, Pointers
- Day 6-7: Interfaces, Error Handling

### Week 3: Advanced
- Day 1-3: Concurrency (Goroutines, Channels)
- Day 4-5: Generics, Testing
- Day 6-7: Web Development

### Week 4: Mastery
- Day 1-2: Best Practices, Code Review
- Day 3-5: Interview Preparation
- Day 6-7: Build a Complete Project

---

## 🔄 Updates & Contributions

This guide is regularly updated to reflect:
- Latest Go features and best practices
- Community feedback and improvements
- New real-world examples
- Additional interview questions

---

## 📌 Key Topics for Interviews

**Most Frequently Asked:**
1. **Goroutines vs Threads** → [Section 9.1](./sections/9-concurrency.md#91-goroutines)
2. **Channels & Select** → [Section 9.2](./sections/9-concurrency.md#92-channels), [9.3](./sections/9-concurrency.md#93-select-statement)
3. **Interfaces** → [Chapter 7](./sections/7-interfaces.md)
4. **Slices vs Arrays** → [Section 6.1](./sections/6-data-structures.md#61-arrays), [6.2](./sections/6-data-structures.md#62-slices)
5. **Error Handling** → [Chapter 8](./sections/8-error-handling.md)
6. **Context** → [Section 9.5](./sections/9-concurrency.md#95-context)

**Coding Challenges:**
- LRU Cache → [Q34](./sections/14-interview-questions.md#q34-implement-a-lru-cache-in-go)
- Rate Limiter → [Q35](./sections/14-interview-questions.md#q35-implement-a-rate-limiter-using-go)
- Generic Stack → [Q39](./sections/14-interview-questions.md#q39-can-you-write-a-generic-stack-in-go)

---

## 🎯 Goals of This Guide

1. **Interview Readiness**: Prepare for technical interviews with comprehensive Q&A
2. **Practical Learning**: Learn through real-world examples, not just theory
3. **Quick Reference**: Serve as a go-to resource during development
4. **Best Practices**: Follow industry-standard patterns and conventions
5. **Complete Coverage**: Cover everything from basics to advanced topics

---

## 📚 Additional Resources

- **Official Go Documentation**: [golang.org/doc](https://golang.org/doc)
- **Go Playground**: [play.golang.org](https://play.golang.org)
- **Go Blog**: [blog.golang.org](https://blog.golang.org)
- **Effective Go**: [golang.org/doc/effective_go](https://golang.org/doc/effective_go)
- **Google Style Guide**: [google.github.io/styleguide/go](https://google.github.io/styleguide/go)

---

## 💡 Tips for Success

1. **Practice Daily**: Write code every day, even if just for 30 minutes
2. **Build Projects**: Apply concepts in real projects
3. **Read Code**: Study well-written Go code (Docker, Kubernetes, etc.)
4. **Join Community**: Participate in Go forums and discussions
5. **Test Yourself**: Complete practice questions after each chapter

---

**Happy Learning! 🚀**

For questions or suggestions, please open an issue or contribute to the repository.

---

**License**: MIT  
**Maintained by**: Kazi Foyez Ahmed  
**Repository**: https://github.com/foyez/go