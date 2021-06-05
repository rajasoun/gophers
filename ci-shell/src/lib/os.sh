#!/usr/bin/env bash

NC=$'\e[0m' # No Color
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'

GIT_CONFIG_FILE="$HOME/.gitconfig"
KEYS_PATH="ssh-keys"
PRIVATE_KEY="$KEYS_PATH/id_rsa_cisco_github"
PUBLIC_KEY="${PRIVATE_KEY}.pub"

# Create Directory if the given directory does not exists
# @param $1 The directory to create
function _create_directory_if_not_exists() {
  DIR_NAME=$1
  ## Create Directory If Not Exists
  if [ ! -d "$DIR_NAME" ]; then
    mkdir -p "$DIR_NAME"
  fi
}

# Displays Time in misn and seconds
function _display_time {
  local T=$1
  local D=$((T / 60 / 60 / 24))
  local H=$((T / 60 / 60 % 24))
  local M=$((T / 60 % 60))
  local S=$((T % 60))
  ((D > 0)) && printf '%d days ' $D
  ((H > 0)) && printf '%d hours ' $H
  ((M > 0)) && printf '%d minutes ' $M
  ((D > 0 || H > 0 || M > 0)) && printf 'and '
  printf '%d seconds\n' $S
}

# Returns true (0) if the given file exists contains the given text and false (1) otherwise. The given text is a
# regular expression.
function _file_contains_text {
  local -r text="$1"
  local -r file="$2"
  grep -q "$text" "$file"
}

# Replace a line of text that matches the given regular expression in a file with the given replacement.
# Only works for single-line replacements.
function _file_replace_text {
  local -r original_text_regex="$1"
  local -r replacement_text="$2"
  local -r file="$3"

  local args=()
  args+=("-i")

  if _is_os_darwin; then
    # OS X requires an extra argument for the -i flag (which we set to empty string) which Linux does no:
    # https://stackoverflow.com/a/2321958/483528
    args+=("")
  fi

  args+=("s|$original_text_regex|$replacement_text|")
  args+=("$file")

  sed "${args[@]}" >/dev/null
}

# Returns true (0) if this is an OS X server or false (1) otherwise.
function _is_os_darwin {
  [[ $(uname -s) == "Darwin" ]]
}

# Returns true (0) if this the given command/app is installed and on the PATH or false (1) otherwise.
function _is_command_found {
  local -r name="$1"
  command -v "$name" >/dev/null ||
    raise_error "${RED}$name is not installed. Exiting...${NC}"
}

# Wrapper function for echo
function prompt() {
  echo -e "${1}" >&2
}

# Example colour function
function all_colors() {

  debug "${RED}RED${NC}"
  debug "${GREEN}GREEN${NC}"
  debug "${ORANGE}ORANGE${NC}"
  debug "${BLUE}BLUE${NC}"
  debug "${LIGHT_BLUE}LIGHT_BLUE${NC}"
  debug "${BOLD}BOLD${NC}"
  debug "${UNDERLINE}UNDERLINE${NC}"

  prompt "$BOLD $UNDERLINE Colour Formatting Example $NC"
  prompt "$RED R $NC $GREEN A  $NC $BLUE J $NC $RED A. $BLUE $NC S"
}

# ls, with chmod-like permissions and more.
# @param $1 The directory to ls
function lls() {
  #If Not Paramter is Passed assumes current directory
  local LLS_PATH=${1:-"."}
  prompt "${GREEN} ls with chmod-like permissions ${NC}"
  # shellcheck disable=SC2012 # Reason: This is for human consumption
  ls -AHl "$LLS_PATH" | awk "{k=0;for(i=0;i<=8;i++)k+=((substr(\$1,i+2,1)~/[rwx]/) \
                            *2^(8-i));if(k)printf(\"%0o \",k);print}"
}

# Run Pre Commit and Git Add on Changed Files
function run_pre_commit() {
  pre-commit run --all-files
  git diff --staged --name-only --diff-filter=ARM | xargs git add
  exit 0
}

# echo message when VERBOSE == 1
function debug(){
  message=$1
    if [ "$VERBOSE" == 1 ]; then
        printf "${ORANGE}\n [DEBUG] %s${NC}\n" "${message}"
    fi
}

# set debug option based on option
function _debug_option(){
  opt="$1" # if -d for debug mode
  choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )
  echo "choice: $choice"
  case ${choice} in
      -d) VERBOSE=1 ;;
  esac
}

# Check Connection
function _check_connection(){
  server=$1
  port=$1
  if nc -z "$server" "$port" 2>/dev/null; then
      echo "$server on $port ✓"
      return 0
  else
      echo "$server on $port  ✗"
      return 1
  fi
}

function _copy_to_clipboard(){
  CONTENT=$1
  MSG=""
  case "$OSTYPE" in
      *msys*|*cygwin*) 
        os="$(uname -o)" 
        if [[ "$os" == "Msys" ]] || [[ "$os" == "Cygwin" ]]; then
          clip < "$CONTENT"
          MSG="Copied content To Windows Clipboard"
          debug "$MSG"
        fi 
        ;;
      Darwin)
        os="$(uname -s)"
        pbcopy < "$CONTENT"
        MSG="Copied content To macOS Clipboard"
        debug "$MSG"
        ;;
      *) 
        os="$(uname -s)"
        debug "${ORANGE}Headless Linux - Manual Copy Required${NC}"
        echo ""
        cat "$CONTENT"
        echo ""
        MSG="Copy Public Key from Terminal..."
        debug "$MSG"
        ;;
  esac
  echo -e  "${GREEN}$MSG${NC}\n"
}

# Prompt to User for Continue or Exit
function _prompt_confirm(){
    # call with a prompt string or use a default
    local response msg="${1:-Do you want to continue} (y/[n])? "; shift
    read -r "$@" -p "$msg" response || echo
    case "$response" in
        [yY][eE][sS]|[yY]) 
            return 0
            ;;
        [nN][no][No]|[nN]) 
            echo "Exiting setup"
            exit 1
            ;;
        *)
            return 1
            ;;
    esac
}

function cci-env-fix(){
	#ToDo: Technical Debt - Refactor - Move the Function to core - Priority : P3
	if [ -f automator/config/cci.env.key ]; then
		prompt "CCI KEY Env Setup"
    # shellcheck disable=SC2002
    # shellcheck disable=SC2046
		export $(cat automator/config/cci.env.key | xargs)
		prompt "export $(cat automator/config/cci.env.key)"
		prompt "DONE !!!"
	else
		rm -fr .cci .sfdx ~/.cumulusci ~/.sfdx
		prompt "${RED}automator/config/cci.env.key does not exists${NC}"
		prompt "${BLUE}Garbage Collecting previous cci and sfdx files${NC}"
		rm -fr .cci .sfdx ~/.cumulusci ~/.sfdx
		prompt "${RED}RUN: automator/cci-deploy.sh${NC}"
	fi

}

function _backup_remove_git_config(){
    if [ -f "$GIT_CONFIG_FILE" ]; then 
        echo "Backing Up $GIT_CONFIG_FILE to $GIT_CONFIG_FILE.bak"
        cp "$GIT_CONFIG_FILE" "$GIT_CONFIG_FILE.bak"
        echo "Removing $GIT_CONFIG_FILE"
        rm -fr "$GIT_CONFIG_FILE"
    fi 
}

function _git_config(){
    _backup_remove_git_config
    if [[ "$USER" == "vscode" ]]; then
        echo "Executing Inside Dev Container. Getting Cisco User"
        MSG="${GREEN} Cisco CEC User ${NC}${ORANGE}(without eMail) : ${NC}"
        read -r -p "$MSG" USER_NAME
    fi
    if [ -n "$USER_NAME" ]; then 
        git config --global user.name "${USER_NAME}"
        git config --global user.email "${USER_NAME}@cisco.com"
        git config --global core.editor "code"
    fi
    debug "$(cat "$GIT_CONFIG_FILE")"
    echo "Git Config for $USER_NAME Done !!!"
}

function _generate_ssh_keys(){
    debug "Backing up $KEYS_PATH to $KEYS_PATH.bak"
    rm -fr "$KEYS_PATH.bak"
    [   -d "$KEYS_PATH" ] && mv "$KEYS_PATH" "$KEYS_PATH.bak"
    [ ! -d "$KEYS_PATH" ] && mkdir -p "$KEYS_PATH"

    echo "Generating SSH Keys for $USER_NAME"
    _is_command_found ssh-keygen
    debug "Generating SSH Keys for $USER_NAME"
    ssh-keygen -q -t rsa -N '' -f "$PRIVATE_KEY" -C "$USER_NAME@cisco.com" <<<y 2>&1 >/dev/null
    
    echo "Set File Permissions"
    # Fix Permission For Private Key
    chmod 400 "$PUBLIC_KEY"
    chmod 400 "$PRIVATE_KEY"
    debug "SSH Keys Generated Successfully"
    debug "SSH Key Scan for Cisco GitHub Successfull" 
}

function _prompt_vpn_connection(){
    echo -e  "${GREEN}Connect to Cisco VPN... ${NC}\n"
    prompt_confirm "Is VPN Connected"
}

function _print_details(){
  debug ""
  debug "========= PUBLIC KEY ============"
  debug "$(cat "$PUBLIC_KEY")"
  debug "======= END PUBLIC KEY ========="

  echo "GoTo:"
  echo ""
  echo "https://www-github.cisco.com/settings/ssh/new"
  echo ""
}

function _configure_ssh(){
    debug "Git Config"
    _git_config

    debug "SSH Key Gneration"
    _generate_ssh_keys

    echo "Copying SSH Public Key to Clipboard"
    _copy_to_clipboard "$PUBLIC_KEY"
    _print_details
    _check_connection "www-github.cisco.com" || _prompt_vpn_connection
    _prompt_confirm "Is SSH Public Added to Cisco GitHub"
}

# ToDo: Technical Debt : Git SSH Fix : Priority P3
function git-ssh-fix(){
	ERROR_MSG="${RED}Private SSH Key Not Present. DONT PANIC.${NC}"
	NEXT_STEP="Run ${GREEN}config${NC} again... Exiting"
  	MSG="$ERROR_MSG \n $NEXT_STEP"
	[[ ! -f "$PRIVATE_KEY" ]] && echo -e "$MSG" && return 1
	# Check In Terminal
	# ssh-add -l > /dev/null || (eval $(ssh-agent -s) && ssh-add $PRIVATE_KEY)
	echo "${BOLD}Git SSH Hack Fix${NC}"
	# check if ssh key is already added
	ssh-add -l > /dev/null || echo "SSH Key "
	if [ "$(ssh-add -l | wc -l )" = 0 ]; then 
		echo "Adding SSH Key"
		eval "$(ssh-agent -s)" && ssh-add $PRIVATE_KEY
	else
		echo "${GREEN}SSH Key Already Added. Hack Fix Not Required !!!${NC}"
		ssh-add -l
	fi
}

# Wrapper To Aid TDD
function _run_main() {
  _create_directory_if_not_exists "$@"
  _is_command_found "$@"
  _display_time "$@"
  _file_exists "$@"
  _file_contains_text "$@"
  _file_replace_text "$@"
  _is_os_darwin "$@"
  _debug_option "$@"
  _check_connection "$@"
  _copy_to_clipboard "$@"
  _prompt_confirm "$@"
  _backup_remove_git_config "$@"
  _git_config "$@"
  _generate_ssh_keys "$@"
  _prompt_vpn_connection "$@"
  _print_details "$@"
  _configure_ssh "$@"

  debug "$@"
  prompt "$@"
  all_colors "$@"
  lls "$@"
  run_pre_commit "$@"
  cci-env-fix "$@"
  git-ssh-fix "$@"

}

# Wrapper To Aid TDD
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  if ! _run_main "$@"; then
    exit 1
  fi
fi
