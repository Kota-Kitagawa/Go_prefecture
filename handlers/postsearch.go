package handlers

import (
    "log"
    "net/http"
    "regexp"
    "Go_prefecture/pkg"
    "github.com/gin-gonic/gin"
)

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

    postalcode, err := pkg.FetchPostal("", Prefecture, City, Detail)
    if err != nil {
        log.Printf("Failed to fetch address: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch address")
        return
    }
    responseFormat := c.Query("format")
    res :=pkg.GetResponse(responseFormat,"postresult.html")
    res.Respond(c,gin.H{
        "PostalCode": postalcode,
    })
}
