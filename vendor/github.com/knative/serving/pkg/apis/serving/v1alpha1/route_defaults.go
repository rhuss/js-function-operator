/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	"context"

	"github.com/knative/serving/pkg/apis/serving/v1beta1"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/ptr"
)

func (r *Route) SetDefaults(ctx context.Context) {
	r.Spec.SetDefaults(apis.WithinSpec(ctx))
}

func (rs *RouteSpec) SetDefaults(ctx context.Context) {
	if v1beta1.IsUpgradeViaDefaulting(ctx) {
		beta := v1beta1.RouteSpec{}
		if rs.ConvertUp(ctx, &beta) == nil {
			alpha := RouteSpec{}
			alpha.ConvertDown(ctx, beta)
			*rs = alpha
		}
	}

	if len(rs.Traffic) == 0 && v1beta1.HasDefaultConfigurationName(ctx) {
		rs.Traffic = []TrafficTarget{{
			TrafficTarget: v1beta1.TrafficTarget{
				Percent:        100,
				LatestRevision: ptr.Bool(true),
			},
		}}
	}

	for i := range rs.Traffic {
		rs.Traffic[i].SetDefaults(ctx)
	}
}
