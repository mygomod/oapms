package service

import (
	"github.com/flosch/pongo2"
	"github.com/smartwalle/pongo2render"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Mailer *mailer

type mailer struct {
	*gomail.Dialer
	tmpls map[string]string
}

func InitMailer() {
	Mailer = &mailer{
		Dialer: gomail.NewDialer(viper.GetString("email.host"), viper.GetInt("email.port"), viper.GetString("email.username"), viper.GetString("email.password")),
		tmpls:  make(map[string]string, 0),
	}
	Mailer.initLoadTmpl()
}

func (m *mailer) initLoadTmpl() {
	tmplRepoDir := "conf/emailtmpl/"
	err := filepath.Walk(tmplRepoDir,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && info.Name() == ".git" {
				return filepath.SkipDir
			}
			if strings.HasSuffix(info.Name(), ".go") {
				return nil
			}
			b, e := ioutil.ReadFile(path)
			if e != nil {
				return nil
			}
			relPath, e := filepath.Rel(tmplRepoDir, path)
			if e != nil {
				return nil
			}
			m.tmpls[relPath] = string(b)
			return nil
		})
	if err != nil {
		panic("load email tmpl err: " + err.Error())
	}
}

func (m *mailer) Send(subject, to string, html string, attachment string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", viper.GetString("email.from"))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", html)
	if attachment != "" {
		msg.Attach(attachment)
	}
	// Send the email to Bob, Cora and Dan.
	if err := m.Dialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func (m *mailer) ParseTpl(filename string, data map[string]interface{}) (html string, err error) {
	var render = pongo2render.NewRender(filename)
	ctx := pongo2.Context(data)
	html, err = render.TemplateFromString(m.tmpls[filename]).Execute(ctx)
	if err != nil {
		return
	}
	return
}
