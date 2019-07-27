package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func writeSchema(root, contents, schemaName string) error {
	pkg := strings.ToLower(LittleQ(ToNonExported(schemaName)))
	dirPath := filepath.Join(root, pkg)
	schemaFile := filepath.Join(dirPath, "schema.go")

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	return ioutil.WriteFile(schemaFile, []byte(contents), os.ModePerm)
}

func writeTable(root, contents, schemaName, name string) error {
	dirPath := buildTableDir(root, schemaName, name)
	path := filepath.Join(dirPath, "table.go")

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	return ioutil.WriteFile(path, []byte(contents), os.ModePerm)
}

func writePackageMembers(root, contents, schemaName, name string) error {
	dirPath := buildTableDir(root, schemaName, name)
	path := filepath.Join(dirPath, "exported.go")

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	return ioutil.WriteFile(path, []byte(contents), os.ModePerm)
}

func buildTableDir(root string, schemaName string, name string) string {
	tablePkgName := strings.ToLower(LittleQ(ToNonExported(name)))
	schemaPkgName := strings.ToLower(LittleQ(ToNonExported(schemaName)))
	dirPath := filepath.Join(root, schemaPkgName, tablePkgName)
	return dirPath
}
