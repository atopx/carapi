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

func Execute(name, output, remote string, vendor, docker, compose bool) (err error) {
	// root dir
	var rootpath = path.Join(output, name)
	if err = os.MkdirAll(rootpath, os.ModePerm); err != nil {
		return err
	}
	// git
	cmd := exec.Command("git", "init")
	cmd.Dir = rootpath
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("git init failed: %v", err)
	}
	cmd.Wait()
	if remote != "" {
		cmd = exec.Command("git", "remote", "add", "origin", remote)
		cmd.Dir = rootpath
		if err = cmd.Start(); err != nil {
			return fmt.Errorf("git add remote error: %v", err)
		}
		cmd.Wait()
	}

	var buf bytes.Buffer

	// app
	var apppath = path.Join(rootpath, "services/app")
	os.MkdirAll(apppath, os.ModePerm)
	write(apppath, "build.dockerfile", model.BuildDockerFile, &buf)
	write(apppath, "local.dockerfile", model.LocalDockerFile, &buf)

	// api
	var apipath = path.Join(rootpath, "services/app/api")
	os.MkdirAll(apipath, os.ModePerm)
	write(apipath, "test.go", model.ApiTestFile, &buf)

	// config
	var cfgpath = path.Join(rootpath, "services/app/config")
	os.MkdirAll(cfgpath, os.ModePerm)
	write(cfgpath, "config.go", model.ConfigGoFile, &buf)
	write(apppath, "config.yml", model.ConfigYamlFile, &buf)

	// core
	var corepath = path.Join(rootpath, "services/app/core")
	os.MkdirAll(corepath, os.ModePerm)

	// docs
	var docpath = path.Join(rootpath, "services/app/docs")
	os.MkdirAll(docpath, os.ModePerm)
	write(docpath, "docs.go", model.DocGoFile, &buf)
	write(docpath, "swagger.json", model.DocSwaggerJsonFile, &buf)
	write(docpath, "swagger.yaml", model.DocSwaggerYamlFile, &buf)

	// libs
	var libpath = path.Join(rootpath, "services/app/libs")
	os.MkdirAll(libpath, os.ModePerm)
	write(libpath, "org_logger.go", model.LibOrgLoggerFile, &buf)

	// middleware
	var midpath = path.Join(rootpath, "services/app/middleware")
	os.MkdirAll(midpath, os.ModePerm)
	write(midpath, "cors.go", model.MiddlewareCorsFile, &buf)
	write(midpath, "logger.go", model.MiddlewareLoggerFile, &buf)
	write(midpath, "swagger.go", model.MiddlewareSwaggerFile, &buf)

	// models
	var modelpath = path.Join(rootpath, "services/app/models")
	os.MkdirAll(modelpath, os.ModePerm)
	write(modelpath, "user.go", model.ModelUserFile, &buf)

	// public
	var pubpath = path.Join(rootpath, "services/app/public")
	os.MkdirAll(pubpath, os.ModePerm)
	write(pubpath, "const.go", model.PublicConstFile, &buf)
	write(pubpath, "variable.go", model.PublicVariableFile, &buf)

	// routes
	var routepath = path.Join(rootpath, "services/app/routes")
	os.MkdirAll(routepath, os.ModePerm)
	write(routepath, "test.go", model.RouterTestFile, &buf)

	// schemas
	var schemapath = path.Join(rootpath, "services/app/schemas")
	os.MkdirAll(schemapath, os.ModePerm)
	write(schemapath, "common.go", model.SchemaCommonFile, &buf)
	write(schemapath, "response.go", model.SchemaResponseFile, &buf)

	// server
	var serverepath = path.Join(rootpath, "services/app/server")
	os.MkdirAll(serverepath, os.ModePerm)
	write(serverepath, "config.go", model.ServerConfigFile, &buf)
	write(serverepath, "database.go", model.ServerDatabaseFile, &buf)
	write(serverepath, "engine.go", model.ServerEngineFile, &buf)
	write(serverepath, "logger.go", model.ServerLoggerFile, &buf)
	write(serverepath, "sentry.go", model.ServerSentryFile, &buf)

	// docker compose
	write(rootpath, "docker-compose.local.yaml", model.DockerComposeLocalFile, &buf)
	write(rootpath, "docker-compose.yaml", model.DockerComposeReleaseFile, &buf)

	log.Println("Success!, cd xxxxx && go run main.go")
	return nil
}

func write(p, n, v string, buf *bytes.Buffer) {
	buf.WriteString(v)
	ioutil.WriteFile(path.Join(p, n), buf.Bytes(), 0644)
	buf.Reset()
}
