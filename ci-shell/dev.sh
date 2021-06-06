#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=/dev/null
source "$SCRIPT_DIR/src/load.sh"

VERBOSE=0
export VERBOSE
GIT_WORKSPACE="$(git rev-parse --show-toplevel)"
APP_NAME=$(basename "$GIT_WORKSPACE")
WORKSPACE="${GIT_WORKSPACE}"
CONFIG_DIR="${WORKSPACE}/.devcontainer"
CONFIG_FILE="devcontainer.json"
DOCKER_IMAGE="vsc-$APP_NAME:$(git rev-parse HEAD)"
debug "DOCKER_IMAGE : $DOCKER_IMAGE"


check jq
_file_exist "$CONFIG_DIR/$CONFIG_FILE" 
# option -d parameter for debug
_debug_option "$2"

opt="$1"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )
case ${choice} in
    "e2e")
        export DOCKER_BUILDKIT=1
        build_container
        e2e_tests
        tear_down
    ;;
    "build")
        build_container
    ;;
    "shell")
        _populate_dev_container_env
        export DOCKER_BUILDKIT=1
        build_container
        shell_2_container
    ;;
    "teardown")
        tear_down
    ;;
    *)
    echo "${RED}Usage: automator/ci.sh <e2e | taerdown | shell> [-d]${NC}"
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
