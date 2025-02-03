-- DDL for the database views.
-- Materialized views are created as tables.

-- Game Preview

CREATE TABLE IF NOT EXISTS "GamePreview" AS
SELECT 
    g.gameid,
    g.name AS game_name,
    genres.genres,
    platforms.platforms,
    a.url AS thumbnail_url,
    a.preview_hash AS thumbnail_hash
FROM 
    Game g
LEFT JOIN 
    (SELECT gg.gameid, GROUP_CONCAT(gen.name, ', ') AS genres
     FROM GameGenre gg
     JOIN Genre gen ON gg.genreid = gen.genreid
     GROUP BY gg.gameid) genres ON genres.gameid = g.gameid
LEFT JOIN 
    (SELECT gp.gameid, GROUP_CONCAT(plat.acronym, ', ') AS platforms
     FROM GamePlatform gp
     JOIN Platform plat ON gp.platid = plat.platid
     GROUP BY gp.gameid) platforms ON platforms.gameid = g.gameid
LEFT JOIN 
    Asset a ON g.gameid = a.gameid AND a.type = 2
GROUP BY 
    g.gameid, g.name, a.url, a.preview_hash;

-- GamePreview for highest rated games (limit 50)

CREATE TABLE IF NOT EXISTS "HighestRatedGamePreview" AS
WITH HighestRatedGames AS (
    SELECT gameid, public_rating, critics_score
    FROM Game
    WHERE public_rating IS NOT NULL
    ORDER BY public_rating DESC, critics_score DESC
    LIMIT 50
)

SELECT GP.*
FROM GamePreview GP
JOIN HighestRatedGames HRG ON GP.gameid = HRG.gameid
ORDER BY HRG.public_rating DESC, HRG.critics_score DESC;

-- GamePreview for lowest rated games (limit 50)

CREATE TABLE IF NOT EXISTS "LowestRatedGamePreview" AS
WITH LowestRatedGames AS (
    SELECT gameid, public_rating, critics_score
    FROM Game
    WHERE public_rating IS NOT NULL
    ORDER BY public_rating ASC, critics_score ASC
    LIMIT 50
)

SELECT GP.*
FROM GamePreview GP
JOIN LowestRatedGames LRG ON GP.gameid = LRG.gameid
ORDER BY LRG.public_rating ASC, LRG.critics_score ASC;

-- GameDetails

CREATE VIEW IF NOT EXISTS "GameDetails" AS
SELECT 
    g.*,
	genres,
	platforms,
	developers,
	publishers,
    screenshots,
    screenshots_hash,
	thumbnail_url,
    thumbnail_hash,
	cover_url,
	cover_hash,
    gameplay_url
FROM 
    Game g
LEFT JOIN 
    (SELECT gg.gameid, GROUP_CONCAT(gen.name, ', ') AS genres
     FROM GameGenre gg
     JOIN Genre gen ON gg.genreid = gen.genreid
     GROUP BY gg.gameid) genres ON genres.gameid = g.gameid
LEFT JOIN 
    (SELECT gpl.gameid, GROUP_CONCAT(pl.acronym, ', ') AS platforms
     FROM GamePlatform gpl
     JOIN Platform pl ON gpl.platid = pl.platid
     GROUP BY gpl.gameid) platforms ON platforms.gameid = g.gameid
LEFT JOIN 
    (SELECT gd.gameid, GROUP_CONCAT(d.name, ', ') AS developers
     FROM GameDeveloper gd
     JOIN Developer d ON gd.devid = d.devid
     GROUP BY gd.gameid) developers ON developers.gameid = g.gameid
LEFT JOIN 
    (SELECT gp.gameid, GROUP_CONCAT(p.name, ', ') AS publishers
     FROM GamePublisher gp
     JOIN Publisher p ON gp.pubid = p.pubid
     GROUP BY gp.gameid) publishers ON publishers.gameid = g.gameid
LEFT JOIN 
    (SELECT a.gameid, GROUP_CONCAT(a.url, ', ') as screenshots, GROUP_CONCAT(a.preview_hash, ', ') as screenshots_hash
     FROM Asset a
     WHERE a.type = 1
     GROUP BY a.gameid) screenshots ON screenshots.gameid = g.gameid
LEFT JOIN 
    (SELECT a.gameid, a.url as thumbnail_url, a.preview_hash as thumbnail_hash
     FROM Asset a
     WHERE a.type = 2
     GROUP BY a.gameid) thumbnail ON thumbnail.gameid = g.gameid
LEFT JOIN 
    (SELECT a.gameid, a.url as cover_url, a.preview_hash as cover_hash
     FROM Asset a
     WHERE a.type = 3
     GROUP BY a.gameid) cover ON cover.gameid = g.gameid
LEFT JOIN 
    (SELECT a.gameid, a.url as gameplay_url
     FROM Asset a
     WHERE a.type = 4
     GROUP BY a.gameid) gameplay ON gameplay.gameid = g.gameid
GROUP BY g.gameid;