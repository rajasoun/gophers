#!/usr/bin/env bash

## To get all functions : bash -c "source src/load.bash && declare -F"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=/dev/null
source "$SCRIPT_DIR/lib/e2e.sh"
# shellcheck source=/dev/null
source "$SCRIPT_DIR/lib/devcontainer.sh"
# shellcheck source=/dev/null
source "$SCRIPT_DIR/lib/docker.sh"
# shellcheck source=/dev/null
source "$SCRIPT_DIR/lib/git.sh"
# shellcheck source=/dev/null
source "$SCRIPT_DIR/lib/os.sh"

