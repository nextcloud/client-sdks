# Nextcloud client SDKs

## Important note

The OpenAPI specifications are currently not generated from the source code available on the server master branch.  
This means there is a risk that the specifications might not match the server code 100%.  
In the worst case your API call will fail, but there shouldn't be any bad consequences.

We are working on generating the specifications from the server code and these SDKs too.  
In the meantime you can help us to test the specifications and SDKs, but be aware that there might be slight changes in the future while we work on generating the specifications from the source code.  
We always welcome feedback about this work-in-progress project, so please feel free to open new issues.

You can also generate client code for the language of your choice, and we might add more languages to this repository later.

## Prerequisites

```bash
git submodule update --init

(cd openapi-extractor && composer update && composer install)
```

## Generating the code

```bash
./generate-spec.sh
./generate-sdks.sh
```

## Examples

See the [examples/](examples/) directory for how to use the code.
