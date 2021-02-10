package advertisingRepository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/customError"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/responseCodes"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AdvertisingRepository struct {
	db *sqlx.DB
}

func NewAdvertisingRepository(db *sqlx.DB) *AdvertisingRepository {
	return &AdvertisingRepository{db: db}
}
func (advRepo *AdvertisingRepository) AddAdvertising(advertising advertisingModel.AdvertisingAdd) (advertisingModel.AdvertisingID, error) {
	var advertisingID advertisingModel.AdvertisingID
	err := advRepo.db.QueryRow(AddAdvertising, advertising.Name, advertising.Description, advertising.Photos[0], pq.Array(advertising.Photos[1:]), advertising.Cost).Scan(&advertisingID.AdvID)
	if err != nil {
		return advertisingID, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	return advertisingID, nil
}
func (advRepo *AdvertisingRepository) GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error) {
	var advertising advertisingModel.Advertising
	err := advRepo.db.QueryRow(GetAdvertising, advertisingID).Scan(&advertising.Name, &advertising.Cost, &advertising.MainPhoto)
	if err != nil {
		return advertising, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	return advertising, nil
}

func (advRepo *AdvertisingRepository) GenerateQueryForGetAdvertising(fields string) string {
	query := "SELECT name,cost,photos[1]"
	fieldsStrings := strings.Split(fields, ",")
	if fieldsStrings[0] == "description" {
		query += ",description"
	}
	if fieldsStrings[1] == "photos" {
		query += ",photos"
	}

	query += " from advertising where advertising_id=$1"
	return query

}

func (advRepo *AdvertisingRepository) ListAdvertising(sort string, desc string) (advertisingModel.AdvertisingList, error) {
	var advertisingList advertisingModel.AdvertisingList
	var advertising []advertisingModel.Advertising

	err := advRepo.db.Select(&advertising, advRepo.GenerateQueryForGetAdvertisingList(sort, desc))
	if err != nil {
		return advertisingList, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	advertisingList.List = advertising
	return advertisingList, nil
}

func (advRepo *AdvertisingRepository) GenerateQueryForGetAdvertisingList(sort string, desc string) string {
	query := "SELECT name,mainPhoto,cost from advertising "

	if sort == "cost" {
		query += "order by cost"
	} else if sort == "created" {
		query += "order by created"
	}

	if desc == "true" {
		query += " DESC"
	} else if desc == "false" {
		query += " ASC"
	}

	query += " LIMIT 10"

	return query
}

func (advRepo *AdvertisingRepository) CheckAdvertisingExist(advertisingID int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT exists (%s)", CheckAdvertisingExist)
	err := advRepo.db.QueryRow(query, advertisingID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return exists, customError.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return exists, nil
}
