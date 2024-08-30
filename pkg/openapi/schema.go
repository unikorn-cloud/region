// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/fWfJ9L9fUj37u5/AwAA//+PuvDR0nsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
