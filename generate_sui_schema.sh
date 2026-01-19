#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

QUERY=$(tr '\n' ' ' < sui_mainnet_introspection_query.graphql)

curl -s -X POST https://graphql.mainnet.sui.io/graphql \
  -H "Content-Type: application/json" \
  -d "{\"query\":\"$QUERY\"}" \
  > sui_mainnet_schema.json

