package main

/*
ENDPOINTS: 3
	/lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerID}
		Description: Get all champion mastery entries sorted by number of champion points descending
		Returns: []ChampionMastery
	/lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}/by-champion/{championId}
		Description: Get a champion mastery by player ID and champion ID
		Returns: ChampionMastery
	/lol/champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}
		Description: Get a player's total champion mastery score, which is the sum of individual champion mastery levels
		Returns: int
*/

// ChampionMastery contains mastery information for the provided player and champion combination
type ChampionMastery struct {
	chestGranted                 bool
	championLevel                int
	championPoints               int
	championID                   float32
	championPointsUntilNextLevel float32
	lastPlayTime                 float64
	tokensEarned                 int
	championPointsSinceLastLevel float32
	summonerID                   string
}
