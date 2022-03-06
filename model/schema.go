package model

// For testing

var schemaTables = `
DROP TABLE IF EXISTS steam_user;
DROP TABLE IF EXISTS player_stats;
DROP TABLE IF EXISTS map;

CREATE TABLE IF NOT EXISTS "steam_user" (
	"id" SERIAL PRIMARY KEY,
	"steamid" VARCHAR(32),
	"personaname" TEXT,
	"lastlogoff" INTEGER,
	"profileurl" TEXT,
	"avatar" TEXT,
	"avatarmedium" TEXT,
	"avatarfull" TEXT,
	"realname" TEXT,
	"primaryclanid" TEXT,
	"timecreated" INTEGER,
	"loccountrycode" TEXT,
	"created_at" TIMESTAMP,
	"updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "player_stats" (
	"id" SERIAL PRIMARY KEY,
	"player_steamid" VARCHAR(32),
	"map_id" INT,
	"kills" INT DEFAULT 0,
	"deaths" INT DEFAULT 0,
	"assists" INT DEFAULT 0,
	"shots" INT DEFAULT 0,
	"hits" INT DEFAULT 0,
	"damage" INT DEFAULT 0,
	"first_blood" INT DEFAULT 0,
	"aces" INT DEFAULT 0,
	"headshots" INT DEFAULT 0,
	"no_scope" INT DEFAULT 0,
	"count" INT DEFAULT 0,
	"playtime" INT DEFAULT 0,
	"match_win" INT DEFAULT 0,
	"match_lose" INT DEFAULT 0,
	"match_draw" INT DEFAULT 0,
	"rounds_won" INT DEFAULT 0,
	"rounds_lost" INT DEFAULT 0,
	"mvp" INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "map" (
	"id" SERIAL PRIMARY KEY, 
	"name" VARCHAR(64)
);
`

var schemaFk = `
ALTER TABLE "player_stats"
    ADD CONSTRAINT fk_map
    FOREIGN KEY (map_id) 
    REFERENCES map (id);
`

var schemaTestData = `

INSERT INTO steam_user(steamid, personaname, lastlogoff, profileurl, avatar, avatarmedium, 
	avatarfull, realname, primaryclanid, timecreated, loccountrycode, created_at) 
	VALUES ('steamid','bozo',123,'kkk','a', 'b', 'c', 'yes', '13', 123, 'dd', NOW());

INSERT INTO map(name) VALUES ('xdream');
INSERT INTO map(name) VALUES ('ski');
INSERT INTO map(name) VALUES ('evil');

INSERT INTO player_stats(player_steamid, map_id) VALUES ('1337steamidyes',1);

INSERT INTO player_stats(player_steamid, map_id, kills, deaths, assists, shots, 
	hits, damage, first_blood, aces, headshots, no_scope, count, playtime, match_win) 
	VALUES ('76561198226912040', 2, 7, 2, 7, 7, 7, 1, 7, 5, 7, 7, 2, 7, 7);

INSERT INTO player_stats(player_steamid, map_id, kills, deaths, assists, shots, 
		hits, damage, first_blood, aces, headshots, no_scope, count, playtime, match_win) 
		VALUES ('76561198226912040', 3, 1, 2, 3, 6, 17, 1, 72, 5, 17, 7, 12, 0, 7);

`

// TODO: add fkey (steamid -> player_steamid ?, ..)
