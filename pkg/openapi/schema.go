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
