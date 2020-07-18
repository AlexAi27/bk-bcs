/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// NetAgentOption option for net agent
type NetAgentOption struct {
	conf.ServiceConfig
	conf.MetricConfig
	conf.LogConfig
	conf.FileConfig

	Cluster              string `json:"cluster" value:"" usage:"cluster for bcs"`
	Kubeconfig           string `json:"kubeconfig" value:"" usage:"kubeconfig for kube-apiserver, Only required if out-of-cluster."`
	KubeResyncPeriod     int    `json:"kube_resync_peried" value:"300" usage:"resync interval for informer factory in seconds; (default 300)"`
	KubeCacheSyncTimeout int    `json:"kube_cachesync_timeout" value:"10" usage:"wait for kube cache sync timeout in seconds; (default 10)"`
	CheckInterval        int    `json:"check_interval" value:"300" usage:"interval for checking ip rules and route tables"`

	CloudNetserviceEndpoints string `json:"cloud_netservice_endpoints" value:"" usage:"cloud netservice endpoints, split by comma"`

	EniMTU int    `json:"eni_mtu" value:"1500" usage:"the mtu of eni"`
	Ifaces string `json:"ifaces" value:"eth1" usage:"use ip of these network interfaces as node identity, split with comma or semicolon"`
}

// New new option
func New() *NetAgentOption {
	return &NetAgentOption{}
}

// Parse parse options
func Parse(opt *NetAgentOption) {
	conf.Parse(opt)

	if len(opt.Cluster) == 0 {
		blog.Fatal("cluster cannot be empty")
	}
	if len(opt.CloudNetserviceEndpoints) == 0 {
		blog.Fatal("cloud netservice endpoints cannot be empty")
	}

	blog.Infof("get option %+v", opt)
}
