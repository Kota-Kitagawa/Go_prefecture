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

func FetchCities(prefecture string, limit, offset int) ([]string, error) {
    query := `
        SELECT CASE
            WHEN field9 = '以下に掲載がない場合' THEN field8
            ELSE field8 || field9
            END AS city
        FROM addresses
        WHERE field7 = ?
        LIMIT ? OFFSET ?
    `
    rows, err := database.DB.Query(query, prefecture, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var cities []string
    for rows.Next() {
        var city string
        if err := rows.Scan(&city); err != nil {
            return nil, err
        }
        cities = append(cities, city)
    }
    return cities, err
}

func FetchPostal(postalcode, Prefecture, City, Normalizedfield9 string) (string, error) {
    query := `SELECT field3 FROM normalized_utf_ken_all WHERE field7 = ? AND field8 = ? AND Normalizedfield9 LIKE ?`
    rows, err := database.DB.Query(query, Prefecture, City, Normalizedfield9)
    if err != nil {
        return "", err
    }
    defer rows.Close()
    for rows.Next() {
        if err := rows.Scan(&postalcode); err != nil {
            return "", err
        }
    }
    return postalcode, nil
}

func FetchPrefecture() ([]string, error) {
    query := `SELECT DISTINCT field7 FROM addresses`
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var prefectures []string
    for rows.Next() {
        var prefecture string
        if err := rows.Scan(&prefecture); err != nil {
            return nil, err
        }
        prefectures = append(prefectures, prefecture)
    }
    return prefectures, nil
}

func FetchPretoCity()([]string,error){
    query := `SELECT DISTINCT field7 FROM addresses`
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var prefectures []string
    for rows.Next() {
        var prefecture string
        if err := rows.Scan(&prefecture); err != nil {
            return nil, err
        }
        prefectures = append(prefectures, prefecture)
    }
    return prefectures, nil
}