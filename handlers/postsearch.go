package handlers

import (
    "log"
    "net/http"
    "regexp"
    "Go_prefecture/database"
    "github.com/gin-gonic/gin"
)

func PostSearchHandler(c *gin.Context){
    Prefecture := c.PostForm("prefecture")
    City := c.PostForm("city")
    Detail := c.PostForm("detail")
    log.Printf("Received address: %s-%s-%s", Prefecture, City, Detail)

    if Prefecture == "" || City == "" || Detail == "" {
        c.String(http.StatusBadRequest, "Address not specified")
        return
    }

    re := regexp.MustCompile(`\s+`)
    Prefecture = re.ReplaceAllString(Prefecture, "")
    City = re.ReplaceAllString(City, "")
    Detail = re.ReplaceAllString(Detail, "")

    var postalcode string

    query := `SELECT field3 FROM normalized_utf_ken_all WHERE field7 = ? AND field8 = ? AND Normalizedfield9 LIKE ?`
    Normalizedfield9 := "%" + Detail + "%"
    rows, err := database.DB.Query(query, Prefecture, City, Normalizedfield9)
    if err != nil {
    	log.Printf("Failed to fetch address: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch address")
        return
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&postalcode)
        if err != nil {
            log.Printf("Failed to scan postal code: %v", err)
            c.String(http.StatusInternalServerError, "Failed to scan postal code")
            return
        }
        log.Printf("Fetched fields: field7: %s, field8: %s, Normalizedfield9: %s, postalcode: %s", Prefecture, City, Detail, postalcode)
        c.HTML(http.StatusOK, "postresult.html", gin.H{
            "PostCode": postalcode,
        })
    } else {
        log.Printf("No postal code found for address: %s-%s-%s", Prefecture, City, Detail)
        c.String(http.StatusNotFound, "No postal code found for the given address")
    }
}