#!/usr/bin/env bash

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
    echo "    -s, --store=VAL    Example for storing cli arg"
    echo
    exit $1
}

declare -a args
unset flag help store

while [[ $# -gt 0 ]]; do
    case "$1" in
        "--") shift && args+=("$@") && break ;;
        "-f"|"--flag") flag="true" ;;
        "-h"|"--help") help="true" ;;
        "-s"|"--store"*) store="$(long_opt $@)" || shift ;;
        *) args+=("$1") ;;
    esac
    shift
done
[[ -z ${args[@]} ]] || set -- "${args[@]}"

[[ -z $help ]] || usage 0
[[ $# -eq 0 ]] || usage 2
