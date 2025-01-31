-- DDL for the database schemas.
-- All instructions have been written to target SQLite engine.
-- --
-- Does not contain any database data.


-- Primary tables

CREATE TABLE IF NOT EXISTS "Genre" (
    "genreid" INTEGER PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "Platform" (
    "platid" INTEGER PRIMARY KEY,
    "acronym" VARCHAR(3) CHECK ( "acronym" in ("GB", "GBC", "GBA") ) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "Publisher" (
    "pubid" INTEGER PRIMARY KEY,
    "name" VARCHAR(128) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "Developer" (
    "devid" INTEGER PRIMARY KEY,
    "name" VARCHAR(128) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "Source" (
    "name" INTEGER(1) REFERENCES "SourceName" ("id")  NOT NULL,
    "gameid" INTEGER REFERENCES "Game" ("gameid") NOT NULL,
    "url" VARCHAR NOT NULL,
    PRIMARY KEY ("name", "gameid")
);

CREATE TABLE IF NOT EXISTS "Asset" (
    "assetid" INTEGER PRIMARY KEY,
    "gameid" INTEGER REFERENCES "Game" ("gameid") NOT NULL,
    "url" VARCHAR NOT NULL,
    "type" INTEGER REFERENCES "AssetType" ("id")  NOT NULL,
    "preview_hash" VARCHAR(40) NULL
);

CREATE TABLE IF NOT EXISTS "Game" (
    "gameid" INTEGER PRIMARY KEY,
    "name" VARCHAR(128) NOT NULL,
    "description" VARCHAR NULL,
    -- TODO: Should include a check for year >= 1984! Currently disabled because the scraped data references the first origin of the game, for any platform.
    "release_year" INTEGER NOT NULL,
    "esrb_rating" INTEGER REFERENCES "ESRB" ("id")  NOT NULL,
    "trivia" VARCHAR NULL,
    "promo" VARCHAR NULL,
    "adult" INTEGER CHECK ( "adult" in (0, 1) ) NOT NULL,
    "public_rating" INTEGER CHECK ( "public_rating" >= 0 AND "public_rating" <= 100 ) NULL,
    "critics_score" INTEGER CHECK ( "critics_score" >= 0 AND "critics_score" <= 100 ) NULL
);


-- Association tables

CREATE TABLE IF NOT EXISTS "GameGenre" (
    "genreid" INTEGER REFERENCES "Genre" ("genreid") NOT NULL,
    "gameid" INTEGER REFERENCES "Game" ("gameid") NOT NULL,
    PRIMARY KEY ("genreid", "gameid")
);

CREATE TABLE IF NOT EXISTS "GamePlatform" (
    "platid" INTEGER REFERENCES "Platform" ("platid") NOT NULL,
    "gameid" INTEGER REFERENCES "Game" ("gameid") NOT NULL,
    PRIMARY KEY ("platid", "gameid")
);

CREATE TABLE IF NOT EXISTS "GamePublisher" (
    "pubid" INTEGER REFERENCES "Publisher" ("pubid") NOT NULL,
    "gameid" INTEGER REFERENCES "Game" ("gameid") NOT NULL,
    PRIMARY KEY ("pubid", "gameid")
);

CREATE TABLE IF NOT EXISTS "GameDeveloper" (
    "devid" INTEGER REFERENCES "Developer" ("devid") NOT NULL,
    "gameid" INTEGER REFERENCES "Game" ("gameid") NOT NULL,
    PRIMARY KEY ("devid", "gameid")
);


-- Support tables

CREATE TABLE IF NOT EXISTS "SourceName" (
    "id" INTEGER(1) PRIMARY KEY CHECK ( "id" in (0) ) NOT NULL,
    "value" VARCHAR CHECK ( "value" in ("MOBYGAMES") ) NOT NULL UNIQUE
) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS "AssetType" (
    "id" INTEGER(1) PRIMARY KEY CHECK ( "id" in (0, 1, 2, 3) ) NOT NULL,
    "value" VARCHAR CHECK ( "value" in ("SCREENSHOT", "THUMBNAIL", "COVER", "GAMEPLAY") ) NOT NULL UNIQUE
) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS "ESRB" (
    "id" INTEGER(1) PRIMARY KEY CHECK ( "id" in (0, 1, 2, 3, 4, 5, 6) ) NOT NULL,
    "value" VARCHAR CHECK ( "value" in ("RATING_PENDING", "EARLY_CHILDHOOD", "EVERYONE", "EVERYONE_ABOVE_10", "TEEN", "MATURE", "ADULT") ) NOT NULL UNIQUE
) WITHOUT ROWID;