#!/bin/bash

platforms=("windows/amd64" "linux/amd64" "darwin/amd64" "darwin/arm64")
output_names=("wallet-scrambler-windows-amd64.exe" "wallet-scrambler-linux-amd64" "wallet-scrambler-macos-amd64" "wallet-scrambler-macos-arm64")

for i in "${!platforms[@]}"; do
    platform=${platforms[$i]}
    output=${output_names[$i]}
    GOOS=${platform%/*} GOARCH=${platform#*/} go build -o $output
    echo "Built $output"
done
