package main

/*
ENDPOINTS: 1
	/lol/status/v3/shard-data
		Description: Get League of Legends status for the given shard
		Returns: ShardStatus
*/

// ShardStatus contains information about the status of a League region
type ShardStatus struct {
	name      string
	regionTag string
	hostname  string
	services  []Service
	slug      string
	locales   []string
}

// Service contains the status of various League services
type Service struct {
	status    string
	incidents []Incident
	name      string
	slug      string
}

// Incident contains information about an incident impacting the status of a League Service
type Incident struct {
	active    bool
	createdAt string
	ID        float64
	updates   []Message
}

// Message contains information given about an Incident
type Message struct {
	severity     string
	author       string
	createdAt    string
	translations []Translation
	updatedAt    string
	content      string
	ID           string
}

// Translation contains translations for Message in different locales
type Translation struct {
	locale    string
	content   string
	updatedAt string
}
