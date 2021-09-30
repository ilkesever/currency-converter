package models

type ConversionRequestDto struct {
	TargetCurrency string              `json:"conversion"`
	Data           []ValueWithCurrency `json:"data"`
}

type ValueWithCurrency struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
