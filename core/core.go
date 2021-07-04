package core

import (
	"bytes"
	"carapi/core/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func generateFiber(rootpath, apppath string, buf bytes.Buffer) {
	// root
	write(rootpath, "config.toml", model.FiberConfigTomlFile, &buf)

	// app/
	write(apppath, "main.go", model.FiberMainFile, &buf)

	// app/api
	var apipath = path.Join(apppath, "api")
	_ = os.MkdirAll(apipath, os.ModePerm)
	write(apipath, "test.go", model.FiberApiTestFile, &buf)
	write(apipath, "task.go", model.FiberApiTaskFile, &buf)

	// app/config
	var cfgpath = path.Join(apppath, "config")
	_ = os.MkdirAll(cfgpath, os.ModePerm)
	write(cfgpath, "config.go", model.FiberConfigGoFile, &buf)

	// app/schema
	var schemapath = path.Join(apppath, "schema")
	_ = os.MkdirAll(schemapath, os.ModePerm)
	write(schemapath, "common.go", model.FiberSchemaCommonFile, &buf)
	write(schemapath, "task.go", model.FiberSchemaTaskFile, &buf)

	// app/setup
	var setuppath = path.Join(apppath, "setup")
	_ = os.MkdirAll(setuppath, os.ModePerm)
	write(setuppath, "config.go", model.SetupConfigFile, &buf)
	write(setuppath, "database.go", model.SetupDatabaseFile, &buf)
	write(setuppath, "logger.go", model.SetupLoggerFile, &buf)

	write(setuppath, "engine.go", model.FiberSetupEngineFile, &buf)
	write(setuppath, "router.go", model.FiberSetupRouterFile, &buf)
}

func generateGin(rootpath, apppath string, buf bytes.Buffer) {
	// root
	write(rootpath, "config.toml", model.GinConfigTomlFile, &buf)

	// app/
	write(apppath, "main.go", model.GinMainFile, &buf)

	// app/api
	var apipath = path.Join(apppath, "api")
	_ = os.MkdirAll(apipath, os.ModePerm)
	write(apipath, "test.go", model.GinApiTestFile, &buf)
	write(apipath, "task.go", model.GinApiTaskFile, &buf)

	// app/config
	var cfgpath = path.Join(apppath, "config")
	_ = os.MkdirAll(cfgpath, os.ModePerm)
	write(cfgpath, "config.go", model.GinConfigGoFile, &buf)

	// app/middleware
	var midpath = path.Join(apppath, "middleware")
	_ = os.MkdirAll(midpath, os.ModePerm)
	write(midpath, "cors.go", model.GinMiddlewareCorsFile, &buf)
	write(midpath, "logger.go", model.GinMiddlewareLoggerFile, &buf)
	write(midpath, "swagger.go", model.GinMiddlewareSwaggerFile, &buf)

	// app/schema
	var schemapath = path.Join(apppath, "schema")
	_ = os.MkdirAll(schemapath, os.ModePerm)
	write(schemapath, "common.go", model.GinSchemaCommonFile, &buf)
	write(schemapath, "task.go", model.GinSchemaTaskFile, &buf)

	// app/setup
	var setuppath = path.Join(apppath, "setup")
	_ = os.MkdirAll(setuppath, os.ModePerm)
	write(setuppath, "config.go", model.SetupConfigFile, &buf)
	write(setuppath, "database.go", model.SetupDatabaseFile, &buf)
	write(setuppath, "logger.go", model.SetupLoggerFile, &buf)

	write(setuppath, "engine.go", model.GinSetupEngineFile, &buf)
	write(setuppath, "router.go", model.GinSetupRouterFile, &buf)
}

func Execute(name, output, remote, frame string, docker, compose bool) (err error) {
	// root dir
	var rootpath = path.Join(output, name)
	if err = os.MkdirAll(rootpath, os.ModePerm); err != nil {
		return err
	}
	var buf bytes.Buffer

	// app/git
	cmd := exec.Command("git", "init")
	cmd.Dir = rootpath
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("git init failed: %v", err)
	}
	if remote != "" {
		cmd = exec.Command("git", "remote", "add", "origin", remote)
		cmd.Dir = rootpath
		if err = cmd.Run(); err != nil {
			return fmt.Errorf("git add remote error: %v", err)
		}
	}
	write(rootpath, ".gitignore", model.GitIgnoreFile, &buf)

	// app dir
	var apppath string
	switch compose {
	case true:
		write(rootpath, "release.sh", model.ReleaseFile, &buf)
		apppath = path.Join(rootpath, "services/app")
		_ = os.MkdirAll(apppath, os.ModePerm)
		write(rootpath, "docker-compose.local.yaml", model.DockerComposeLocalFile, &buf)
		write(rootpath, "docker-compose.yaml", model.DockerComposeReleaseFile, &buf)
	case false:
		apppath = rootpath
	}

	switch frame {
	case "gin":
		generateGin(rootpath, apppath, buf)
	case "fiber":
		generateFiber(rootpath, apppath, buf)
	}

	// app/dockerfile
	if docker {
		write(apppath, "build.dockerfile", model.BuildDockerFile, &buf)
		write(apppath, "local.dockerfile", model.LocalDockerFile, &buf)
	}

	// app/docs
	var docpath = path.Join(apppath, "docs")
	_ = os.MkdirAll(docpath, os.ModePerm)
	write(docpath, "docs.go", model.DocGoFile, &buf)
	write(docpath, "swagger.json", model.DocSwaggerJsonFile, &buf)
	write(docpath, "swagger.yaml", model.DocSwaggerYamlFile, &buf)

	// app/model
	var modelpath = path.Join(apppath, "model")
	_ = os.MkdirAll(modelpath, os.ModePerm)
	write(modelpath, "base.go", model.ModelBaseFile, &buf)
	write(modelpath, "task.go", model.ModelTaskFile, &buf)

	// app/service
	var servicepath = path.Join(apppath, "service")
	_ = os.MkdirAll(servicepath, os.ModePerm)
	write(servicepath, "task.go", model.ServiceTaskFile, &buf)

	// app/public
	var pubpath = path.Join(apppath, "public")
	_ = os.MkdirAll(pubpath, os.ModePerm)
	write(pubpath, "const.go", model.PublicConstFile, &buf)
	write(pubpath, "handle.go", model.PublicHandleFile, &buf)
	write(pubpath, "utils.go", model.PublicUtilsFile, &buf)

	// go mod
	log.Println("Running `go mod init app`")
	cmd = exec.Command("go", "mod", "init", "app")
	cmd.Dir = apppath
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("go mod init error: %v", err)
	}
	log.Println("Downloading dependency Library...")
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = apppath
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("go mod tidy error: %v", err)
	}
	log.Printf("Success!, cd %s && go run main.go", apppath)
	return nil
}

func write(p, n, v string, buf *bytes.Buffer) {
	buf.WriteString(v)
	p = path.Join(p, n)
	_ = ioutil.WriteFile(p, buf.Bytes(), 0644)
	log.Printf("write to file %s success\n", p)
	buf.Reset()
}
