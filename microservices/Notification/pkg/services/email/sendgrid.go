package email_service

import (
	"bytes"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"notificaiton/pkg/structs"
	"notificaiton/pkg/utils/common"
	"os"
	"strings"
)

type SendGrid struct {
}

func (sndgrd *SendGrid) SendEmail(emailFeatures structs.EmailFeatures) *common.ResponseMessage {
	fromString := emailFeatures.From
	plainText := emailFeatures.PlainText
	htmlContent := emailFeatures.HtmlContent
	htmlTemplateData := emailFeatures.HtmlTemplateData
	dir, _ := os.Getwd()
	htmlContentTemplateFileName := ""
	info, fileExistError := os.Stat(dir + "/pkg/templates/email/" + emailFeatures.HtmlContentTemplateFileName)
	if !os.IsNotExist(fileExistError) && !info.IsDir() {
		htmlContentTemplateFileName = dir + "/pkg/templates/email/" + emailFeatures.HtmlContentTemplateFileName
	}

	subject := emailFeatures.Subject
	envelopes := emailFeatures.Tos

	m := mail.NewV3Mail()
	from := mail.NewEmail(fromString["name"], fromString["email"])
	m.SetFrom(from)

	personalization := mail.NewPersonalization()
	//personalization.AddFrom(from)
	var tos []*mail.Email
	for _, to := range envelopes {
		emailAddress := mail.NewEmail(to["name"], to["email"])
		tos = append(tos, emailAddress)
	}
	personalization.AddTos(tos...)
	personalization.Subject = subject
	m.AddPersonalizations(personalization)
	if strings.Trim(plainText, " ") != "" {
		plainTextContent := mail.NewContent("text/plain", plainText)
		m.AddContent(plainTextContent)
	}

	if strings.Trim(htmlContent, " ") != "" {
		htmlTextContent := mail.NewContent("text/html", htmlContent)
		m.AddContent(htmlTextContent)
	}

	if strings.Trim(htmlContentTemplateFileName, " ") != "" {
		templateContent, err := template.ParseFiles(htmlContentTemplateFileName)
		if err == nil {
			buf := new(bytes.Buffer)
			templateContentError := templateContent.Execute(buf, htmlTemplateData)
			if templateContentError == nil {
				htmlTextContent := mail.NewContent("text/html", buf.String())
				m.AddContent(htmlTextContent)
			}
		}

	}

	//message := mail.NewSingleEmail(from, subject, emailAddress, plainText, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	emailResponse, err := client.Send(m)
	if err != nil {
		return common.ReplyUtil(false, "", err.Error())
	}
	return common.ReplyUtil(true, "", emailResponse.Body)
}

func (sndgrd *SendGrid) IsActiveProvider() bool {
	return true
}
