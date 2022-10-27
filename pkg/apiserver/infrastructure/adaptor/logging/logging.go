package adaptor

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/options"
)

type PrometheusAdaptor struct {
}

func (p *PrometheusAdaptor) Labels(ctx context.Context, opts ...options.Option) []string

func (p *PrometheusAdaptor) Logs(ctx context.Context, opts ...options.Option) []string
