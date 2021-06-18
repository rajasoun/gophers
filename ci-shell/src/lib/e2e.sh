#!/usr/bin/env bash

# End to End Validation test for container
function _check_installed() {
     command=$1
    # { type curl || type wget; } >/dev/null
    { type "$command"; }  >/dev/null
  }

# Returns true (0) if this the given command/app is installed and on the PATH or false (1) otherwise.
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

# Returns true (0) if this the given command/app is installed and on the PATH or false (1) otherwise.
function check_detail() {
    LABEL=$1
    shift
    echo -e "\n🧪 Testing $LABEL"
    if "$@"; then
        echo "✅  Passed!"
        return 0
    else
        echo_std_err "❌ $LABEL check Failed!"
        return 1
    fi
}

function check_system_details() {
    #check_detail "non-root-user" id "${USER}"
    check_detail "locale" [ "$(locale -a | grep en_US.utf8)" ]
    check_detail "sudo" sudo echo "sudo works."
    check_detail "zsh" zsh --version
    check_detail "oh-my-zsh" [ -d "$HOME/.oh-my-zsh" ]
    check_detail "login-shell-path" [ -f "/etc/profile.d/00-restore-env.sh" ]
    check_detail "code" which code
}

function report_results() {
    FAILED=("$@")
    echo "${FAILED[@]}"
    if [ ${#FAILED[@]} -ne 0 ]; then
        echo_std_err "\n💥  Failed tests:" "${#FAILED[@]}"
        return 1
    else
        echo -e "\n💯  All Passed!"
        return 0
    fi
}

function check_dev_packages(){
   local FAILED=()
   pkgs=(git go shellspec kcov jq nc gh curl python3)
   for pkg in "${pkgs[@]}"
   do 
        check "$pkg" || FAILED+=("$pkg")
   done
   report_results "${FAILED[@]}"
}

function check_toolz_packages(){
   local FAILED=()
   pkgs=(husky commitlint commitizen shellcheck)
   for pkg in "${pkgs[@]}"
   do 
        check "$pkg" || FAILED+=("$pkg")
   done
   report_results "${FAILED[@]}"
}

# Check common os packages are installed
# 
function _common_packages() {
    local FAILED=()
    LABEL=$1
    shift
    echo -e "\n🧪 Testing $LABEL"
    if dpkg-query --show -f='${Package}: ${Version}\n' "$@"; then
        echo "💯  All Passed! ✅"
        return 0
    else
        echo_std_err "$LABEL check Failed! ❌"
        FAILED+=("$LABEL")
        return 1
    fi
}

function check_common_packages(){
    PACKAGE_LIST="apt-utils \
        git \
        openssh-client \
        less \
        iproute2 \
        procps \
        curl \
        wget \
        unzip \
        nano \
        jq \
        lsb-release \
        ca-certificates \
        apt-transport-https \
        dialog \
        gnupg2 \
        libc6 \
        libgcc1 \
        libgssapi-krb5-2 \
        liblttng-ust0 \
        libstdc++6 \
        zlib1g \
        locales \
        sudo"

    echo -e "\n🧪 Testing ${PACKAGE_LIST}"
    # shellcheck disable=SC2086
    _common_packages "common-os-packages" ${PACKAGE_LIST}
}





