package models
import (
	"database/sql"
	"Go_prefecture/database"
)

func GetPosttalCode(prefecture,city,town,street string) (string,error) {
	var postalCode string
	query :=`SELECT field3 FORM address 
			WHERE field7 = ? AND field8 = ? AND field9 = ?
			ORDER BY 
			CASE WHEN ? LIKE field4 || '%' THEN 1 ELSE 2 END
			LIMIT 1`
	err := database.DB.QueryRow(query,prefecture,city,town,street).Scan(&postalCode)
	if err != nil {
		return "",err
	}
	return postalCode,nil 
}