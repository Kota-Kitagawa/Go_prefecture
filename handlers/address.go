package handlers

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "Go_prefecture/database"
)

func AddressHandler(c *gin.Context) {
    postalCode1 := c.PostForm("postalcode1")
    postalCode2 := c.PostForm("postalcode2")
    log.Printf("Received postal codes: %s-%s", postalCode1, postalCode2)

    if postalCode1 == "" || postalCode2 == "" {
        c.String(http.StatusBadRequest, "Postal code not specified")
        return
    }

    postalCode := postalCode1 + postalCode2
    log.Printf("Combined postal code: %s", postalCode)

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
        log.Printf("Failed to fetch address: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch address")
        return
    }

    log.Printf("Fetched fields: field7: %s, field8: %s, field9: %s, fullAddress: %s", field7, field8, field9, fullAddress)

    log.Printf("Constructed full address: %s", fullAddress)

    c.HTML(http.StatusOK, "addressresult.html", gin.H{
        "Field7": field7,
        "Field8": field8,
        "Field9": field9,
        "FullAddress": fullAddress,
    })
}
