package main

type Match struct {
	SeasonID              int
	QueueID               int
	GameID                int
	ParticipantIdentities []ParticipantIdentity
	GameVersion           string
	PlatformID            string
	GameMode              string
	MapID                 int
	GameType              string
	Teams                 []TeamStats
	Participants          []Participant
	GameDuration          int
	GameCreation          int
}

type ParticipantIdentity struct {
	Player        Player
	ParticipantID int
}

type Player struct {
	CurrentPlatformID string
	SummonerName      string
	MatchHistoryURI   string
	PlatformID        string
	CurrentAccountID  string
	ProfileIcon       int
	SummonerID        string
	AccountID         string
}

type TeamStats struct {
	FirstDragon          bool
	FirstInhibitor       bool
	Bans                 []TeamBans
	BaronKills           int
	FirstRiftHerald      bool
	FirstBaron           bool
	RiftHeraldKills      int
	FirstBlood           bool
	TeamID               int // 100 for blue side, 200 for red side
	FirstTower           bool
	VilemawKills         int
	InhibitorKills       int
	TowerKills           int
	DominionVictoryScore int
	Win                  string
	DragonKills          int
}

type TeamBans struct {
	PickTurn   int
	ChampionID int
}

type Participant struct {
	Stats                        ParticipantStats
	ParticipantID                int
	Runes                        []Rune // Only exists for legacy games, not included for matches played with Runes Reforged
	Timeline                     ParticipantTimeline
	TeamID                       int
	Spell1ID                     int
	Spell2ID                     int
	Masteries                    []Mastery // Only exists for legacy games, not included for matches played with Runes Reforged
	HighestAchievementSeasonTier string    // Highest ranked tier achieved for the previous season. Used to display in-game loading screen
	ChampionID                   int
}

type ParticipantStats struct {
	FirstBloodAssist                bool
	VisionScore                     int
	MagicDamageDealtToChampions     int
	DamageDealtToObjectives         int
	TotalTimeCrowdControlDealt      int
	LongestTimeSpentLiving          int
	Perk1Var1                       int
	Perk1Var2                       int
	Perk1Var3                       int
	TripleKills                     int
	Perk3Var3                       int
	NodeNeutralizeAssist            int
	Perk3Var2                       int
	PlayerScore9                    int
	PlayerScore8                    int
	Kills                           int
	PlayerScore1                    int
	PlayerScore0                    int
	PlayerScore3                    int
	PlayerScore2                    int
	PlayeerScore5                   int
	PlayerScore4                    int
	PlayerScore7                    int
	PlayerScore6                    int
	Perk5Var1                       int
	Perk5Var2                       int
	Perk5Var3                       int
	TotalScoreRank                  int
	NeutralMinionsKilled            int
	DamageDealtToTurrets            int
	PhysicalDamageDealtToChampions  int
	NodeCapture                     int
	LargetMultiKill                 int
	Perk2Var1                       int
	Perk2Var2                       int
	Perk2Var3                       int
	TotalUnitsHealed                int
	Perk4Var1                       int
	Perk4Var2                       int
	Perk4Var3                       int
	WardsKilled                     int
	LargestCriticalStrike           int
	LargestKillingSpree             int
	QuadraKills                     int
	TeamObjective                   int
	MagicDamageDealt                int
	Item2                           int
	Item3                           int
	Item0                           int
	NeutralMinionsKilledTeamJungle  int
	Item6                           int
	Item4                           int
	Item5                           int
	Perk1                           int // Primary path rune
	Perk0                           int // Primary path keystone rune
	Perk3                           int // Primary path rune
	Perk2                           int // Primary path rune
	Perk5                           int // Secondary path rune
	Perk4                           int // Secondary path rune
	Perk3Var1                       int
	DamageSelfMitigated             int
	MagicalDamageTaken              int
	FirstInhibitorKill              bool
	TrueDamageTaken                 int
	NodeNeutralize                  int
	Assists                         int
	CombatPlayerScore               int
	PerkPrimaryStyle                int // Primary rune path
	GoldSpent                       int
	TrueDamageDealt                 int
	ParticipantID                   int
	TotalDamageTaken                int
	PhysicalDamageDealt             int
	SightWardsBoughtInGame          int
	TotalDamageDealtToChampions     int
	PhysicalDamageTaken             int
	TotalPlayerScore                int
	Win                             bool
	ObjectivePlayerScore            int
	TotalDamageDealt                int
	Item1                           int
	NeutralMinionsKilledEnemyJungle int
	Deaths                          int
	WardsPlaced                     int
	PerkSubStyle                    int // Secondary rune path
	TurretKills                     int
	FirstBloodKill                  bool
	TrueDamageDealtToChampions      int
	GoldEarned                      int
	KillingSprees                   int
	UnrealKills                     int
	AltarsCaptured                  int
	FirstTowerAssist                bool
	FirstTowerKill                  bool
	ChampLevel                      int
	DoubleKills                     int
	NodeCaptureAssist               int
	InhibitorKills                  int
	FirstInhibitorAssist            bool
	Perk0Var1                       int
	Perk0Var2                       int
	Perk0Var3                       int
	VisionWardsBoughtInGame         int
	AltarsNeutralized               int
	PentaKills                      int
	TotalHeal                       int
	TotalMinionsKilled              int
	TimeCCingOthers                 int
}

type Rune struct {
	RuneID int
	Rank   int
}

type ParticipantTimeline struct {
	Lane                        string // Participant's calculated lane. MID and BOT are legacy values. (Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM)
	ParticipantID               int
	CsDiffPerMinDeltas          map[string]float32 // Creep score difference versus the calculated lane opponent(s) for a specified period.
	GoldPerMinDeltas            map[string]float32 // Gold for a specified period.
	XpDiffPerMinDeltas          map[string]float32 // Experience difference versus the calculated lane opponent(s) for a specified period.
	CreepsPerMinDeltas          map[string]float32 // Creeps for a specified period.
	XpPerMinDeltas              map[string]float32 // Experience change for a specified period.
	Role                        string             // Participant's calculated role. (Legal values: DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT)
	DamageTakenDiffPerMinDeltas map[string]float32 // Damage taken difference versus the calculated lane opponent(s) for a specified period.
	DamageTakenPerMinDeltas     map[string]float32 // Damage taken for a specified period.
}

type Mastery struct {
	MasteryID int
	Rank      int
}

// MatchList contains a list of MatchReferences
type MatchList struct {
	Matches    []MatchReference
	TotalGames int
	StartIndex int
	EndIndex   int
}

// MatchReference contains basic information about the game and a reference to the MatchID
type MatchReference struct {
	Lane       string
	GameID     int // MatchID
	Champion   int
	PlatformID string
	Season     int
	Queue      int
	Role       string
	Timestamp  int
}
