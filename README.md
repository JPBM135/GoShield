# First GO Project

This web application hashes the input string using the specified algorithm.

## Usage

```bash
go run src/*.go
```

## Endpoints

### POST /hash

The default algorithm is SHA512. To specify a different algorithm, use the `algorithm` query parameter.

> Supported algorithms: sha256, sha512, md5, sha1
