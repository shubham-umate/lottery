package main

type Participant struct {
	Id               uint   `json:"id"`
	ParticipantEmail string `json:"participantEmail"`
	LotteryId        uint   `json:"lotteryId"`
}

type Lottery struct {
	Id           uint   `json:"id"`
	LotteryName  string `json:"lotteryName"`
	Limit        int    `json:"limit"`
	Participants int    `json:"participants"`
	Winner       uint   `json:"-"`
}

type Winner struct {
	Id          uint   `json:"id"`
	LotteryId   uint   `json:"lotteryId"`
	WinnerEmail string `json:"winnerEmail"`
}
