/*
Copyright 2022-2024 EscherCloud.
Copyright 2024 the Unikorn Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstack

import (
	"errors"
)

var (
	// ErrResourceNotFound is returned when a named resource cannot
	// be looked up (we have to do it ourselves) and it cannot be found.
	ErrResourceNotFound = errors.New("requested resource not found")

	// ErrResourceDependency is returned when a resource is in unexpected
	// state or condition.
	ErrResouceDependency = errors.New("resource dependency error")
)
