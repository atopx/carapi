package core

import (
	"bytes"
	"fmt"
	"ginhelper/core/model"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func Execute(name, output, remote string, docker, compose bool) (err error) {
	// root dir
	var rootpath = path.Join(output, name)
	if err = os.MkdirAll(rootpath, os.ModePerm); err != nil {
		return err
	}

	var buf bytes.Buffer
	var apppath string

	switch compose {
	case true:
		apppath = path.Join(rootpath, "services/app")
		_ = os.MkdirAll(apppath, os.ModePerm)
		write(rootpath, "docker-compose.local.yaml", model.DockerComposeLocalFile, &buf)
		write(rootpath, "docker-compose.yaml", model.DockerComposeReleaseFile, &buf)
		write(rootpath, "config.yaml", model.ConfigYamlFile, &buf)
	case false:
		apppath = rootpath
	}

	// git
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

	// app
	write(apppath, "main.go", model.MainFile, &buf)
	if docker {
		write(apppath, "build.dockerfile", model.BuildDockerFile, &buf)
		write(apppath, "local.dockerfile", model.LocalDockerFile, &buf)
	}

	// api
	var apipath = path.Join(apppath, "api")
	_ = os.MkdirAll(apipath, os.ModePerm)
	write(apipath, "test.go", model.ApiTestFile, &buf)

	// config
	var cfgpath = path.Join(apppath, "config")
	_ = os.MkdirAll(cfgpath, os.ModePerm)
	write(cfgpath, "config.go", model.ConfigGoFile, &buf)
	write(apppath, "config.yaml", model.ConfigYamlFile, &buf)

	// core
	var corepath = path.Join(apppath, "core")
	_ = os.MkdirAll(corepath, os.ModePerm)

	// docs
	var docpath = path.Join(apppath, "docs")
	_ = os.MkdirAll(docpath, os.ModePerm)
	write(docpath, "docs.go", model.DocGoFile, &buf)
	write(docpath, "swagger.json", model.DocSwaggerJsonFile, &buf)
	write(docpath, "swagger.yaml", model.DocSwaggerYamlFile, &buf)

	// libs
	var libpath = path.Join(apppath, "libs")
	_ = os.MkdirAll(libpath, os.ModePerm)
	write(libpath, "orm_logger.go", model.LibOrgLoggerFile, &buf)

	// middleware
	var midpath = path.Join(apppath, "middleware")
	_ = os.MkdirAll(midpath, os.ModePerm)
	write(midpath, "cors.go", model.MiddlewareCorsFile, &buf)
	write(midpath, "logger.go", model.MiddlewareLoggerFile, &buf)
	write(midpath, "swagger.go", model.MiddlewareSwaggerFile, &buf)

	// models
	var modelpath = path.Join(apppath, "models")
	_ = os.MkdirAll(modelpath, os.ModePerm)
	write(modelpath, "user.go", model.ModelUserFile, &buf)

	// public
	var pubpath = path.Join(apppath, "public")
	_ = os.MkdirAll(pubpath, os.ModePerm)
	write(pubpath, "const.go", model.PublicConstFile, &buf)
	write(pubpath, "variable.go", model.PublicVariableFile, &buf)

	// routes
	var routepath = path.Join(apppath, "routers")
	_ = os.MkdirAll(routepath, os.ModePerm)
	write(routepath, "test.go", model.RouterTestFile, &buf)

	// schemas
	var schemapath = path.Join(apppath, "schemas")
	_ = os.MkdirAll(schemapath, os.ModePerm)
	write(schemapath, "common.go", model.SchemaCommonFile, &buf)
	write(schemapath, "response.go", model.SchemaResponseFile, &buf)

	// server
	var serverepath = path.Join(apppath, "server")
	_ = os.MkdirAll(serverepath, os.ModePerm)
	write(serverepath, "config.go", model.ServerConfigFile, &buf)
	write(serverepath, "database.go", model.ServerDatabaseFile, &buf)
	write(serverepath, "engine.go", model.ServerEngineFile, &buf)
	write(serverepath, "logger.go", model.ServerLoggerFile, &buf)
	write(serverepath, "sentry.go", model.ServerSentryFile, &buf)

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
