#!/usr/bin/env bash

err() { echo -e "${color:+\e[31m}[!] $@\e[0m"; }

errx() { echo -e "${color:+\e[31m}[!] ${@:2}\e[0m"; exit $1; }

good() { echo -e "${color:+\e[32m}[+] $@\e[0m"; }

info() { echo -e "${color:+\e[37m}[*] $@\e[0m"; }

long_opt() {
    local arg shift="0"
    case "$1" in
        "--"*"="*) arg="${1#*=}"; [[ -n $arg ]] || usage 1 ;;
        *) shift="1"; shift; [[ $# -gt 0 ]] || usage 1; arg="$1" ;;
    esac
    echo "$arg"
    return $shift
}

subinfo() { echo -e "${color:+\e[36m}[=] $@\e[0m"; }

usage() {
    echo "Usage: ${0##*/} [OPTIONS]"
    echo
    echo "TODO"
    echo
    echo "Options:"
    echo "    -h, --help         Display this help message"
    echo "    --nocolor          Disable colorized output"
    echo "    -t, --todo         Example flag"
    echo "    -t, --todo=TODO    Example for storing cli arg"
    echo
    exit $1
}

warn() { echo -e "${color:+\e[33m}[-] $@\e[0m"; }

# Check for missing dependencies
# for dep in TODO; do
#     [[ -n $(command -v $dep) ]] || errx 3 "$dep is not installed"
# done; unset dep

declare -a args
unset help todo
color="true"

# Parse command line options
while [[ $# -gt 0 ]]; do
    case "$1" in
        "--") shift && args+=("$@") && break ;;
        "-h"|"--help") help="true" ;;
        "--nocolor") unset color ;;
        "-t"|"--todo") todo="true" ;;
        "-t"|"--todo"*) todo="$(long_opt "$@")" || shift ;;
        *) args+=("$1") ;;
    esac
    shift
done
[[ -z ${args[@]} ]] || set -- "${args[@]}"

# Check for valid params
[[ -z $help ]] || usage 0
[[ $# -eq 0 ]] || usage 2

# TODO
