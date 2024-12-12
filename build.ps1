$platforms = @("windows/amd64", "linux/amd64", "darwin/amd64", "darwin/arm64")
$outputNames = @("wallet-scrambler-windows-amd64.exe", "wallet-scrambler-linux-amd64", "wallet-scrambler-macos-amd64", "wallet-scrambler-macos-arm64")

for ($i = 0; $i -lt $platforms.Count; $i++) {
    $platform = $platforms[$i]
    $output = $outputNames[$i]
    $split = $platform -split "/"
    $env:GOOS = $split[0]
    $env:GOARCH = $split[1]
    if (Test-Path -Path $output) {
        Remove-Item -Path $output -Force
        Write-Host "Removed existing file: $output"
    }
    & go build -o $output
    Write-Host "Built $output"
}
