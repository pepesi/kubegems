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

package tenanthandler

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"kubegems.io/kubegems/pkg/i18n"
	"kubegems.io/kubegems/pkg/utils/statistics"
)

func (h *TenantHandler) ValidateTenantResourceQuota(ctx context.Context, clustername string, clusterOversold []byte, origin, need []byte) error {
	originlist, newlist := v1.ResourceList{}, v1.ResourceList{}
	if err := json.Unmarshal(origin, &originlist); err != nil {
		return err
	}
	if err := json.Unmarshal(need, &newlist); err != nil {
		return err
	}
	cli, err := h.GetAgents().ClientOf(ctx, clustername)
	if err != nil {
		return err
	}
	statistics := &statistics.ClusterResourceStatistics{}
	if err := cli.Extend().ClusterResourceStatistics(ctx, statistics); err != nil {
		return i18n.Errorf(ctx, "validate tenant resoure quota failed, can't get cluster resource statistics: %w", err)
	}
	oversoldrates := ParseOversoldConfig(clusterOversold)
	if err := CheckOverSold(*statistics, oversoldrates, originlist, newlist); err != nil {
		return i18n.Errorf(ctx, "validate tenant resoure quota failed: %w", err)
	}
	return nil
}

// CheckOverSold
// To check:  ((Capacity * OversoldRate) - (AllTenantAllocated - CurrentTenantAllocated))  >  NewTenantAllocated
func CheckOverSold(clusterstatistics statistics.ClusterResourceStatistics, oversoldRate map[v1.ResourceName]float32, before, after v1.ResourceList) error {
	allAllocated := clusterstatistics.TenantAllocated.DeepCopy()

	oversoldCapacity := v1.ResourceList{}
	statistics.ResourceListCollect(oversoldCapacity, clusterstatistics.Capacity.DeepCopy(),
		func(resname v1.ResourceName, into *resource.Quantity, val resource.Quantity) {
			into.Set(int64(float32(val.Value()) * oversoldRate[resname]))
		})
	msgs := []string{}
	for resourceName, quantity := range after {
		if oversoldRate[resourceName] == 0 {
			continue
		}
		capacity := oversoldCapacity[resourceName].DeepCopy()
		if capacity.IsZero() {
			continue
		}

		available := Sub(capacity, Sub(allAllocated[resourceName], before[resourceName]))
		if quantity.Cmp(available) == 1 /*quantity > available*/ {
			msgs = append(msgs, fmt.Sprintf("resource [%s] available %s ,but request %s", resourceName, available.String(), quantity.String()))
		}
	}
	if len(msgs) > 0 {
		return fmt.Errorf("%v", msgs)
	}
	return nil
}

func Sub(a, b resource.Quantity) resource.Quantity {
	a = a.DeepCopy()
	a.Sub(b)
	return a
}
