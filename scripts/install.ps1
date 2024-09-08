# Fetch the latest version
$versionUrl = "https://api.github.com/repos/vossenwout/crev/releases/latest"
$response = Invoke-RestMethod -Uri $versionUrl
$VERSION = $response.tag_name
$BASE_URL = "https://github.com/vossenwout/crev/releases/download/$VERSION"

# Detect architecture
if ([System.Environment]::Is64BitOperatingSystem) {
    if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
        $ARCH = "arm64"
    } else {
        $ARCH = "x86_64"
    }
} else {
    $ARCH = "i386"
}

# Form download URL
$FILE = "crev_Windows_${ARCH}.zip"

Write-Host "Downloading $FILE from $BASE_URL..."

# Download and extract the binary
$downloadUrl = "$BASE_URL/$FILE"
$destination = "$env:TEMP\crev.zip"
Invoke-WebRequest -Uri $downloadUrl -OutFile $destination

Write-Host "Extracting $destination..."
Expand-Archive -Path $destination -DestinationPath $env:TEMP\crev -Force

# Move the binary to a directory in the PATH (C:\Program Files by default)
$installPath = "C:\Program Files\crev"
if (!(Test-Path -Path $installPath)) {
    New-Item -ItemType Directory -Path $installPath
}
Move-Item "$env:TEMP\crev\crev.exe" "$installPath\crev.exe"

# Optionally add to PATH if not already
if (-not ([Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine) -contains $installPath)) {
    [Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine) + ";$installPath", [System.EnvironmentVariableTarget]::Machine)
    Write-Host "crev path added to system PATH. You may need to restart your terminal."
}

# Cleanup
Remove-Item $destination -Force
Remove-Item "$env:TEMP\crev" -Recurse -Force

Write-Host "crev has been installed successfully!"
