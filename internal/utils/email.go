package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"path/filepath"
)

func SendEmailWithAttachment(
	smtpHost, smtpPort, from, password,
	to, subject, body, filePath string,
) error {

	// Read file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// MIME boundary
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Headers
	headers := textproto.MIMEHeader{}
	headers.Set("From", from)
	headers.Set("To", to)
	headers.Set("Subject", subject)
	headers.Set("MIME-Version", "1.0")
	headers.Set("Content-Type", "multipart/mixed; boundary="+writer.Boundary())

	for k, v := range headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", k, v[0]))
	}
	buf.WriteString("\r\n")

	// Body
	bodyPart, _ := writer.CreatePart(
		textproto.MIMEHeader{
			"Content-Type": {"text/plain"},
		},
	)
	bodyPart.Write([]byte(body))

	// Attachment
	attachmentPart, _ := writer.CreatePart(
		textproto.MIMEHeader{
			"Content-Type":              {"application/octet-stream"},
			"Content-Transfer-Encoding": {"base64"},
			"Content-Disposition": {
				fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(filePath)),
			},
		},
	)

	encoded := base64.StdEncoding.EncodeToString(fileData)
	attachmentPart.Write([]byte(encoded))

	writer.Close()

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{to},
		buf.Bytes(),
	)
}
