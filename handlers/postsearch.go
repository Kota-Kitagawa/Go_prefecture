package handlers

import (
    "log"
    "net/http"
    "regexp"
    "Go_prefecture/pkg/database"
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

func PostSearchHTMLHandler(c *gin.Context) {
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

    postalcode, err := fetchPostal("", Prefecture, City, Detail)
    if err != nil {
        log.Printf("Failed to fetch address: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch address")
        return
    }
    c.HTML(http.StatusOK, "postresult.html", gin.H{
        "PostCode": postalcode,
    })
}

func PostSearchJSONHandler(c *gin.Context) {
    prefecture := c.PostForm("prefecture")
    log.Printf("Received prefecture: %s", prefecture)

    cities, err := fetchCities(prefecture)
    if err != nil {
        log.Printf("Failed to fetch cities: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cities"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "cities": cities,
    })
}
func PostalHandler(c *gin.Context) {
	c.HTML(200, "postcode.html", nil)
}

func AddressSearchHandler(c *gin.Context){
	c.HTML(200,"postsearch.html",nil)
}