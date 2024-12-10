// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
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
