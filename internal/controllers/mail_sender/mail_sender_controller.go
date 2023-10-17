package mail_sender

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iliyasali2107/archiver/internal/dto"
)

type MailSenderCtrl struct {
	svc MailSenderdSvc
}

type MailSenderdSvc interface {
	SendMail(dto.SendMailRequest) error
}

const (
	fileFormKey   = "file"
	emailsFormKey = "emails"
)

func NewMailSenderCtrl(msSvc MailSenderdSvc) *MailSenderCtrl {
	return &MailSenderCtrl{svc: msSvc}
}

func (fsc *MailSenderCtrl) SendFile(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	emailsStr := c.Request.FormValue(emailsFormKey)
	receivers := strings.Split(emailsStr, ",")

	req := dto.SendMailRequest{ReceiverEmails: receivers, FileHeader: fh}

	err = fsc.svc.SendMail(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
