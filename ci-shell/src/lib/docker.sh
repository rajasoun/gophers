#!/usr/bin/env bash

_docker="sudo docker"
#ToDo: Technical Debt: Priority-1 
#      Use this Function to elininate complex logic for getting environment function
function load_env_variables(){
    CONFIG_JSON=$(_get_devcontainer_json_corrected)
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    debug "${GREEN} Building container ${NC}"
    check "docker"  || raise_error "Docker Not Installed. Exiting... "

    _get_docker_file_path "$CONFIG_JSON" 
    debug "DOCKER_FILE_PATH: $DOCKER_FILE_PATH"

    BUILD_CONTEXT="."
    BUILD_ARGS=$(_get_build_args "$CONFIG_JSON")
    debug "BUILD_ARGS : $BUILD_ARGS"
    DOCKER_IMAGE="vsc-$(_get_app_name_from_git_workspace):$(git rev-parse HEAD)"
    debug "DOCKER_IMAGE : $DOCKER_IMAGE"

    REMOTE_USER=$(_get_remote_user "$CONFIG_JSON")
    SHELL=$(_get_shell "$CONFIG_JSON")
    PORTS=$(_get_port "$CONFIG_JSON")
    ENVS=$(_get_envs "$CONFIG_JSON")
    MOUNT=$(_get_mount_points "$CONFIG_JSON")
    APP_NAME=$(_get_app_name_from_git_workspace)
    WORK_DIR="${PWD}"
    echo "true"
    export CONFIG_JSON DOCKER_FILE_PATH BUILD_CONTEXT BUILD_ARGS DOCKER_IMAGE
    export REMOTE_USER PORTS ENVS MOUNT APP_NAME SHELL WORK_DIR
    return 0
}

function build_container(){
    load_env_variables
    DOCKER_CMD="$_docker build -f $DOCKER_FILE_PATH -t $DOCKER_IMAGE $BUILD_ARGS $BUILD_CONTEXT"
    debug "Docker Build Command : $DOCKER_CMD"
    $DOCKER_CMD 2>/dev/stdout || raise_error "Container Build Failed"
    return 0
}

function build_container_and_time_it(){
    start=$(date +%s)
    build_container
    end=$(date +%s)
    runtime=$((end-start))
    echo "Container Local Build Success ✅ : $(_display_time $runtime)" 2>/dev/stdout
    return 0
}

function shell_2_container(){
    echo "${GREEN} Starting container ${NC}"
    MOUNT="$MOUNT -v $(pwd):$(pwd)"
    DOCKER_RUN_OPTS="$REMOTE_USER $PORTS $ENVS $MOUNT -w $WORK_DIR $DOCKER_IMAGE $SHELL"
    DOCKER_CMD="$_docker run --sig-proxy=false -a STDOUT -a STDERR  --rm -it $DOCKER_RUN_OPTS"
    debug "Docker Run Command : $DOCKER_CMD"
    $DOCKER_CMD
}

function e2e_tests(){
    load_env_variables
    echo "${GREEN} Starting container ${NC}"
    MOUNT="$MOUNT -v $(pwd):$(pwd)"
    DOCKER_RUN_OPTS="$REMOTE_USER $PORTS $ENVS $MOUNT -w $WORK_DIR $DOCKER_IMAGE $SHELL"
    DOCKER_CMD="$_docker run --sig-proxy=false -a STDOUT -a STDERR  --rm $DOCKER_RUN_OPTS"
    debug "Docker Run Command : $DOCKER_CMD -c shellspec -c ci-shell/spec --tag unit,integration,iaac --kcov"
    $DOCKER_CMD -c "shellspec -c ci-shell/spec --tag unit,integration,iaac --kcov"
}

function tear_down(){
    load_env_variables
    if $_docker ps | awk -v app="$DOCKER_IMAGE" 'NR > 1 && $NF == app{ret=1; exit} END{exit !ret}'; then
        echo "Stopping Container : $DOCKER_IMAGE ..."
        $_docker stop "$DOCKER_IMAGE" && $_docker rm -f "$DOCKER_IMAGE" 2>/dev/stdout
    fi
    DOCKER_IMAGE_ID=$($_docker images | grep "$APP_NAME" | awk '{print $3}') 2>/dev/null
    if [ -n "$DOCKER_IMAGE_ID"  ]; then
        debug "DOCKER_IMAGE_ID : $DOCKER_IMAGE_ID"
        echo "Removing Container Image : $DOCKER_IMAGE_ID"
        $_docker rmi "$DOCKER_IMAGE_ID" -f
    fi
    echo "true"
    return 0
}
