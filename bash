#!/usr/bin/env bash

### Helpers begin
check_deps() {
    local missing
    for d in "${deps[@]}"; do
        if [[ -z $(command -v "$d") ]]; then
            # Force absolute path
            if [[ ! -e "/$d" ]]; then
                err "$d was not found"
                missing="true"
            fi
        fi
    done; unset d
    [[ -z $missing ]] || exit 128
}
err() { echo -e "${color:+\e[31m}[!] $*${color:+\e[0m}" >&2; }
errx() { err "${*:2}"; exit "$1"; }
good() { echo -e "${color:+\e[32m}[+] $*${color:+\e[0m}"; }
info() { echo -e "${color:+\e[37m}[*] $*${color:+\e[0m}"; }
long_opt() {
    local arg shift="0"
    case "$1" in
        "--"*"="*) arg="${1#*=}"; [[ -n $arg ]] || return 127 ;;
        *) shift="1"; shift; [[ $# -gt 0 ]] || return 127; arg="$1" ;;
    esac
    echo "$arg"
    return $shift
}
subinfo() { echo -e "${color:+\e[36m}[=] $*${color:+\e[0m}"; }
warn() { echo -e "${color:+\e[33m}[-] $*${color:+\e[0m}"; }
### Helpers end

usage() {
    cat <<EOF
Usage: ${0##*/} [OPTIONS]

DESCRIPTION
    TODO.

OPTIONS
    -h, --help         Display this help message
        --no-color     Disable colorized output
    -t, --todo         Example flag
    -t, --todo=TODO    Example for storing cli arg

EOF
    exit "$1"
}

declare -a args
unset help
color="true"

# Parse command line options
while [[ $# -gt 0 ]]; do
    case "$1" in
        "--") shift; args+=("$@"); break ;;
        "-h"|"--help") help="true" ;;
        "--no-color") unset color ;;
        # "-t"|"--todo") todo="true" ;;
        # "-t"|"--todo"*) todo="$(long_opt "$@")" ;;
        *) args+=("$1") ;;
    esac
    case "$?" in
        0) ;;
        1) shift ;;
        *) usage $? ;;
    esac
    shift
done
[[ ${#args[@]} -eq 0 ]] || set -- "${args[@]}"

# Help info
[[ -z $help ]] || usage 0

# Check for missing dependencies
declare -a deps
# deps+=("todo")
check_deps

# Check for valid params
[[ $# -eq 0 ]] || usage 1

# TODO
