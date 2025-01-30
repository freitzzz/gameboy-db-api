#!/usr/bin/env bun

import fs from 'fs';
import path from 'path';
import { parse } from 'node-html-parser';  // Bun's built-in HTML parser
import { htmlToText } from 'html-to-text';

// Input and output paths
const inputFolder = process.argv[2];
const outputFile = process.argv[3];

// Convert HTML content to raw text using html-to-text library
function convertToRawText(htmlContent) {
  return htmlToText(htmlContent, {
    wordwrap: 130,
    ignoreHref: true,  // Don't convert links
    ignoreImage: true, // Don't include image tags
  }).replace(/\[(http|https):\/\/.*\]/g, '');
}

// Extract content from promo.html (inside the first <blockquote>)
function extractPromoContent(root) {
  const blockquote = root.querySelector('blockquote');
  return blockquote ? convertToRawText(blockquote.innerHTML) : null;
}

// Extract content from trivia.html (inside the element with id="gameTrivia")
function extractTriviaContent(root) {
  const trivia = root.querySelector('#gameTrivia');
  return trivia ? convertToRawText(trivia.innerHTML) : null;
}

// Extract fields from the standard game page
function extractGamePageContent(root) {
  const result = {};

  // Description
  const description = root.querySelector('#description-text');
  if (description) result.description = convertToRawText(description.innerHTML);

  // Screenshots (Extract the src as plain text for raw format)
  const screenshots = root.querySelectorAll('#gameShots a img');
  result.screenshots = screenshots.map((img) => img.getAttribute('src'));

  // Cover URL
  const coverMeta = root.querySelector('meta[property="og:image"]');
  result.cover_url = coverMeta ? coverMeta.getAttribute('content') : null;

  // ESRB Rating
  const esrbDt = root.querySelector('dt:contains("ESRB Rating")');
  if (esrbDt) {
    const esrbDd = esrbDt.nextElementSibling;
    result.esrb_rating = esrbDd ? esrbDd.textContent.trim() : null;
  }

  // Video URL
  const videosSection = root.querySelector('#gameVideos');
  result.gameplay_url = videosSection ? videosSection.querySelector("div[data-src]").getAttribute('data-src') : null;


  return result;
}

// Process all HTML files for a specific game
function processGameFolder(gameId, folderPath) {
  const gameData = { game_id: Number.parseInt(gameId) };

  const files = fs.readdirSync(folderPath);
  files.forEach((filename) => {
    const filePath = path.join(folderPath, filename);
    if (filename.endsWith('.html')) {
      const content = fs.readFileSync(filePath, 'utf-8');
      const root = parse(content); // Using Bun's built-in HTML parser

      if (filename.endsWith('_promo.html')) {
        gameData.promo = extractPromoContent(root);
      } else if (filename.endsWith('_trivia.html')) {
        gameData.trivia = extractTriviaContent(root);
      } else {
        Object.assign(gameData, extractGamePageContent(root));
      }
    }
  });

  return gameData;
}

// Main function to process all games and save to output
async function main() {
  const allGames = [];

  const folders = fs.readdirSync(inputFolder);
  for (const folderName of folders) {
    const folderPath = path.join(inputFolder, folderName);
    if (fs.statSync(folderPath).isDirectory()) {
      const gameId = folderName;
      const gameData = processGameFolder(gameId, folderPath);
      allGames.push(gameData);
    }
  }

  fs.writeFileSync(outputFile, JSON.stringify(allGames, null, 2));
  console.log(`HTML extraction and transformation complete. Data saved to ${outputFile}`);
}

// Run the main function
main();
