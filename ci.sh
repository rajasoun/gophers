#!/usr/bin/env bash

# Load env variables 
function load_config_from_dotenv(){
    CONFIG_FILE=".env"
    if [ -f "$CONFIG_FILE" ]
    then
        export $(echo $(cat "$CONFIG_FILE" | sed 's/#.*//g' | sed 's/\r//g' | xargs))
        #export $(echo $(cat "$CONFIG_FILE" | sed 's/#.*//g' | sed 's/\r//g' | xargs) | envsubst)
    fi
    return 0
}

function git_config(){
    load_config_from_dotenv
    gh auth status --hostname $GH_HOST
    retVal=$?
    if [ $retVal -ne 0 ]; then
      gh auth login --with-token <<< $GITHUB_TOKEN 
    fi
    if [ ! -f "$HOME/.gitconfig" ]; 
    then
      # printf "User EMail : " && read -r USER_EMAIL
      # printf "User ID : " && read -r USER_ID
      git config --global user.email "$USER_EMAIL"
      git config --global user.name "$USER_ID"
    fi 
}

function delete_all_releases_tags(){
    # Delete all Release
    list=$(gh release list | sed 's/|/ /' | awk '{print $2}')
    for line in $list;
    do
      gh release delete -y "$line"; 
    done
    # Delete all remote tags
    git tag -l | xargs -n 1 git push --delete origin
    # Delete local tags
    git tag | xargs git tag -d
}

function git_delete_latest_release_tag(){
  git fetch --tags # checkout action does not get these
  current_tag=$(git tag --sort=-v:refname --list  | head -n 1)
  echo "current_tag: $current_tag"
  git tag --list "$current_tag" | xargs -I % echo "git tag -d %; git push --delete origin %" | sh
  gh release delete -y "$current_tag"
}

# Load configs by convention
function load_config(){
    load_config_from_dotenv
    #APP_NAME=$(basename "$(git remote get-url origin)")
    APP_NAME=$(basename $(git rev-parse --show-toplevel))
    GIT_COMMIT=$(git log -1 --format=%h)
    CONTAINER="$APP_NAME"
    BUILD_CONTEXT="."
    export APP_NAME CONTAINER BUILD_CONTEXT
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
  _docker build \
      -f "ci-shell/Dockerfile" \
      -t "$CONTAINER" \
      --build-arg GIT_COMMIT="$GIT_COMMIT" \
      "$BUILD_CONTEXT" 
  return 0
}

function shell(){
  _docker run \
          --rm -it \
          --name "$APP_NAME" \
          --hostname "$APP_NAME" \
          -v "$(pwd):$(pwd)" -w "$(pwd)" \
          -v /var/run/docker.sock:/var/run/docker.sock  \
          "$CONTAINER" 
  return 0
}

function run_ci_shell_e2e_tests(){
  _docker run \
          --rm \
          --name "$APP_NAME" \
          --hostname "$APP_NAME" \
          -v "$(pwd):$(pwd)" -w "$(pwd)" \
          -v /var/run/docker.sock:/var/run/docker.sock  \
          "$CONTAINER" sh -c "shellspec -c ci-shell/spec --tag ci-build"
  return 0
}

function run_dev_container_e2e_tests(){
  _docker run \
          --rm \
          --name $APP_NAME \
          --hostname $APP_NAME \
          -v $(pwd):$(pwd) -w $(pwd) \
          -v /var/run/docker.sock:/var/run/docker.sock  \
          $CONTAINER sh -c "ci-shell/dev.sh e2e "
  return 0
}

function teardown(){
  echo  "Clean Container"
  docker rmi "$CONTAINER" 
}

opt="$1"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )
load_config 
case $choice in
    git-config)
      git_config
    ;;
    build)
      echo -e "\nBuild Container"
      build || echo "Docker Build Failed"
    ;;
    shell)
      echo "Run  Container"
      shell || echo "Docker Run Failed"
    ;;
    e2e)
      echo "Test  Container"
      build > /dev/null 2>&1 # Build ci-shell
      run_ci_shell_e2e_tests 
      run_dev_container_e2e_tests  # e2e Test devcontainer within ci-shell
      teardown
    ;;
    teardown)
      teardown
    ;;
    *)
    echo "${RED}Usage: ci.sh  (build | shell | teardown | e2e) [-d]${NC}"
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