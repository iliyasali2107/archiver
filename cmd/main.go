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

	archiveSvc := &services.ArchiveService{}

	controller := controllers.NewController(archiveSvc, archiveSvc)

	archiveGroup.POST("/info", controller.ArchiveInfoCtrl.GetArchiveInfo)
	archiveGroup.POST("/compress", controller.ArchiveCompresCtrl.Compress)

	fmt.Println("running ...")
	r.Run(":8080")

}
