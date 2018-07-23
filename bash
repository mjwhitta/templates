#!/usr/bin/env bash

err() { echo -e "${clr:+\e[31m}[!] $@\e[0m"; }

errx() { echo -e "${clr:+\e[31m}[!] ${@:2}\e[0m"; exit $1; }

info() { echo -e "${clr:+\e[37m}[*] $@\e[0m"; }

long_opt() {
    local arg shift="0"
    case "$1" in
        "--"*"="*) arg="${1#*=}"; [[ -n $arg ]] || usage 1 ;;
        *) shift="1"; shift; [[ $# -gt 0 ]] || usage 1; arg="$1" ;;
    esac
    echo "$arg"
    return $shift
}

usage() {
    echo "Usage: ${0/*\//} [OPTIONS]"
    echo
    echo "TODO"
    echo
    echo "Options:"
    echo "    -f, --flag         Example flag"
    echo "    -h, --help         Display this help message"
    echo "    --nocolor          Disable colorized output"
    echo "    -s, --store=VAL    Example for storing cli arg"
    echo
    exit $1
}

warn() { echo -e "${clr:+\e[33m}[-] $@\e[0m"; }

declare -a args
unset flag help store
clr="true"

while [[ $# -gt 0 ]]; do
    case "$1" in
        "--") shift && args+=("$@") && break ;;
        "-f"|"--flag") flag="true" ;;
        "-h"|"--help") help="true" ;;
        "--nocolor") unset clr ;;
        "-s"|"--store"*) store="$(long_opt $@)" || shift ;;
        *) args+=("$1") ;;
    esac
    shift
done
[[ -z ${args[@]} ]] || set -- "${args[@]}"

[[ -z $help ]] || usage 0
[[ $# -eq 0 ]] || usage 2
