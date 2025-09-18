package email

import (
	"bytes"
	"embed"
	"html/template"
	"time"
)


//go:embed templates/*.html
var emailTemplates embed.FS

// TemplateData is a helper to ensure common fields like Title and Year are always available
type TemplateData map[string]interface{}

// RenderTemplate renders an email template with the given data
func RenderTemplate(templateName string, data TemplateData) (string, error) {
	// Add default/common values
	if _, ok := data["Title"]; !ok {
		data["Title"] = "Go SaaS Kit"
	}
	if _, ok := data["Year"]; !ok {
		data["Year"] = time.Now().Year()
	}

	// Parse layout + content template
	tmpl, err := template.ParseFS(emailTemplates,
		"templates/layout.html",
		"templates/"+templateName,
	)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := tmpl.ExecuteTemplate(&body, "layout", data); err != nil {
		return "", err
	}

	return body.String(), nil
}
