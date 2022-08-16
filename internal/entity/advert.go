package entity

import (
	"time"
	"unicode/utf8"
)

const (
	maxLenUrls = 3
	maxLenName = 200
	maxLenDesc = 1000
)

func NewAdvert() *Advert {
	return &Advert{
		Name:        "Название для объявления отсутствует",
		Description: "Описание для объявления отсутствует",
	}
}

func NewAdvertResponse(adv Advert) AdvertResponse {
	return AdvertResponse{
		Name:  adv.Name,
		Price: adv.Price,
		Urls:  adv.Urls,
	}
}

type Advert struct {
	Id          int64     `json:"id"`
	CreatedDt   time.Time `json:"created_dt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Urls        []string  `json:"urls"`
}

type AdvertResponse struct {
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Urls  []string `json:"urls"`
}

func (adv *Advert) Check() (err error) {
	if len(adv.Urls) > maxLenUrls == true {
		err = ErrorTooManyUrls
	} else if utf8.RuneCountInString(adv.Name) > maxLenName {
		err = ErrorTooManyNameSymbols
	} else if utf8.RuneCountInString(adv.Description) > maxLenDesc {
		err = ErrorTooManyDescriptionSymbols
	} else if adv.Price < 0 {
		err = ErrorNegativePrice
	}
	return
}
