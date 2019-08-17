package rander

import (
	"github.com/Masterminds/sprig"
	"github.com/nonetheless/gorm-migrate/pkg/utils"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type VersionTemplates []*VersionTemplate

type VersionTemplate struct {
	Version     string
	PreVersion  string
	PackageName *string
}

type VersionTemplateContext struct {
	wd          string
	VersionTemps VersionTemplates
	dirPath     string
	PackageName string
	newVersion  *VersionTemplate
}

func (c *VersionTemplateContext) createTemplate(preVersion string) (error) {
	version := utils.GeneratorX(12)
	err := os.Mkdir(c.dirPath+"/"+version, os.ModePerm)
	if err != nil {
		return err
	}
	newVersion := VersionTemplate{
		Version:     version,
		PreVersion:  preVersion,
		PackageName: &c.PackageName,
	}
	data, err := yaml.Marshal(newVersion)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.dirPath+"/"+version+"/version.yaml", data, 0644)
	if err != nil {
		return err
	}
	// create golang file
	inputFile := c.wd + "/template/temp.gotmpl"
	outputFile := c.dirPath + "/" + newVersion.Version + "/migrate.go"
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(inputFile)
	if err != nil {
		return err
	}
	err = utils.WriteTemplate(outputFile, func(writer io.Writer) error {
		err := t.Execute(writer, newVersion)
		return err
	})
	if err != nil {
		return err
	}
	c.newVersion = &newVersion
	c.VersionTemps = append(c.VersionTemps, &newVersion)
	return nil
}

func (c *VersionTemplateContext) createDoc() error {
	inputFile := c.wd + "/template/doc.gotmpl"
	outputFile := c.dirPath + "/doc.go"
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(inputFile)
	if err != nil {
		return err
	}
	err = utils.WriteTemplate(outputFile, func(writer io.Writer) error {
		err := t.Execute(writer, c.newVersion)
		return err
	})
	return err
}

func (c *VersionTemplateContext) createParams() error {
	inputFile := c.wd + "/template/param.gotmpl"
	outputFile := c.dirPath + "/param/param.go"
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(inputFile)
	if err != nil {
		return err
	}
	err = utils.WriteTemplate(outputFile, func(writer io.Writer) error {
		err := t.Execute(writer, c.newVersion)
		return err
	})
	return err
}

func (c *VersionTemplateContext) createInit() error {
	inputFile := c.wd + "/template/init.gotmpl"
	outputFile := c.dirPath + "/init.go"
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(inputFile)
	if err != nil {
		return err
	}
	err = utils.WriteTemplate(outputFile, func(writer io.Writer) error {
		err := t.Execute(writer, c)
		return err
	})
	return err
}

func RanderTemplate(path, packageName string) error {
	dir, err := utils.ReadOrCreate(path)
	if err != nil {
		return err
	}
	versionMap := make(map[string]string)
	for _, fi := range dir {
		if fi.IsDir() {
			if fi.Name() == "param"{
				continue
			}
			version, err := ioutil.ReadFile(path + "/" + fi.Name() + "/version.yaml")
			if err != nil {
				return err
			}
			versionYaml := VersionTemplate{}
			err = yaml.Unmarshal(version, &versionYaml)
			if err != nil {
				return err
			}
			versionMap[versionYaml.PreVersion] = versionYaml.Version
		}
	}
	pre := ""
	versions := make(VersionTemplates, 0)
	for {
		if value, ok := versionMap[pre]; ok {
			versions = append(versions, &VersionTemplate{
				Version:     value,
				PreVersion:  pre,
				PackageName: &packageName,
			})
			pre = value
		} else {
			break
		}
	}
	wd, err := os.Getwd()
	if err != nil {
		return  err
	}
	context := VersionTemplateContext{
		wd:          wd,
		VersionTemps:  versions,
		dirPath:     path,
		PackageName: packageName,
		newVersion:  nil,
	}
	err = context.createTemplate(pre)
	if err != nil {
		return err
	}
	err = context.createDoc()
	if err != nil {
		return err
	}
	err = context.createParams()
	if err != nil {
		return err
	}
	err = context.createInit()
	if err != nil {
		return err
	}
	return nil
}
