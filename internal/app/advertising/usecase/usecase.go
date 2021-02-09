package advertisingUsecase

import (
	"github.com/Kostikans/AvitoTechadvertising/internal/app/advertising"
	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
)

type AdvertisingUsecase struct {
	AdvertisingRepo advertising.Repository
}

func NewAdvertisingUsecase(AdvertisingRepo advertising.Repository) *AdvertisingUsecase {
	return &AdvertisingUsecase{AdvertisingRepo: AdvertisingRepo}
}
func (AdvUsecase *AdvertisingUsecase) AddAdvertising(advertising advertisingModel.Advertising) (int, error) {
	return AdvUsecase.AddAdvertising(advertising)
}
func (AdvUsecase *AdvertisingUsecase) GetAdvertising(advertisingID int) (advertisingModel.Advertising, error) {
	return AdvUsecase.GetAdvertising(advertisingID)
}
func (AdvUsecase *AdvertisingUsecase) ListAdvertising(field string, desc bool) (advertisingModel.AdvertisingList, error) {
	return AdvUsecase.ListAdvertising(field, desc)
}
