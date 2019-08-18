package rander

import (
	"github.com/Masterminds/sprig"
	api "github.com/nonetheless/gorm-migrate/pkg"
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

func (v *VersionTemplate) Printf(out api.MigrateOut){
	tempVersion := v.PreVersion
	if tempVersion == "" {
		tempVersion = "            "
	}
	out.Infof(tempVersion + "   --------->     " + v.Version + "\n")
}

func (v *VersionTemplate) RPrintf(out api.MigrateOut){
	tempVersion := v.PreVersion
	if tempVersion == "" {
		tempVersion = "            "
	}
	out.Infof(tempVersion + "   <---------     " + v.Version + "\n")
}

type VersionTemplateContext struct {
	VersionTemps VersionTemplates
	dirPath      string
	PackageName  string
	newVersion   *VersionTemplate
	out          api.MigrateOut
	preVersion   string
}

func (c *VersionTemplateContext) createTemplate() (error) {
	version := utils.GeneratorX(12)
	err := os.Mkdir(c.dirPath+"/"+version, os.ModePerm)
	if err != nil {
		return err
	}
	newVersion := VersionTemplate{
		Version:     version,
		PreVersion:  c.preVersion,
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
	inputFile := "template/temp.gotmpl"
	outputFile := c.dirPath + "/" + newVersion.Version + "/migrate.go"
	data, err = utils.AssetTemplate("template/temp.gotmpl")
	if err != nil {
		return err
	}
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).Parse(string(data))
	if err != nil {
		return err
	}
	err = utils.WriteTemplate(outputFile, func(writer io.Writer) error {
		err := t.Execute(writer, newVersion)
		return err
	})
	c.out.Infof("create new version:" + newVersion.Version + "\n")
	if err != nil {
		return err
	}
	c.newVersion = &newVersion
	c.VersionTemps = append(c.VersionTemps, &newVersion)
	return nil
}

func (c *VersionTemplateContext) createDoc() error {
	inputFile := "template/doc.gotmpl"
	outputFile := c.dirPath + "/doc.go"
	data, err := utils.AssetTemplate("template/doc.gotmpl")
	if err != nil {
		return err
	}
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).Parse(string(data))
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
	inputFile := "template/param.gotmpl"
	outputFile := c.dirPath + "/param/param.go"
	data, err := utils.AssetTemplate("template/param.gotmpl")
	if err != nil {
		return err
	}
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).Parse(string(data))
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
	inputFile := "template/init.gotmpl"
	outputFile := c.dirPath + "/init.go"
	data, err := utils.AssetTemplate("template/init.gotmpl")
	if err != nil {
		return err
	}
	t, err := template.New(filepath.Base(inputFile)).Funcs(sprig.TxtFuncMap()).Parse(string(data))
	if err != nil {
		return err
	}
	err = utils.WriteTemplate(outputFile, func(writer io.Writer) error {
		err := t.Execute(writer, c)
		return err
	})
	c.out.Infof("Migration version list:\n")
	printVersion(c)
	return err
}

func printVersion(c *VersionTemplateContext) {
	for _, version := range c.VersionTemps {
		version.Printf(c.out)
	}
}

func createRanderContext(dir []os.FileInfo, path, packageName string, out api.MigrateOut) (*VersionTemplateContext, error) {
	versionMap := make(map[string]string)
	for _, fi := range dir {
		if fi.IsDir() {
			if fi.Name() == "param" {
				continue
			}
			version, err := ioutil.ReadFile(path + "/" + fi.Name() + "/version.yaml")
			if err != nil {
				return nil, err
			}
			versionYaml := VersionTemplate{}
			err = yaml.Unmarshal(version, &versionYaml)
			if err != nil {
				return nil, err
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
	context := VersionTemplateContext{
		VersionTemps: versions,
		dirPath:      path,
		PackageName:  packageName,
		newVersion:   nil,
		out:          out,
		preVersion:   pre,
	}
	return &context, nil

}

func RanderTemplate(path, packageName string, out api.MigrateOut) error {
	dir, err := utils.ReadOrCreate(path)
	if err != nil {
		return err
	}
	context, err := createRanderContext(dir, path, packageName, out)
	if err != nil {
		return err
	}
	err = context.createDoc()
	if err != nil {
		return err
	}
	err = context.createTemplate()
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

func Stamp(path string, out api.MigrateOut) error{
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	context, err := createRanderContext(dir, path, "packageName", out)
	if err != nil {
		return err
	}
	context.out.Infof("Stamp test_migration version:\n")
	printVersion(context)
	return nil
}