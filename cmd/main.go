package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iliyasali2107/archiver/internal/controllers"
	"github.com/iliyasali2107/archiver/internal/services"
)

func main() {

	r := gin.Default()

	archiveGroup := r.Group("/archive")

	archiveInfoSvc := services.ArchiveService{}

	controller := controllers.NewController(&archiveInfoSvc)

	archiveGroup.POST("/info", controller.ArchiveInfoCtrl.GetArchiveInfo)

	fmt.Println("running ...")
	r.Run(":8080")

}
