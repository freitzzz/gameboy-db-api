-- DML for the database support tables.
-- All instructions have been written to target SQLite engine.
-- --

-- Source names

INSERT INTO "SourceName" VALUES (1, "MOBYGAMES");

-- Asset types

INSERT INTO "AssetType" VALUES (1, "SCREENSHOT");
INSERT INTO "AssetType" VALUES (2, "THUMBNAIL");
INSERT INTO "AssetType" VALUES (3, "COVER");
INSERT INTO "AssetType" VALUES (4, "GAMEPLAY");

-- ESRB Rating

INSERT INTO "ESRB" VALUES (1, "RATING_PENDING");
INSERT INTO "ESRB" VALUES (2, "EARLY_CHILDHOOD");
INSERT INTO "ESRB" VALUES (3, "EVERYONE");
INSERT INTO "ESRB" VALUES (4, "EVERYONE_ABOVE_10");
INSERT INTO "ESRB" VALUES (5, "TEEN");
INSERT INTO "ESRB" VALUES (6, "MATURE");
INSERT INTO "ESRB" VALUES (7, "ADULT");