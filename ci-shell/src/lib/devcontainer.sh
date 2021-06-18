#!/usr/bin/env bash

## IMPORTANT: ToDo: Touching this file with debug and echo - will break functionality 
## as we are returning the values from this function 
## Figure other alternatives
function _query_json(){
    JSON_STRING=$1
    QUERY_STRING=$2
    [ -n "$JSON_STRING" ]  || raise_error "Parameter JSON_STRING is Empty! "
    [ -n "$QUERY_STRING" ] || raise_error "Parameter QUERY_STRING is Empty! "
    QUERY_RESULT=$(echo "$JSON_STRING" | jq -r "$QUERY_STRING")
    echo "$QUERY_RESULT"
    return 0
}

function _get_devcontainer_json(){
    DEV_CONTAINER_JSON_PATH="$(_get_git_workspace)/.devcontainer/devcontainer.json"
    _file_exist "$DEV_CONTAINER_JSON_PATH" || raise_error "devcontainer.json Not Found"
    ## Load devcontainer.json to CONFIG_JSON
    _JSON="$(< "$DEV_CONTAINER_JSON_PATH"  grep -v // )"
    FORMATTED_JSON=$(_query_json "$_JSON" ".")
    echo  "$FORMATTED_JSON"
    return 0
}

# Returns formatted devcontainer.json (0) or errors out  (1) otherwise.
function _get_devcontainer_json_corrected(){
    json=$(_get_devcontainer_json | tail -n +2)
    echo "$json"
}

# Returns dokcer file path (0) or errors out  (1) otherwise.
function _get_docker_file_path(){
    CONFIG_JSON=$1
    # Throw Error if the $CONFIG_JSON is empty
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    ## Get Dockerfile name and derive Location Name from CONFIG_JSON
    DOCKER_FILE_NAME=$(_query_json "$CONFIG_JSON" ".build.dockerfile")
    DOCKER_FILE_PATH="$(_get_git_workspace)/.devcontainer/$DOCKER_FILE_NAME"

    _file_exist "$DOCKER_FILE_PATH" || raise_error "$DOCKER_FILE_PATH Not Found"
    echo "$DOCKER_FILE_PATH"
    export DOCKER_FILE_PATH
    return 0
}

# Returns Remote User Name 
function _get_remote_user(){
    CONFIG_JSON=$1
    # Throw Error if the $CONFIG_JSON is empty
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    ## Get Remote User Name from CONFIG_JSON
    REMOTE_USER=$(_query_json "$CONFIG_JSON" ".remoteUser")
    if ! [ "$REMOTE_USER" == "null" ]; then
        REMOTE_USER="-u ${REMOTE_USER}"
    fi
    echo "${REMOTE_USER}"
    return 0
}

# Returns build args 
function _get_build_args(){
    CONFIG_JSON=$1
    # Throw Error if the $CONFIG_JSON is empty
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    ## Get Build Args from CONFIG_JSON
    ARGS=$(_query_json "$CONFIG_JSON" '.build.args | to_entries? | map("--build-arg \(.key)=\"\(.value)\"")? | join(" ")')
    ARGS=$(echo "$ARGS" | tr -d '"')
    echo "$ARGS"
    return 0
}

# Returns shell 
function _get_shell(){
    CONFIG_JSON=$1
    # Throw Error if the $CONFIG_JSON is empty
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    ## Get Build Args from CONFIG_JSON
    SHELL=$(_query_json "$CONFIG_JSON" '.settings."terminal.integrated.shell.linux"')
    echo "$SHELL"
    return 0
}

# Returns ports 
function _get_port(){
    CONFIG_JSON=$1
    # Throw Error if the $CONFIG_JSON is empty
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    ## Get Build Args from CONFIG_JSON
    PORTS=$(_query_json "$CONFIG_JSON" '.forwardPorts | map("-p \(.):\(.)")? | join(" ")')
    echo "$PORTS"
    return 0
}

# Returns envs 
function _get_envs(){
    CONFIG_JSON=$1
    # Throw Error if the $CONFIG_JSON is empty
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    ## Get Build Args from CONFIG_JSON
    ENVS=$(_query_json "$CONFIG_JSON" '.remoteEnv | to_entries? | map("-e \(.key)=\(.value)")? | join(" ")')
    echo "$ENVS"
    return 0
}

# Returns Mounts 
function _get_mount_points(){
    CONFIG_JSON=$1
    # Throw Error if the $CONFIG_JSON is empty
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    ## Get Build mount points 
    WORK_DIR="/workspaces"
    MOUNT="--mount type=bind,source=${WORKSPACE},target=${WORK_DIR}"
    echo "$MOUNT"
    return 0
}
