#!/usr/bin/env bash

function build_container(){
    CONFIG_JSON=$(_get_devcontainer_json_corrected)
    [ -n "$CONFIG_JSON" ] || raise_error "Parameter CONFIG_JSON is Empty! "
    debug "${GREEN} Building container ${NC}"
    check "docker"  || raise_error "Docker Not Installed. Exiting... "
    
    BUILD_CONTEXT="."
    DOCKER_FILE=$(_get_docker_file_path "$(_get_devcontainer_json_corrected)")
    BUILD_ARGS=$(_get_build_args "$CONFIG_JSON")
    debug "BUILD_ARGS : $BUILD_ARGS"
    DOCKER_IMAGE="vsc-$(_get_app_name_from_git_workspace):$(git rev-parse HEAD)"
    debug "DOCKER_IMAGE : $DOCKER_IMAGE"

    start=$(date +%s)
    DOCKER_CMD="docker build -f $DOCKER_FILE -t $DOCKER_IMAGE $BUILD_ARGS $BUILD_CONTEXT"
    debug "Docker Build Command : $DOCKER_CMD"
    $DOCKER_CMD
    end=$(date +%s)
    runtime=$((end-start))
    echo "Container Local Build Success ✅ : $(_display_time $runtime)"
}

function shell_2_container(){
    prompt "${GREEN} Starting container ${NC}"
    DOCKER_RUN_OPTS="$REMOTE_USER $PORTS $ENVS $MOUNT -w $WORK_DIR $DOCKER_IMAGE $SHELL"
    DOCKER_CMD="docker run --sig-proxy=false -a STDOUT -a STDERR  --rm -it $DOCKER_RUN_OPTS"
    debug "Docker Run Command : $DOCKER_CMD"
    $DOCKER_CMD
}

function e2e_tests(){
    prompt "${GREEN} Starting container ${NC}"
    DOCKER_RUN_OPTS="$REMOTE_USER $PORTS $ENVS $MOUNT -w $WORK_DIR $DOCKER_IMAGE $SHELL"
    DOCKER_CMD="docker run --sig-proxy=false -a STDOUT -a STDERR  --rm -it $DOCKER_RUN_OPTS"
    debug "Docker Run Command : $DOCKER_CMD -c "shellspec -c ci-shell/spec --kcov""
    $DOCKER_CMD -c "shellspec -c ci-shell/spec --kcov"
}

function tear_down(){
    if docker ps | awk -v app="$DOCKER_IMAGE" 'NR > 1 && $NF == app{ret=1; exit} END{exit !ret}'; then
        prompt "Stopping Container : $DOCKER_IMAGE ..."
        docker stop "$DOCKER_IMAGE" && docker rm -f "$DOCKER_IMAGE"
    fi
    DOCKER_IMAGE_ID=$(docker images | grep "$APP_NAME" | awk '{print $3}') 2>/dev/null
    if [ -n "$DOCKER_IMAGE_ID"  ]; then
        debug "DOCKER_IMAGE_ID : $DOCKER_IMAGE_ID"
        prompt "Removing Container Image : $DOCKER_IMAGE_ID"
        docker rmi "$DOCKER_IMAGE_ID" -f
    fi
}
