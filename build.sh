#!/bin/bash

platforms=("windows/amd64" "linux/amd64" "darwin/amd64" "darwin/arm64")
output_names=("wallet-scrambler-windows-amd64.exe" "wallet-scrambler-linux-amd64" "wallet-scrambler-macos-amd64" "wallet-scrambler-macos-arm64")

for i in "${!platforms[@]}"; do
    platform=${platforms[$i]}
    output=${output_names[$i]}
    
    GOOS=${platform%/*}
    GOARCH=${platform#*/}
    
    GOOS=$GOOS GOARCH=$GOARCH go build -o $output
    if [ $? -eq 0 ]; then
        echo "Built $output"
    else
        echo "Failed to build $output" >&2
        exit 1
    fi

    signature_file="${output}.asc"
    if [ -f "$signature_file" ]; then
        rm -f "$signature_file"
        echo "Removed existing signature: $signature_file"
    fi
    
    gpg --detach-sign --armor -o "$signature_file" "$output"
    if [ $? -eq 0 ]; then
        echo "Successfully signed $output -> $signature_file"
    else
        echo "Failed to sign $output" >&2
        exit 1
    fi
done
