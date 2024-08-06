package utils

import "database/sql"

func CityCollect(db *sql.DB) []string {

	var iller []string

	query := "SELECT il_adi FROM iller"

	illerDB, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var il string

	for illerDB.Next() {

		err := illerDB.Scan(&il)

		if err != nil {
			panic(err.Error())
		}

		iller = append(iller, il)
	}

	defer illerDB.Close()
	return iller

}

func GetDistrict(cityName string, db *sql.DB) []string {

	var districts []string
	var district string

	query := "SELECT ilceler.ilce_adi FROM ilceler INNER JOIN iller ON ilceler.il_id = iller.id WHERE iller.il_adi =?"

	districtDB, err := db.Query(query, cityName)

	if err != nil {
		panic(err.Error())
	}

	for districtDB.Next() {
		err := districtDB.Scan(&district)

		if err != nil {
			panic(err.Error())
		}

		districts = append(districts, district)
	}

	defer districtDB.Close()
	return districts
}
