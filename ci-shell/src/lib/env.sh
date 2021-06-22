#!/usr/bin/env bash

function init_env_variables(){
    VERBOSE=0

    DEFAULT_SHELL="go"
    DEV_SHELL="${1:-$DEFAULT_SHELL}"

    GIT_WORKSPACE="$(git rev-parse --show-toplevel)"
    APP_NAME=$(basename "$GIT_WORKSPACE")
    WORKSPACE="${GIT_WORKSPACE}/vscode-iaac"

    CONFIG_DIR="${WORKSPACE}/${DEV_SHELL}/.devcontainer"
    debug "CONFIG_DIR : $CONFIG_DIR"
    CONFIG_FILE="devcontainer.json"
    DOCKER_IMAGE="vsc-$APP_NAME:$(git rev-parse HEAD)"
    debug "DOCKER_IMAGE : $DOCKER_IMAGE"

    DEV_CONTAINER_JSON_PATH="$CONFIG_DIR/$CONFIG_FILE"
    debug "DEV_CONTAINER_JSON_PATH : $DEV_CONTAINER_JSON_PATH"

    export VERBOSE DEV_SHELL  GIT_WORKSPACE APP_NAME WORKSPACE
    export CONFIG_DIR CONFIG_FILE DOCKER_IMAGE
    export DEV_CONTAINER_JSON_PATH
}
init_env_variables "go"
