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

package agents

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/rest"
	"kubegems.io/kubegems/pkg/utils/kube"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	AgentModeApiServer = "apiServerProxy"
	AgentModeAHTTP     = "http"
	AgentModeHTTPS     = "https"
)

type Client interface {
	client.WithWatch
	DoRequest(ctx context.Context, req Request) error
	DoRawRequest(ctx context.Context, clientreq Request) (*http.Response, error)
	DialWebsocket(ctx context.Context, path string, headers http.Header) (*websocket.Conn, *http.Response, error)
	Extend() *ExtendClient
	Name() string
	BaseAddr() url.URL
	APIServerAddr() url.URL
	APIServerVersion() string
	RestConfig() *rest.Config
	// Deprecated: remove
	Proxy(ctx context.Context, obj client.Object, port int, req *http.Request, writer http.ResponseWriter, rewritefunc func(r *http.Response) error) error
}

var _ Client = &DelegateClient{}

func NewDelegateClientClient(options *ClientOptions, name string, apiserver *url.URL, discovery discovery.DiscoveryInterface, tracer trace.Tracer) Client {
	cli := NewTypedClient(options, kube.GetScheme())
	return &DelegateClient{
		name:            name,
		apiserverAddr:   apiserver,
		baseaddr:        options.Addr,
		discovery:       memory.NewMemCacheClient(discovery),
		TypedClient:     cli,
		ExtendClient:    NewExtendClientFrom(cli),
		WebsocketClient: NewWebsocketClient(options),
	}
}

type DelegateClient struct {
	*TypedClient
	*ExtendClient
	*WebsocketClient
	name          string
	baseaddr      *url.URL
	apiserverAddr *url.URL
	discovery     discovery.DiscoveryInterface
	kubeconfig    *rest.Config
}

func (c *DelegateClient) Extend() *ExtendClient {
	return c.ExtendClient
}

func (c *DelegateClient) Name() string {
	return c.name
}

func (c *DelegateClient) BaseAddr() url.URL {
	return *c.baseaddr
}

func (c *DelegateClient) APIServerAddr() url.URL {
	return *c.apiserverAddr
}

func (c *DelegateClient) RestConfig() *rest.Config {
	return c.kubeconfig
}

func (c *DelegateClient) APIServerVersion() string {
	version, err := c.discovery.ServerVersion()
	if err != nil {
		return ""
	}
	return version.String()
}
