package converter

import "context"

type ConverterInterface interface {
	GetCurrencyRate(ctx context.Context, conversion string) (float64, error)
}
