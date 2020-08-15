# go-uuid: Simple command line utility to generate UUIDs

## Installation

    go get github.com/eugeneching/go-uuid/cmd/uuid

## Usage

    Usage of uuid:
          --from-ulid   generate UUID from ULID
          --uuidv1      generate UUID v1
          --uuidv4      generate UUID v4 (default)
      -u, --uppercase   output uppercase
      -n, --number      number of UUIDs to generate (default 1)

## Example

    uuid --from-ulid -n 10
