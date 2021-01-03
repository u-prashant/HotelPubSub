package database

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// StoreOfferInDB stores the hotel offer into db
// TODO:
// - add error handling of each Create operation
// - perform all the operations as single transaction for rolling back ?
func (db *DbCtxt) StoreOfferInDB(offers *HotelOffers) {
	var (
		hotels    []*Hotel
		rooms     []*Room
		ratePlans []*RatePlan
	)
	for _, offer := range offers.Offers {
		var hotel *Hotel
		err := json.Unmarshal(offer.Hotel, &hotel)
		if err != nil {
			log.Errorf("error in unmarshalling hotel data[%s] err[%s]", offer.Hotel, err.Error())
			return
		}
		hotel.Object = string(offer.Hotel)
		hotels = append(hotels, hotel)
		rooms = append(rooms, &Room{HotelID: hotel.ID, Object: string(offer.Room)})
		ratePlans = append(ratePlans, &RatePlan{HotelID: hotel.ID, Object: string(offer.RatePlan)})
	}
	db.client.Table("hotels").CreateInBatches(hotels, len(hotels))
	db.client.Table("rooms").CreateInBatches(rooms, len(rooms))
	db.client.Table("rate_plans").CreateInBatches(ratePlans, len(ratePlans))
	return
}
