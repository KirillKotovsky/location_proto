

# 📦 location_proto

This repository contains [Protocol Buffers](https://protobuf.dev/) definitions used as shared API contracts between microservices in the `location_*` project family.

> 🛑 Generated `.pb.go` files are excluded from version control — all consumers must generate them manually or via CI.

---

## ⚙️ Requirements

To generate Go code from `.proto` definitions, ensure the following are installed:

- [`protoc`](https://grpc.io/docs/protoc-installation/) (Protocol Buffers compiler)
- Go >= 1.22
- Go plugins for protoc:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

Ensure `$GOPATH/bin` is in your `PATH`.

---

## 🧰 Option 1: CI-based Generation (Recommended)

GitHub Actions workflow automatically runs on tag or manual dispatch:

- Installs `protoc` and plugins
- Runs `make generate` to generate `.pb.go` files
- Ensures the project builds via `go build ./...`

To trigger manually:

1. Go to the **Actions** tab
2. Run the `Release Proto (Manual)` workflow

---

## 💻 Option 2: Local Generation

### On Unix/macOS:

```bash
brew install protobuf  # if not installed
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin"

make generate
```

### On Windows (PowerShell):

1. [Download protoc binaries](https://github.com/protocolbuffers/protobuf/releases)
2. Add `protoc.exe` to your `PATH`
3. Install plugins:
   ```powershell
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
4. Add `%GOPATH%\bin` to `PATH`
5. Run:
   ```powershell
   make generate
   ```

---

## 📂 Output

Generated files will appear alongside `.proto` definitions:

- `proto/location/location.pb.go`
- `proto/location/location_grpc.pb.go`

These files are required in any Go project importing this module.

---

## 📌 Notes

- Always re-run `make generate` after modifying `.proto` files
- Avoid committing generated files — this repo uses `.gitignore` to enforce that

---

## 🧪 CI Workflow Summary

| Step               | Description                      |
|--------------------|----------------------------------|
| `make generate`    | Generates gRPC and Go code       |
| `go mod tidy`      | Updates dependency graph         |
| `go build ./...`   | Verifies compilability           |

---

## 🧭 License

MIT — free to use and modify.