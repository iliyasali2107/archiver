package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iliyasali2107/archiver/internal/config"
	"github.com/iliyasali2107/archiver/internal/controllers"
	"github.com/iliyasali2107/archiver/internal/services/archive"
	"github.com/iliyasali2107/archiver/internal/services/mail"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Println("qwerqwer")
		log.Fatal(err)
	}

	archiveSvc := archive.NewArchiveSvc()
	mailSvc := mail.NewMailSvc(cfg)

	controller := controllers.NewController(archiveSvc, archiveSvc, mailSvc)

	r := gin.Default()
	archiveGroup := r.Group("/archive")

	archiveGroup.POST("/info", controller.ArchiveInfoCtrl.GetArchiveInfo)
	archiveGroup.POST("/compress", controller.ArchiveCompresCtrl.Compress)
	archiveGroup.POST("/send", controller.MailSenderCtrl.SendFile)

	fmt.Println("running ...")
	r.Run(":8080")
}
