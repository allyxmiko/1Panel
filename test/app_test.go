package test

import (
	"testing"

	"github.com/1Panel-dev/1Panel/backend/app/dto/request"
	"github.com/1Panel-dev/1Panel/backend/app/service"
)

func TestApp(t *testing.T) {
	var req request.AppSearch
	appService := service.NewIAppService()
	req.Name = ""
	req.Page = 1
	req.PageSize = 50
	req.Tags = []string{}
	req.Recommend = true
	appService.PageApp(req)
}
