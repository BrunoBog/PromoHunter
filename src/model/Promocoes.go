package model

type Promocoes struct {
	Promobit PromobitPromo `json:"Promobit"  bson:"PromobitPromo"`
	HardMob  HardMobPromo  `json:"HardMob"  bson:"HardMobPromo"`
}
