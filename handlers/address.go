package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "Go_prefecture/internal/pkg"
)


func AddressHandler(c *gin.Context) {
    postalCode := c.Param("postalCode")
    field7, field8, field9, fullAddress, err := pkg.FetchAddress(postalCode)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error fetching address")
        return
    }
    responseFormat := c.Query("format")
    res := pkg.GetResponse(responseFormat,"postCode.html")
    res.Respond(c,gin.H{
        "Field7":      field7,
        "Field8":      field8,
        "Field9":      field9,
        "FullAddress": fullAddress,
    })
}