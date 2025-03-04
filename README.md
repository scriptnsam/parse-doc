# API Docs Generator (MVP)

### 🚀 Automate API Documentation for Go, Python, and JS files

This CLI tool:
✅ Extracts function definitions  
✅ Uses AI (Cohere API) to generate descriptions  
✅ Outputs docs in **Markdown**

## 📦 Installation

```sh
go build -o parse-doc main.go
```

## 👷 Usage

```sh
parse-doc /path/to/code/files
```

## Example Output

```md
### Function: addNumbers(a, b)

**Description:** Adds two numbers and returns the result.
**Parameters:**

- `a`: First number
- `b`: Second number
  **Returns:** Sum of `a` and `b`
```
