package entity

import "time"

type UserToken struct {
	Id                 int       `json:"id"`
	UserId             int       `json:"userId"`
	Token              string    `json:"token"`
	ExpireAt           time.Time `json:"expireAt"`
	IsDeleted          bool      `json:"isDeleted"`
	ResendTime         time.Time `json:"resendTime"`
	RecordCreateUserId string    `json:"recordCreateUserId"`
	RecordCreateTime   time.Time `json:"recordCreateTime"`
	RecordUpdateUserId string    `json:"recordUpdateUserId"`
	RecordUpdateTime   time.Time `json:"recordUpdateTime"`
}
