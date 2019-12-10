package main

type MatchDto struct {
	SeasonID              int
	QueueID               int
	GameID                float64
	ParticipantIdentities []ParticipantIdentityDto
	GameVersion           string
	PlatformID            string
	GameMode              string
}

type ParticipantIdentityDto struct {
}
