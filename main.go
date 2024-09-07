package main

import (
	"html/template"
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Contact struct {
		FirstName string `yaml:"first name"`
		LastName  string `yaml:"last name"`
		Mob       string `yaml:"mob"`
		Email     string `yaml:"email"`
		Web       string `yaml:"website"`
		Location  string `yaml:"location,omitempty"`
	} `yaml:"contact"`

	Experience []struct {
		Role    string   `yaml:"role"`
		Org     string   `yaml:"org"`
		Summary []string `yaml:"summary"`
		Start   string   `yaml:"start"`
		End     string   `yaml:"end"`
	} `yaml:"experience,omitempty"`

	Education []struct {
		Title       string `yaml:"title"`
		Institution string `yaml:"institution"`
		Result      string `yaml:"result"`
		Description string `yaml:"desc"`
		Start       string `yaml:"start"`
		End         string `yaml:"end"`
	} `yaml:"education,omitempty"`

	References []struct {
		Fullname string `yaml:"name"`
		Org      string `yaml:"org"`
		Email    string `yaml:"email"`
		Phone    string `yaml:"phone"`
	} `yaml:"references,omitempty"`

	Projects []struct {
		Name        string `yaml:"name"`
		Permalink   string `yaml:"link"`
		Description string `yaml:"description"`
	} `yaml:"projects,omitempty"`

	Skills template.HTML `yaml:"skills,omitempty"`
	Extra  template.HTML `yaml:"extra,omitempty"`
}

func main() {
	// read config file content
	cf, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// unmasrshal config content
	c := Config{}
	err = yaml.Unmarshal(cf, &c)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("index.tmpl", "header.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.OpenFile("index.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(outputFile, c)
	if err != nil {
		log.Fatal(err)
	}
	outputFile.Close()
}
