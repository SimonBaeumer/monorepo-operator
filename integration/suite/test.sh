#!/usr/bin/env bash

$dir=$(dirname $0)

./../suite/setup.sh
./../suite/test-init.sh
./../suite/test-sync.sh
./../suite/test-sync-tag.sh
./../suite/test-remove-branches.sh