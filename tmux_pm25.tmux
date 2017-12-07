#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )"

pm25_interpolation="\#{test_int}"
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
  all_interpolated=${all_interpolated/$test_interpolation/$test_value}
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
#tmux set -g status-right "#(get_rand) $(get_opt)"
