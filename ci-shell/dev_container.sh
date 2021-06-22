#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=/dev/null
source "$SCRIPT_DIR/src/load.sh"

DOCKER_BUILDKIT=1
DEBUG_OFF="" 
DEBUG_TOGGLE="${2:-$DEBUG_OFF}"
DEV_SHELL="${3:-sfdx}"

export DOCKER_BUILDKIT

init_env_variables "$DEV_SHELL"
_debug_option "$DEBUG_TOGGLE"
check jq
_file_exist "$DEV_CONTAINER_JSON_PATH" 

opt="$1"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )
echo "Starting --> $choice for $DEV_SHELL"
case ${choice} in
    "e2e")
        build_container > /dev/null 2>&1
        e2e_tests
        tear_down
    ;;
    "build")
        build_container
    ;;
    "shell")
        build_container
        shell_2_container
    ;;
    "teardown")
        tear_down
    ;;
    *)
    echo "${RED}Usage: automator/ci.sh <build | e2e | taerdown | shell> [-d]${NC}"
cat <<-EOF
Commands:
---------
  build       -> Build Container
  shell       -> Shell into the Dev Container
  teardown    -> Teardown Dev Container
  e2e         -> Build Dev Container,Run End to End IaaC Test Scripts and Teardown
EOF
    ;;
esac
