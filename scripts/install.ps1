# Fetch the latest version
$versionUrl = "https://api.github.com/repos/vossenwout/crev/releases/latest"
try {
    $response = Invoke-RestMethod -Uri $versionUrl
    $VERSION = $response.tag_name
} catch {
    Write-Host "Error: Failed to fetch the latest version."
    return
}

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
try {
    Invoke-WebRequest -Uri $downloadUrl -OutFile $destination
    Write-Host "Download completed."
} catch {
    Write-Host "Error: Failed to download $FILE."
    return
}

Write-Host "Extracting $destination..."
try {
    Expand-Archive -Path $destination -DestinationPath $env:TEMP\crev -Force
    Write-Host "Extraction completed."
} catch {
    Write-Host "Error: Failed to extract $destination."
    return
}

# Move the binary to a directory in the PATH (C:\Program Files by default)
$installPath = "C:\Program Files\crev"
if (!(Test-Path -Path $installPath)) {
    try {
        New-Item -ItemType Directory -Path $installPath
    } catch {
        Write-Host "Error: Failed to create installation directory at $installPath."
        return
    }
}

try {
    Move-Item "$env:TEMP\crev\crev.exe" "$installPath\crev.exe" -Force
    Write-Host "crev.exe moved to $installPath."
} catch {
    Write-Host "Error: Failed to move crev.exe to $installPath."
    return
}

# Optionally add to PATH if not already
if (-not ([Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine) -contains $installPath)) {
    try {
        [Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine) + ";$installPath", [System.EnvironmentVariableTarget]::Machine)
        Write-Host "crev path added to system PATH. You may need to restart your terminal."
    } catch {
        Write-Host "Error: Failed to update system PATH."
        return
    }
}

# Cleanup
try {
    Remove-Item $destination -Force
    Remove-Item "$env:TEMP\crev" -Recurse -Force
    Write-Host "Cleanup completed."
} catch {
    Write-Host "Error: Failed to clean up temporary files."
}

Write-Host "crev has been installed successfully!"
