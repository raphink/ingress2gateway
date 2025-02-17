/*
Copyright 2024 The Kubernetes Authors.

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

package extensions

import (
	gkegatewayv1 "github.com/GoogleCloudPlatform/gke-gateway-api/apis/networking/v1"
	"github.com/kubernetes-sigs/ingress2gateway/pkg/i2gw/intermediate"
)

func BuildBackendPolicySessionAffinityConfig(serviceIR intermediate.ProviderSpecificServiceIR) *gkegatewayv1.SessionAffinityConfig {
	affinityType := serviceIR.Gce.SessionAffinity.AffinityType
	saConfig := gkegatewayv1.SessionAffinityConfig{
		Type: &affinityType,
	}
	if affinityType == "GENERATED_COOKIE" {
		saConfig.CookieTTLSec = serviceIR.Gce.SessionAffinity.CookieTTLSec
	}
	return &saConfig
}

func BuildBackendPolicySecurityPolicyConfig(serviceIR intermediate.ProviderSpecificServiceIR) *string {
	securityPolicy := serviceIR.Gce.SecurityPolicy.Name
	return &securityPolicy
}
