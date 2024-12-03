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
	"AQAA//8G1nU2/voAAA==",
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
