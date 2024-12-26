package pkg

func GetResponse(format string, templateName string)ResponseStrategy{
	if format == "json"{
		return &JSONResponse{}
	}
	return &HTMLResponse{TemplateName:templateName}
}