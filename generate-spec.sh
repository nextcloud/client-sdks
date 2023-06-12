#!/bin/bash
set -euxo pipefail
cd "$(dirname "$0")"

# shellcheck disable=SC2046
./openapi-extractor/merge-specs \
    --first-status-code \
		--core ./server/core/openapi.json \
		--merged openapi.json \
		$(ls ./server/apps/*/openapi.json)
