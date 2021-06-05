#!/usr/bin/env bash

function _populate_dev_container_env(){
    prompt "${GREEN} Using workspace ${WORKSPACE} ${NC}"

    ## Load devcontainer.json to CONFIG_JSON
    CONFIG=$(< "$CONFIG_DIR/$CONFIG_FILE"  grep -v // | jq)
    CONFIG_JSON="echo $CONFIG"
    debug "CONFIG : $CONFIG"

    ## Get Dockerfile name and derive Location Name from CONFIG_JSON
    DOCKER_FILE_NAME=$($CONFIG_JSON | jq -r .dockerFile)
    if [ "$DOCKER_FILE_NAME" == "null" ]; then
        DOCKER_FILE_NAME=$($CONFIG_JSON | jq -r .build.dockerfile)
    fi
    DOCKER_FILE="$CONFIG_DIR/$DOCKER_FILE_NAME"
    _file_exist "$CONFIG_DIR/$DOCKER_FILE_NAME" || raise_error "$DOCKER_FILE_NAME Not Found"
    debug "DOCKER_FILE: ${DOCKER_FILE}"

    ## Get Remote User Name from CONFIG_JSON
    REMOTE_USER=$($CONFIG_JSON | jq -r .remoteUser)
    if ! [ "$REMOTE_USER" == "null" ]; then
        REMOTE_USER="-u ${REMOTE_USER}"
    fi
    debug "REMOTE_USER: ${REMOTE_USER}"

    ## Get Build Args from CONFIG_JSON
    ARGS=$($CONFIG_JSON | jq -r '.build.args | to_entries? | map("--build-arg \(.key)=\"\(.value)\"")? | join(" ")')
    ARGS=$(echo "$ARGS" | tr -d '"')
    debug "ARGS: ${ARGS}"

    ## Get Shell from CONFIG_JSON
    SHELL=$($CONFIG_JSON | jq -r '.settings."terminal.integrated.shell.linux"')
    debug "SHELL: ${SHELL}"

    ## Get Ports from CONFIG_JSON
    PORTS=$($CONFIG_JSON | jq -r '.forwardPorts | map("-p \(.):\(.)")? | join(" ")')
    debug "PORTS: ${PORTS}"

    ## Get Envs from CONFIG_JSON
    ENVS=$($CONFIG_JSON | jq -r '.remoteEnv | to_entries? | map("-e \(.key)=\(.value)")? | join(" ")')
    debug "ENVS: ${ENVS}"
    WORK_DIR="/workspace"
    debug "WORK_DIR: ${WORK_DIR}"

    ## Get Mount points from CONFIG_JSON
    MOUNT="${MOUNT} --mount type=bind,source=${WORKSPACE},target=${WORK_DIR}"
    debug "MOUNT: ${MOUNT}"
    debug "ARGS: $ARGS"
}

function build_container(){
    echo "${GREEN} Building container ${NC}"
    start=$(date +%s)
    DOCKER_CMD="docker build -f $DOCKER_FILE -t $DOCKER_IMAGE $ARGS ."
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
