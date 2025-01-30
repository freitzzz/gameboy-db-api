#!/usr/bin/env bash

root="$(dirname "$0")"

{
    find "${root}/../moby-games/pages" -mindepth 2 -maxdepth 2 -name '*.html' |
        sed 's|/[^/]*$||' |
        sed 's|/||' |
        sort |
        uniq -c |
        awk '$1 != 3 { print $2 }' |
        xargs -n 1 basename

    comm -23 <(sed 's|https://www.mobygames.com/game/[^/]\+/||; s|/$||' "${root}/game-urls.lst" | sort) \
        <(find "${root}/../moby-games/pages/" -mindepth 2 -maxdepth 2 -name '*.html' |
            sed 's|.*/||; s/_\(promo\|trivia\)\.html$//; s/\.html$//' | sort -u)
} | while read name; do
    grep "${name}" "${root}/game-urls.lst"
done
