package advertising

import advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

type Usecase interface {
	AddAdvertising(advertising advertisingModel.Advertising) int
	GetAdvertising(advertisingID int) advertisingModel.Advertising
	ListAdvertising(field string, desc bool) advertisingModel.AdvertisingList
}
