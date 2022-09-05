package apis

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"kubegems.io/kubegems/pkg/agent/cluster"
	"kubegems.io/kubegems/pkg/log"
)

type ClusterHandler struct {
	cluster        cluster.Interface
	lastUpdateTime *time.Time
	rwmu           sync.RWMutex
	cacheData      []*v1.APIResourceList
}

// @Tags         Agent.V1
// @Summary      获取k8s api-resources
// @Description  获取k8s api-resources
// @Accept       json
// @Produce      json
// @Param        cluster  path      string                                  true  "cluster"
// @Success      200      {object}  handlers.ResponseStruct{Data=[]object}  "resp"
// @Router       /v1/proxy/cluster/{cluster}/api-resources [get]
func (h *ClusterHandler) APIResources(c *gin.Context) {
	ret, refreshed, err := h.getApiResources()
	if refreshed {
		h.rwmu.Lock()
		defer h.rwmu.Unlock()
		h.cacheData = ret
	}
	if err != nil {
		if discovery.IsGroupDiscoveryFailedError(err) {
			OK(c, ret)
			return
		} else {
			NotOK(c, err)
			return
		}
	}
	OK(c, ret)
}

func (h *ClusterHandler) getApiResources() (ret []*v1.APIResourceList, refreshed bool, err error) {
	if time.Since(*h.lastUpdateTime) > time.Duration(time.Minute*5) && len(h.cacheData) > 0 {
		ret = h.cacheData
		return
	}
	ret, err = h.cluster.Discovery().ServerPreferredResources()
	if err != nil {
		log.Error(err, "failed to refresh api-resources")
		return
	}
	refreshed = true
	return
}
