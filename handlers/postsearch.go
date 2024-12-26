package handlers

import (
    "log"
    "net/http"
    "regexp"
    "Go_prefecture/internal/database"
    "github.com/gin-gonic/gin"
)

func fetchPostal(postalcode, Prefecture, City, Normalizedfield9 string) (string, error) {
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

func PostSearchHandler(c *gin.Context) {
    Prefecture := c.Query("prefecture")
    City := c.Query("city")
    Detail := c.Query("detail")
    log.Printf("Received address: %s-%s-%s", Prefecture, City, Detail)
    if Prefecture == "" || City == "" || Detail == "" {
        c.String(http.StatusBadRequest, "Address not specified")
        return
    }
    re := regexp.MustCompile(`\s+`)
    Prefecture = re.ReplaceAllString(Prefecture, "")
    City = re.ReplaceAllString(City, "")
    Detail = re.ReplaceAllString(Detail, "")

    postalcode, err := fetchPostal("", Prefecture, City, Detail)
    if err != nil {
        log.Printf("Failed to fetch address: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch address")
        return
    }
    responseFormat := c.Query("format")
    if responseFormat == "json" {
        c.JSON(http.StatusOK, gin.H{
            "PostalCode": postalcode,
        })
    } else {
        c.HTML(http.StatusOK, "postresult.html", gin.H{
            "PostalCode": postalcode,
        })
    }
    
}


func PostalHandler(c *gin.Context) {
	c.HTML(200, "postcode.html", nil)
}

func AddressSearchHandler(c *gin.Context){
	c.HTML(200,"postsearch.html",nil)
}