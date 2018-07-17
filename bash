#!/usr/bin/env bash

usage() {
    echo "Usage: ${0/*\//} [OPTIONS]"
    echo
    echo "TODO"
    echo
    echo "Options:"
    echo "    -f, --flag         Example flag"
    echo "    -f, --flag=FLAG    Example for storing cli arg"
    echo "    -h, --help         Display this help message"
    echo
    exit $1
}

declare -a args
unset flag help

while [[ $# -gt 0 ]]; do
    case "$1" in
        "--") shift && args+=("$@") && break ;;
        "-f"|"--flag") flag="true" ;;
        "-f"|"--flag"*)
            case "$1" in
                "--"*"="*) arg="${1#*=}"; [[ -n $arg ]] || usage 1 ;;
                *) shift; [[ $# -gt 0 ]] || usage 1; arg="$1" ;;
            esac
            flag="--flag $arg"
            ;;
        "-h"|"--help") help="true" ;;
        *) args+=("$1") ;;
    esac
    shift
done
[[ -z ${args[@]} ]] || set -- "${args[@]}"

[[ -z $help ]] || usage 0
[[ $# -eq 0 ]] || usage 2
