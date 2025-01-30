# scraping

The database is bootstrapped with [MobyGames](https://www.mobygames.com/) data. I will not be providing the scrape sources, but all the scripts used can be found in this folder.

1. Run `gen-urls-list.sh` (requires manual scrape of game metadata via MobyGames Browser API)
2. Run `run-scrape.sh` (requires setting up Cloudflare Worker to download pages)
    1. Run `verify-scrape.sh` to verify that no game pages are left to scrape
3. Run `merge-scrape-sh`