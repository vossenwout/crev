# CREV (Code REView AI)
CLI tool to review code using AI.


## installation
Linux/MacOS binaries
```bash
sudo curl -L https://raw.githubusercontent.com/vossenwout/crev/feature/add-install-scripts/scripts/install.sh | bash
```

Brew (MacOS / Linux)
```bash
brew install vossenwout/crev/crev
```

windows (Run as Administrator in powershell)
```bash
Invoke-Expression (Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/vossenwout/crev/feature/add-install-scripts/scripts/install.ps1').Content
```

## Releasing
Push a new tag to the repository.
```bash
git tag v0.0.1
git push origin v0.0.1
```
