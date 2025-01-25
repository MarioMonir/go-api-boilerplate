# Go API Boilerplate

This repository demonstrates a basic Go web server with the following key features:

- **Multiple Handlers:** Supports multiple routes with separate handler functions.

- **Modular Structure:**

  - Handlers are defined in separate files within the `handlers` package.

  - Server logic is encapsulated in the `server` package.

- **Readability and Maintainability:** Promotes good code organization and readability.

**Project Structure:**

```
├── handlers/
│ ├── root.go
│ └── user.go
│
├── server/
│ └── server.go
│
├── main.go
│
└── README.md
```

**How to Run:**

1. **Clone the repository:**

   ```bash
   git clone https://github.com/MarioMonir/go-api-boilerplate
   ```

2. **Run Project:**
   ```bash
    go run .
   ```
