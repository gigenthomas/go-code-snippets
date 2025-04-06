# 🚪 Go API Template

A minimal and modular Go API setup using:

- [**Viper**](https://github.com/spf13/viper) – configuration management
- [**Chi**](https://github.com/go-chi/chi) – lightweight and idiomatic HTTP router
- [**Zap**](https://github.com/uber-go/zap) – high-performance structured logging

---

## 📆 Project Structure

```
.
├── config/          # Viper-based configuration loader
├── logger/          # Zap logger initialization
├── middleware/      # Custom middleware (e.g., logger injection)
├── routes/          # Chi route handlers
├── main.go          # Entry point
└── go.mod           # Go module file
```

---

## 🚀 Features

- ✅ Environment-based configuration with Viper
- ✅ Fast and flexible routing using Chi
- ✅ Structured and performant logging with Zap
- ✅ Middleware setup for logging and request context
- ✅ Clean, modular design for easy scalability

---

## 🛠️ Getting Started

### Prerequisites

- Go 1.18+

### Installation

```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
go mod tidy
```

---

## 🧪 Running the App

```bash
go run main.go
```

The server will start on the port defined in your config (default: `:8080`):

```
Listening on http://localhost:8080
```

---

## ⚙️ Configuration with Viper

Create a `config.yaml` in the root or `/conf/` directory:

```yaml
server:
  port: ":8080"

log:
  level: "info"
  encoding: "json"
```

Viper will auto-detect and load the config file and support environment variable overrides.

---

## 🌐 Example Route

```go
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    logger := middleware.GetLogger(r.Context())
    logger.Info("Hello route hit")
    w.Write([]byte("Hello, World!"))
})
```

---

## 📚 Logging with Zap

Use `zap.L()` for the global logger or pass your own:

```go
zap.L().Info("API server started", zap.String("port", cfg.Server.Port))
```

---

## 🤩 Middleware

Example: Injecting logger into request context:

```go
r.Use(middleware.LogMiddleware)
```

Then access it inside any handler:

```go
logger := middleware.GetLogger(r.Context())
```

---

## 📆 Dependencies

| Package | Description                        |
| ------- | ---------------------------------- |
| `viper` | Config management with env support |
| `chi`   | HTTP router for Go                 |
| `zap`   | Structured logging                 |

---



