// Copyright 2022 The kubegems.io Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package response

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func OK(w http.ResponseWriter, data interface{}) {
	Raw(w, http.StatusOK, Response{Data: data}, nil)
}

func NotFound(w http.ResponseWriter, message string) {
	Error(w, NewError(http.StatusNotFound, message))
}

func BadRequest(w http.ResponseWriter, message string) {
	Error(w, StatusError{Status: http.StatusBadRequest, Message: message})
}

func ServerError(w http.ResponseWriter, err error) {
	Error(w, StatusError{Status: http.StatusInternalServerError, Message: err.Error()})
}

func Error(w http.ResponseWriter, err error) {
	serr := &StatusError{}
	if errors.As(err, &serr) {
		Raw(w, serr.Status, Response{Message: err.Error(), Error: err}, nil)
	} else {
		Raw(w, http.StatusBadRequest, Response{Message: err.Error(), Error: err}, nil)
	}
}

func Raw(w http.ResponseWriter, status int, data interface{}, headers map[string]string) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	switch val := data.(type) {
	case io.Reader:
		setContentTypeIfNotSet(w.Header(), "application/octet-stream")
		w.WriteHeader(status)
		_, _ = io.Copy(w, val)
	case string:
		setContentTypeIfNotSet(w.Header(), "text/plain")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(val))
	case []byte:
		setContentTypeIfNotSet(w.Header(), "application/octet-stream")
		w.WriteHeader(status)
		_, _ = w.Write(val)
	case nil:
		w.WriteHeader(status)
		// do not write a nil representation
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(data)
	}
}

func setContentTypeIfNotSet(hds http.Header, val string) {
	if val := hds.Get("Content-Type"); val == "" {
		hds.Set("Content-Type", val)
	}
}

type StatusError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e StatusError) Error() string {
	return e.Message
}

func NewError(status int, message string) *StatusError {
	return &StatusError{Status: status, Message: message}
}
