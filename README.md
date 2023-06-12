# Nextcloud client SDKs [PROOF OF CONCEPT]

This is a repository to show the ability to generated client SDKs in different languages for interacting with Nextcloud APIs.

**Please do not use this code in a production environment, it's merely an experiment**.

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
