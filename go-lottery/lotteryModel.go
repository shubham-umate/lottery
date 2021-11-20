package main

type Participant struct {
	Id               uint   `json:"id"`
	LotteryId        uint   `json:"lotteryId"`
	ParticipantEmail string `json:"participantEmail" gorm:"unique"`
}

type Lottery struct {
	Id           uint   `json:"id"`
	LotteryId    uint   `json:"lotteryId"`
	LotteryName  string `json:"lotteryName"`
	Limit        string `json:"limit"`
	Participants int    `json:"participants"`
}
