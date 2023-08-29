package core

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/blocks"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

type Config struct {
	Port string `yaml:"port"`
}

type Moujin struct {
	Views  *blocks.Blocks
	Config *Config
	Router *gin.Engine
}

func (c *Config) SetDefaultConf() error {
	yamlFile, err := os.ReadFile("./conf.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}

func (m *Moujin) SetDefaultViews() error {
	m.Views = blocks.New("./src/views")
	err := m.Views.Load()
	if err != nil {
		return err
	}
	return nil
}

func (m *Moujin) SetDefaultRouter() {
	m.Router = gin.Default()
	m.Router.StaticFile("/favicon.ico", "./src/views/assets/reimu.png")
	m.Router.Static("/css", "./src/views/styles")
	m.Router.Static("/js", "./src/views/scripts")
	m.Router.Static("/ass", "./src/views/assets")
	m.Router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", ContentTypeHTML)
		data := map[string]interface{}{
			"s":     http.StatusOK,
			"Title": "Home",
		}
		err := m.Views.ExecuteTemplate(c.Writer, "home", "main", data)
		if err != nil {
			log.Printf("%s", err)
		}
	})
	m.Router.GET("/aboutMe", func(c *gin.Context) {
		c.Header("Content-Type", ContentTypeHTML)
		data := map[string]interface{}{
			"s":     http.StatusOK,
			"Title": "About me",
		}
		err := m.Views.ExecuteTemplate(c.Writer, "aboutMe", "main", data)
		if err != nil {
			log.Printf("%s", err)
		}
	})
}
