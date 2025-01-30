#!/usr/bin/env bash

root="$(dirname "$0")"
worker_url="..."

cat "${root}/game-urls.lst" | while read url; do
    name=$(basename $url)
    game_id=$(awk -F '/' '{print $(NF-2)}' <<<$url)

    echo "Processing ${game_id}/${name}..."
    curl -X POST "${worker_url}" -H "x-download-url: ${url}" -H "x-file-name: pages/${game_id}/${name}.html"
    curl -X POST "${worker_url}" -H "x-download-url: ${url}adblurbs/" -H "x-file-name: pages/${game_id}/${name}_promo.html"
    curl -X POST "${worker_url}" -H "x-download-url: ${url}trivia/" -H "x-file-name: pages/${game_id}/${name}_trivia.html"
done
