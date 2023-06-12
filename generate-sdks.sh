#!/bin/bash
set -euxo pipefail
cd "$(dirname "$0")"

common_options=(-i openapi.json --skip-validate-spec --git-user-id nextcloud --git-repo-id api-sdk)

# Python
rm -rf python
# shellcheck disable=SC2086
openapi-generator-cli generate ${common_options[*]} -o python -g python-nextgen

# Typescript
rm -rf typescript
# shellcheck disable=SC2086
openapi-generator-cli generate ${common_options[*]} -o typescript -g typescript-axios --additional-properties=npmName=nextcloud-api-sdk
(cd typescript && npm install)

# Go
rm -rf go
# shellcheck disable=SC2086
openapi-generator-cli generate ${common_options[*]} -o go -g go
(cd go && go get -u)

# Rust
rm -rf rust
# shellcheck disable=SC2086
openapi-generator-cli generate ${common_options[*]} -o rust -g rust --additional-properties=supportMultipleResponses=true
(cd rust && cargo update)
