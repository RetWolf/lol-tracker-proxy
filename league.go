package main

/*
ENDPOINTS: 6
	/lol/league/v4/challengerleagues/by-queue/{queue}
		Description: Get the challenger league for given queue
		Returns: LeagueList
	/lol/league/v4/entries/by-summoner/{encryptedSummonerId}
		Description: Get league entries in all queues for a given summoner ID
		Returns: []LeagueEntry
	/lol/league/v4/entries/{queue}/{tier}/{division}
		Description: Get all the league entries
		Returns: []LeagueEntry
	/lol/league/v4/grandmasterleagues/by-queue/{queue}
		Description: Get the grandmaster league of a specific queue
		Returns: LeagueList
	/lol/league/v4/leagues/{leagueId}
		Description: Get league with given ID, including inactive entries
		Returns: LeagueList
	/lol/league/v4/masterleagues/by-queue/{queue}
		Description: Get the master league for given queue
		Returns: LeagueList
*/

// LeagueList contains a ladder listing with player entries
type LeagueList struct {
	leagueID string
	tier     string
	entries  []LeagueItem
	queue    string
	name     string
}

// LeagueItem contains information about a player on a given Queue/Tier's ladder
type LeagueItem struct {
	summonerName string
	hotStreak    bool
	miniSeries   MiniSeries
	wins         int
	veteran      bool
	losses       int
	freshBlood   bool
	inactive     bool
	rank         string
	summonerID   string
	leaguePoints int
}

// LeagueEntry contains information about a player on a given Queue/Tier's ladder
type LeagueEntry struct {
	queueType    string
	summonerName string
	hotStreak    bool
	miniSeries   MiniSeries
	wins         int
	veteran      bool
	losses       int
	rank         string
	leagueID     string
	inactive     bool
	freshBlood   bool
	tier         string
	summonerID   string
	leaguePoints int
}

// MiniSeries contains stats on the current promotional series the player is in, if he/she is in one
type MiniSeries struct {
	progress string
	losses   int
	target   int
	wins     int
}
