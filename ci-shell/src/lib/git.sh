#!/usr/bin/env bash

# Returns true (0) if current directory is git configured or false (1) otherwise.
function _is_git_repo(){
    #git rev-parse --is-inside-work-tree >/dev/null 2>&1
    if [ "$(git rev-parse --is-inside-work-tree)" = true ]; then
        echo "true" 
        return 0
    else
        echo_std_err "false" 
        return 1
    fi
}

# Returns current directory git workspace (0) or specify error (1) otherwise.
function _get_git_workspace(){
    if [ "$(_is_git_repo)" ]; then
        GIT_WORKSPACE="$(git rev-parse --show-toplevel)" || return 1
        echo "$GIT_WORKSPACE"
        return 0
    else 
        echo_std_err "Not a Valid Git Repo! ❌" 
        return 1
    fi
}

# Returns application name for a valid git repository (0) or errors out (1)
function _get_app_name_from_git_workspace(){
    GIT_WORKSPACE=$(_get_git_workspace) 
    APP_NAME=$(basename "$GIT_WORKSPACE")
    if [ -n "$APP_NAME" ];then
        echo "$APP_NAME"
        export APP_NAME
        return 0
    else 
        echo_std_err "App Name is Empty! ❌" 
        return 1      
    fi

}