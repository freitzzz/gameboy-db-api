#!/usr/bin/env bash

root="$(dirname "$0")"

cat "${root}/../moby-games/browser.json" | jq -r '.[].data.games[].internal_url' | sort -u >${root}/game-urls.lst
