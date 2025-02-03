#!/usr/bin/env bun

import fs from 'fs';
import path from 'path';
import { Database } from "bun:sqlite";

// Paths to database files
const DB_FILE = path.join(import.meta.dir, 'db.sqlite');
const SCHEMA_FILE = path.join(import.meta.dir, 'schema.sql');
const VIEWS_FILE = path.join(import.meta.dir, 'views.sql');
const SUPPORT_DML_FILE = path.join(import.meta.dir, 'support.sql');
const JSON_FILE = path.join(import.meta.dir, '../scraping/merge.json');

// Delete the database if it already exists
if (fs.existsSync(DB_FILE)) {
    fs.rmSync(DB_FILE);
    console.log("Existing database deleted.");
}

// Read schema, views, support dml, and JSON data
const schemaDDL = fs.readFileSync(SCHEMA_FILE, 'utf-8');
const viewsDDL = fs.readFileSync(VIEWS_FILE, 'utf-8');
const enumDML = fs.readFileSync(SUPPORT_DML_FILE, 'utf-8');
const data = JSON.parse(fs.readFileSync(JSON_FILE, 'utf-8'));

// Initialize SQLite database
const db = new Database(DB_FILE);

const esrbEnum = {
    'Rating Pending': 1,
    'Early Childhood': 2,
    'Everyone': 3,
    'Kids to Adults': 4,
    'Everyone 10+': 5,
    'Teen': 6,
    'Mature': 7,
    'Adult': 8,
};

const platformEnum = {
    'GB': 1,
    'GBC': 2,
    'GBA': 3,
}

// Initialize schema
function initializeSchema() {
    db.transaction(() => {
        schemaDDL.split(';').forEach(stmt => {
            if (stmt.trim()) db.exec(stmt.trim());
        });
    })();
    console.log("Schema initialized successfully.");
}

// Initialize views
function initializeViews() {
    db.transaction(() => {
        viewsDDL.split(';').forEach(stmt => {
            if (stmt.trim()) db.exec(stmt.trim());
        });
    })();
    console.log("Views initialized successfully.");
}

// Populate enum tables using a cursor
function populateEnumTables() {
    db.transaction(() => {
        enumDML.split(';').forEach(stmt => {
            if (stmt.trim()) db.exec(stmt.trim());
        });
    })();
    console.log("Enum tables populated successfully.");
}

function getId(table, idColumn, valueColumn, value) {
    const stmt = db.prepare(`SELECT ${idColumn} FROM ${table} WHERE ${valueColumn} = ? LIMIT 1`);
    const result = stmt.get(value);
    return result ? Object.values(result)[0] : null;
}

// Helper function to fetch rowid for a value in a table
function getRowId(table, column, value) {
    return getId(table, 'rowid', column, value);
}

// Insert data using cursors
function populateData() {
    const insertGameStmt = db.prepare(
        `INSERT INTO Game (name, description, release_year, esrb_rating, trivia, promo, adult, public_rating, critics_score)
     VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`
    );
    const insertAssetStmt = db.prepare(
        `INSERT INTO Asset (gameid, url, type, preview_hash) VALUES (?, ?, ?, ?);`
    );
    const insertDeveloperStmt = db.prepare(`INSERT INTO Developer (name) VALUES (?);`);
    const insertPublisherStmt = db.prepare(`INSERT INTO Publisher (name) VALUES (?);`);
    const insertGenreStmt = db.prepare(`INSERT INTO Genre (name) VALUES (?);`);
    const insertPlatformStmt = db.prepare(`INSERT INTO Platform (platid, acronym) VALUES (?, ?);`);
    const insertSourceStmt = db.prepare(
        `INSERT INTO Source (name, gameid, url) VALUES (?, ?, ?);`
    );
    const insertGameGenreStmt = db.prepare(
        `INSERT INTO GameGenre (genreid, gameid) VALUES (?, ?);`
    );
    const insertGamePlatformStmt = db.prepare(
        `INSERT INTO GamePlatform (platid, gameid) VALUES (?, ?);`
    );
    const insertGamePublisherStmt = db.prepare(
        `INSERT INTO GamePublisher (pubid, gameid) VALUES (?, ?);`
    );
    const insertGameDeveloperStmt = db.prepare(
        `INSERT INTO GameDeveloper (devid, gameid) VALUES (?, ?);`
    );

    const platforms = new Set();
    const genres = new Set();
    const publishers = new Set();
    const developers = new Set();

    // Step 1: Insert developers, publishers, genres, and platforms (needed for many-to-many associations)
    db.transaction(() => {
        data.forEach(game => {
            game.developers?.forEach(dev => developers.add(dev));
            game.publishers?.forEach(pub => publishers.add(pub));
            game.genres?.forEach(genre => genres.add(genre));
            game.platforms?.forEach(platform => platforms.add(platform));
        });

        developers.forEach(dev => insertDeveloperStmt.run(dev));
        publishers.forEach(pub => insertPublisherStmt.run(pub));
        genres.forEach(genre => insertGenreStmt.run(genre));
        platforms.forEach(platform => insertPlatformStmt.run(platformEnum[platform], platform));
    })();

    console.log("Developer, Publisher, Genre, and Platform tables populated.");

    // Step 2: Insert games and their associated assets
    db.transaction(() => {
        data.forEach(game => {
            // Insert game and get the row ID
            const result = insertGameStmt.run(
                game.name,
                game.description ?? null,
                game.release_year,
                esrbEnum[game.esrb_rating] ?? 0,
                game.trivia ?? null,
                game.promo ?? null,
                game.adult ? 1 : 0,
                game.public_rating,
                game.critics_score
            );
            const gameId = result.lastInsertRowid;

            // Insert associated assets for the game
            if (game.thumbnail_url) insertAssetStmt.run(gameId, game.thumbnail_url, 2, null);
            if (game.cover_url) insertAssetStmt.run(gameId, game.cover_url, 3, null);
            if (game.gameplay_url) insertAssetStmt.run(gameId, game.gameplay_url, 4, null);
            if (Array.isArray(game.screenshots)) {
                game.screenshots.forEach(url => insertAssetStmt.run(gameId, url, 1, null));
            }

            // Insert source URL if available
            if (game.source_url) {
                insertSourceStmt.run(0, gameId, game.source_url); // 0 as MobyGames is the only source
            }

            // Insert game associations with correct rowid
            game.genres?.forEach(genre => {
                const genreId = getRowId('Genre', 'name', genre);
                if (genreId) insertGameGenreStmt.run(genreId, gameId);
            });

            game.platforms?.forEach(platform => {
                const platformId = getId('Platform', 'platid', 'acronym', platform);
                if (platformId) insertGamePlatformStmt.run(platformId, gameId);
            });

            game.publishers?.forEach(pub => {
                const pubId = getRowId('Publisher', 'name', pub);
                if (pubId) insertGamePublisherStmt.run(pubId, gameId);
            });

            game.developers?.forEach(dev => {
                const devId = getRowId('Developer', 'name', dev);
                if (devId) insertGameDeveloperStmt.run(devId, gameId);
            });
        });
    })();

    // Finalize the prepared statements
    insertGameStmt.finalize();
    insertAssetStmt.finalize();
    insertDeveloperStmt.finalize();
    insertPublisherStmt.finalize();
    insertGenreStmt.finalize();
    insertPlatformStmt.finalize();
    insertSourceStmt.finalize();
    insertGameGenreStmt.finalize();
    insertGamePlatformStmt.finalize();
    insertGamePublisherStmt.finalize();
    insertGameDeveloperStmt.finalize();

    console.log("Data populated successfully.");
}

// Execute schema, populate enums, and insert data
initializeSchema();
populateEnumTables();
populateData();
initializeViews();

// Close the database
db.close();
console.log("Database closed.");
