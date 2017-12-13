#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )"

pm25_interpolation="\#{pm25}"
if [[ $OSTYPE == "darwin"* ]]; then
  pm25="#($CURRENT_DIR/bin/pm25_mac)"
elif [[ $OSTYPE == "linux"* ]]; then
  pm25="#($CURRENT_DIR/bin/pm25_linux)"
else
  echo "Platform not supported"
  exit 1
fi

get_tmux_option () {
  local option="$1"
  local default="$2"
  local value="$(tmux show-option -gqv "$option")"
  if [ -z "$value" ]; then
    echo "$default"
  else
    echo "$value"
  fi
}

set_tmux_option () {
  local option=$1
  local value=$2
  tmux set-option -gq "$option" "$value"
}

do_interpolation () {
  local all_interpolated="$1"
  all_interpolated=${all_interpolated/$pm25_interpolation/$pm25}
  echo "$all_interpolated"
}

update_tmux_option () {
  local option="$1"
  local value=$(get_tmux_option "$option")
  local new_value=$(do_interpolation "$value")
  set_tmux_option "$option" "$new_value"
}

main () {
  update_tmux_option "status-right"
  update_tmux_option "status-left"
}
main
