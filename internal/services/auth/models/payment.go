package models

import "time"

type Payment struct {
	Id        int64
	UserId    int64
	PaymentId string
	Amount    string
	Status    string
	Paid      bool
	Type      string
	Credited  bool
	Time      time.Time
}

type Payments []Payment
