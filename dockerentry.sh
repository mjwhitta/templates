#!/usr/bin/env bash

# Default group (user:1000)
gid="${DKR_GID:-1000}"
gname="${DKR_GNAME:-user}"

# Default user (user:1000)
uid="${DKR_UID:-1000}"
uname="${DKR_UNAME:-user}"

# Does user with UID already exist?
old_uid="$(id -nu "$uid" 2>/dev/null)"

# If root or UID of 0 provided, just run command
if [[ $(id -u) -ne 0 ]] || [[ $uid -eq 0 ]]; then
    exec "$@"
fi

# Create group if it doesn't exist
groupadd -f "$gname"
groupmod -g "$gid" "$gname" 2>/dev/null

# Create user if they don't exist
[[ -z $old_uid ]] || usermod -g "$gid" -l "$uname" "$old_uid"
useradd -g "$gid" -ou "$uid" "$uname" 2>/dev/null

# Sudo (a few different ways)
groupadd -f sudo
groupadd -f wheel
usermod -a -G sudo,wheel "$uname"
echo "$uname ALL=(ALL) NOPASSWD: ALL" >"/etc/sudoers.d/$uname"

# Run command as specified user
if [[ -n $(command -v sudo) ]]; then
    sudo -u "$uname" "$@"
else
    su -c "$@" "$uname"
fi
