# CREV (Code REView AI)
CLI tool to review code using AI.

## Dependencies

### Development:
## How to run the code

```bash
go run cmd/ai-code-review/main.go
```

## How to run tests
    
```bash
go test ./tests/... -count=1
```

## How to lint the code
Linter:
https://golangci-lint.run/welcome/install/

(MacOS)
```bash
brew install golangci-lint
```
```bash
golangci-lint run
```

## Releasing
Push a new tag to the repository.
```bash
git tag v0.0.1
git push origin v0.0.1
```
