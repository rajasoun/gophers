#!/usr/bin/env bash

function echo_std_err(){
    echo -e "$@" 1>&2
}

 function _check_installed() {
     command=$1
    # { type curl || type wget; } >/dev/null
    { type "$command"; }  >/dev/null
  }


function check() {
    LABEL=$1
    shift
    if _check_installed "$LABEL" ; then
        echo -e "🧪 Testing $LABEL  Passed! ✅" 
        return 0
    else
        echo_std_err "🧪 Testing $LABEL  Failed! ❌" 
        return 1
    fi
}

function report_results() {
    FAILED=("$@")
    echo "${FAILED[@]}"
    if [ ${#FAILED[@]} -ne 0 ]; then
        echo_std_err "\n💥  Failed tests:" "${#FAILED[@]}"
        # echo_std_err "\nOverall Status: Failed! ❌" 
        return 1
    else
        echo -e "\n💯  All Passed!"
        return 0
    fi
}

function check_common_os_packages(){
   local FAILED=()
   pkgs=(git shellspec kcov jq nc gh curl)
   for pkg in "${pkgs[@]}"
   do 
        check "$pkg" || FAILED+=("$pkg")
   done
   report_results "${FAILED[@]}"
}


