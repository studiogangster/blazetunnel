#!/usr/bin/env bash

package_name="blazetunnel"

platforms=("windows/amd64/windows/x64" "windows/386/windows/x86" "darwin/amd64/mac/x64" "darwin/386/mac/x86"  "linux/amd64/linux/x64" "linux/386/linux/x86" )

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    FOLDER=${platform_split[2]}
    VARIANT=${platform_split[3]}
    output_name='apps/'$FOLDER'/'$VARIANT'/'$package_name #'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name ./../
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done