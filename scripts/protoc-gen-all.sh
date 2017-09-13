#!/usr/bin/env bash

# Delete all ".pb.go" files in a folder.
deleteGenerated() {
    # If there's no argument, exit.
    if [ -z "$1" ]; then
        return 0
    fi

    pkg=$1
    for file in $pkg/*; do

        # It's a file.
        if [ -f "$file" ]; then

            # It ends with ".pb.go".
            if [ ${file: -6} == ".pb.go" ]; then
                echo "Deleting ${file}"
                $(rm -rf ${file})
            fi
        fi
    done

    return 0
}

# Generate all ".proto" files in a folder.
generate() {
    # If there's no argument, exit.
    if [ -z "$1" ]; then
        return 0
    fi

    pkg=$1
    echo "Generating in ${pkg}"
    $(protoc --go_out=plugins=grpc:$GO_SRC_PATH --proto_path=$GO_SRC_PATH "$pkg"/*.proto)

    return 0
}

traverse() {
    # If there's no argument, exit.
    if [ -z "$1" ]; then
        return 0
    fi

    pkg=$1

    # Check if there are any generated or proto files in $pkg.
    genFile=$(find $pkg -maxdepth 1 -type f -name "*.pb.go" | head -n1)
    protoFile=$(find $pkg -maxdepth 1 -type f -name "*.proto" | head -n1)

    # Delete and generate protobuf files.
    if [ ! -z $genFile ]; then
        deleteGenerated $pkg
    fi
    if [ ! -z $protoFile ]; then
        generate $pkg
    fi

    # Print new line.
    if [ ! -z $genFile ] || [ ! -z $protoFile ]; then
        echo
    fi

    # Recursively traverse any sub folders.
    for d in $pkg/*; do
        if [ -d "$d" ]; then
            traverse $d
        fi
    done
    return 0
}

## Main.

export SCRIPT_PATH=$(cd $(dirname "${BASH_SOURCE[0]}" ) && pwd)
export PROJECT_PATH=$(cd $SCRIPT_PATH && cd .. && pwd)
export GO_SRC_PATH=$GOPATH/src

traverse $PROJECT_PATH
