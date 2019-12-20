#!/usr/bin/env bash
set -e

./../suite/setup.sh
./../suite/test-init.sh
./../suite/test-sync.sh
./../suite/test-sync-tag.sh
./../suite/test-remove-branches.sh
./../suite/test-clone.sh
