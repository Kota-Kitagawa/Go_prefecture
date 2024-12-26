package handlers

import (
    "net/http"
    "Go_prefecture/internal/database"
    "github.com/gin-gonic/gin"
)

func fetchAddress(postalCode string) (string, string, string, string, error) {
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

func AddressHandler(c *gin.Context) {
    postalCode := c.Query("postalcode1") + c.Query("postalcode2")
    field7, field8, field9, fullAddress, err := fetchAddress(postalCode)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error fetching address")
        return
    }
    responseFormat := c.Query("format")
    if responseFormat == "json" {
        c.JSON(http.StatusOK, gin.H{
            "Field7":      field7,
            "Field8":      field8,
            "Field9":      field9,
            "FullAddress": fullAddress,
        })
    } else {
        c.HTML(http.StatusOK, "addressresult.html", gin.H{
            "Field7":      field7,
            "Field8":      field8,
            "Field9":      field9,
            "FullAddress": fullAddress,
        })
    }
}