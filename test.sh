#!/bin/bash -e

export NAME=ext_authz
export PORT_PROXY="9999"

# shellcheck source=examples/verify-common.sh
. "./verify-common.sh"

run_log "Test services responds with 403"
responds_with_header \
    "HTTP/1.1 401 Unauthorized"\
    "http://localhost:${PORT_PROXY}"


run_log "Test authenticated service responds with 200"
responds_with_header \
    "HTTP/1.1 200 OK" \
    -H "Authorization: Bearer token1" \
    "http://localhost:${PORT_PROXY}/service"