package handlers

import (
    "log"
    "net/http"
    "Go_prefecture/pkg/database"
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

func AddressHTMLHandler(c *gin.Context) {
    postalCode1 := c.PostForm("postalcode1")
    postalCode2 := c.PostForm("postalcode2")
    log.Printf("Received postal codes: %s-%s", postalCode1, postalCode2)

    if postalCode1 == "" || postalCode2 == "" {
        c.String(http.StatusBadRequest, "Postal code not specified")
        return
    }

    postalCode := postalCode1 + postalCode2
    log.Printf("Combined postal code: %s", postalCode)

    field7, field8, field9, fullAddress, err := fetchAddress(postalCode)
    if err != nil {
        log.Printf("Failed to fetch address: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch address")
        return
    }

    c.HTML(http.StatusOK, "addressresult.html", gin.H{
        "Field7": field7,
        "Field8": field8,
        "Field9": field9,
        "FullAddress": fullAddress,
    })
}

func AddressJSONHandler(c *gin.Context) {
    postalCode1 := c.PostForm("postalcode1")
    postalCode2 := c.PostForm("postalcode2")
    log.Printf("Received postal codes: %s-%s", postalCode1, postalCode2)

    if postalCode1 == "" || postalCode2 == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Postal code not specified"})
        return
    }

    postalCode := postalCode1 + postalCode2
    log.Printf("Combined postal code: %s", postalCode)

    field7, field8, field9, fullAddress, err := fetchAddress(postalCode)
    if err != nil {
        log.Printf("Failed to fetch address: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch address"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Field7": field7,
        "Field8": field8,
        "Field9": field9,
        "FullAddress": fullAddress,
    })
}