#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )"

pm25_interpolation="\#{pm25}"
pm25="#($CURRENT_DIR/pm25)"

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

check_command () {
  if hash go 2>/dev/null; then
    if [ ! -f $CURRENT_DIR/pm25 ]; then
      go build $CURRENT_DIR/pm25.go
    fi
  else
    echo "Go not found"
    exit 1
  fi
}

main () {
  check_command
  update_tmux_option "status-right"
  update_tmux_option "status-left"
}
main
#tmux set -g status-right "#(get_rand) $(get_opt)"
