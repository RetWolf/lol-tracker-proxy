package main

/*
ENDPOINTS: 1
	/lol/platform/v3/champion-rotations
		Description: Returns champion rotations, including free-to-play and low-level free-to-play rotations
		Returns: ChampionInfo
*/

// ChampionInfo contains free champion IDs and the max new player level
type ChampionInfo struct {
	freeChampionIDsForNewPlayers []int
	freeChampionIDs              []int
	maxNewPlayerLevel            int
}
