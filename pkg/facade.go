package pkg
import(
	"Go_prefecture/internal/database"
)

func FetchAddress(postalCode string) (string, string, string, string, error) {
    var field7, field8, field9, fullAddress string

    query := `
        SELECT field7, field8, field9,
        CASE
            WHEN field9 = '以下に掲載がない場合' THEN field7 || field8
            ELSE field7 || field8 || field9
        END AS Fulladdress
        FROM addresses
        WHERE field3 = ?
    `
    err := database.DB.QueryRow(query, postalCode).Scan(&field7, &field8, &field9, &fullAddress)
    if err != nil {
        return "", "", "", "", err
    }

    return field7, field8, field9, fullAddress, nil
}