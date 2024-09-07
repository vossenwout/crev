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


## If you want to set up pre-commit hooks

1. Setup virtual environment
```bash
python -m venv .venv
```

2. Activate virtual environment
```bash
source .venv/bin/activate
```

3. Install pre-commit
```bash
pip install pre-commit
```

4. Install pre-commit hooks
```bash
pre-commit install
```

5. Run pre-commit hooks
```bash
pre-commit run --all-files
```


