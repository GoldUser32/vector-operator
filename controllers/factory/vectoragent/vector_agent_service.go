/*
Copyright 2022.

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

package vectoragent

import (
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/util/intstr"
)

func (ctrl *Controller) createVectorAgentService() *corev1.Service {
	labels := ctrl.labelsForVectorAgent()

	service := &corev1.Service{
		ObjectMeta: ctrl.objectMetaVectorAgent(labels),
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "prom-exporter",
					Protocol:   corev1.Protocol("TCP"),
					Port:       9090,
					TargetPort: intstr.FromInt(9090),
				},
			},
			Selector: labels,
		},
	}
	return service
}
