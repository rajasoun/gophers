#!/usr/bin/env bash

APP_NAME="ci-shell"
CONTAINER="$APP_NAME:latest"

# Injecting env variables 
function load_config(){
    CONFIG_FILE=$1
    # shellcheck disable=SC2002,SC2005,SC2046
    # export $(cat .env | sed 's/#.*//g' | xargs)
    if [ -f "$CONFIG_FILE" ]
    then
        export $(echo $(cat "$CONFIG_FILE" | sed 's/#.*//g' | sed 's/\r//g' | xargs) | envsubst)
    fi
    return 0
}

function help(){
    echo "Usage: $0  {build|shell|clean}" >&2
    echo
    echo "   build           Build Container"
    echo "   shell           Run Container"
    echo "   clean           Stop and Remove Container"
    echo "   test            Build, Run Test and Clean"
    echo
    return 1
}

# Replace a line of text that matches the given regular expression in a file with the given replacement.
# Only works for single-line replacements.
function file_replace_text {
  local -r original_text_regex="$1"
  local -r replacement_text="$2"
  local -r file="$3"

  local args=()
  args+=("-i")

  if os_is_darwin; then
    # OS X requires an extra argument for the -i flag (which we set to empty string) which Linux does no:
    # https://stackoverflow.com/a/2321958/483528
    args+=("")
  fi

  args+=("s|$original_text_regex|$replacement_text|")
  args+=("$file")

  sed "${args[@]}" > /dev/null
  return 0
}

# Workaround for Path Limitations in Windows
function _docker() {
  export MSYS_NO_PATHCONV=1
  export MSYS2_ARG_CONV_EXCL='*'

  case "$OSTYPE" in
      *msys*|*cygwin*) os="$(uname -o)" ;;
      *) os="$(uname)";;
  esac

  if [[ "$os" == "Msys" ]] || [[ "$os" == "Cygwin" ]]; then
      # shellcheck disable=SC2230
      realdocker="$(which -a docker | grep -v "$(readlink -f "$0")" | head -1)"
      printf "%s\0" "$@" > /tmp/args.txt
      file_replace_text "/var" "//var" "/tmp/args.txt"
      # --tty or -t requires winpty
      if grep -ZE '^--tty|^-[^-].*t|^-t.*' /tmp/args.txt; then
          #exec winpty /bin/bash -c "xargs -0a /tmp/args.txt '$realdocker'"
          winpty /bin/bash -c "xargs -0a /tmp/args.txt '$realdocker'"
          return 0
      fi
  fi
  docker "$@"
  return 0
}

function build(){
  docker build -f $APP_NAME/Dockerfile -t $CONTAINER .
  return 0
}

function shell(){
  app_name=$(basename "$(git remote get-url origin)")
  _docker run --rm -it --name $APP_NAME --hostname $APP_NAME \
          -v $(pwd):$(pwd) -w $(pwd) \
          -v /var/run/docker.sock:/var/run/docker.sock  \
          $CONTAINER
  return 0
}

opt="$1"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )
load_config ".env"
case $choice in
    build)
      echo -e "\nBuild Container"
      build || echo "Docker Build Failed"
      ;;
    shell)
      echo "Run  Container"
      shell || echo "Docker Run Failed"
      ;;
    test)
      echo "Test  Container"
      build # Build ci-shell
      ci-shell/dev.sh e2e # e2e Test devcontainer within ci-shell
      docker rmi "$CONTAINER" || echo "Docker Image Remove Failed"
      ;;
    clean)
      echo  "Clean Container"
      docker rmi "$CONTAINER" || echo "Docker Image Remove Failed"
      ;;
    *)  help ;;
esac