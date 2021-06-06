#!/usr/bin/env bash

NC=$'\e[0m' # No Color
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'

# Displays text in colors
function all_colors() {
  debug "${RED}RED${NC}"
  debug "${GREEN}GREEN${NC}"
  debug "${ORANGE}ORANGE${NC}"
  debug "${BLUE}BLUE${NC}"
  debug "${LIGHT_BLUE}LIGHT_BLUE${NC}"
  debug "${BOLD}BOLD${NC}"
  debug "${UNDERLINE}UNDERLINE${NC}"
}

# Wrapper function for echo to stderr
function echo_std_err(){
    echo -e "$@" 1>&2
}

# promt error message and exit
function raise_error(){
  echo -e "${RED}${1} | Exiting...${NC}" >&2
  exit 1
}

# set debug option based on option
function _debug_option(){
  opt="$1" # if -d for debug mode
  choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )
  case ${choice} in
      -d) VERBOSE=1 ;;
  esac
}

# echo message when VERBOSE == 1
function debug(){
  message=$1
    if [ "$VERBOSE" == 1 ]; then
        printf "${ORANGE}\n [DEBUG] %s${NC}\n" "${message}"
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

# Returns true (0) if this the given file path is found or false (1) otherwise.
function _file_exist() {
    FILE=$1
    if [ -e "$FILE" ]; then
        echo -e "$FILE  Found! ✅" 
        return 0
    else
        echo_std_err "$FILE  Not Found! ❌" 
        return 1
    fi
}

