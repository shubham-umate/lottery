package main

type Participant struct {
	Id               uint   `json:"id"`
	LotteryId        uint   `json:"lotteryId"`
	ParticipantEmail string `json:"participantEmail"`
}

type Lottery struct {
	Id           uint   `json:"id"`
	LotteryName  string `json:"lotteryName"`
	Limit        int    `json:"limit"`
	Participants int    `json:"participants"`
}
