#!/usr/bin/env bash

root="$(dirname "$0")"

browser_json="${root}/../moby-games/browser.json"
scraped_pages_dir="${root}/../moby-games/pages"

browser_data_json="${root}/browser-data.json"
scraped_data_json="${root}/pages-data.json"
merge_json="${root}/merge.json"

jq '[.[] | .data.games[] | {
  game_id: .game_id,
  name: .title,
  release_year: (if .release_date then (.release_date | split("-")[0] | tonumber) else null end),
  adult: .adult,
  genres: [.genres[].name],
  platforms: [.platforms[] | select(.name | startswith("Game Boy")) | (.name | split(" ") | map(.[0:1]) | join(""))],
  developers: [.companies[] | select(.title_id == 2) | .name],
  publishers: [.companies[] | select(.title_id == 1) | .name],
  public_rating: (if .moby_score then (.moby_score * 10 | ceil) else null end),
  critics_score: (.critic_score // null),
  thumbnail_url: .cover.tiny_url,
  source_url: .internal_url
}] | unique_by(.game_id)' "$browser_json" >"$browser_data_json"

bun run "${root}/merge-scraped-pages.js" "$scraped_pages_dir" "$scraped_data_json"

bun --eval "import data from '${browser_data_json}'; import data2 from '${scraped_data_json}'; const data3 = data.map((x) => Object.assign(x, data2.find((y) => y.game_id === x.game_id))); console.log(JSON.stringify(data3, null, 2))" >"${merge_json}"

rm "$browser_data_json" "$scraped_data_json"
