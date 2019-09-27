/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	v2 "bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"
	"bk-bcs/bcs-mesos/pkg/client/internalclientset/scheme"

	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type BkbcsV2Interface interface {
	RESTClient() rest.Interface
	AdmissionWebhookConfigurationsGetter
	AgentsGetter
	AgentSchedInfosGetter
	ApplicationsGetter
	BcsClusterAgentSettingsGetter
	BcsCommandInfosGetter
	BcsConfigMapsGetter
	BcsEndpointsGetter
	BcsSecretsGetter
	BcsServicesGetter
	CrdsGetter
	CrrsGetter
	DeploymentsGetter
	TasksGetter
	TaskGroupsGetter
	VersionsGetter
}

// BkbcsV2Client is used to interact with features provided by the bkbcs.bkbcs.tencent.com group.
type BkbcsV2Client struct {
	restClient rest.Interface
}

func (c *BkbcsV2Client) AdmissionWebhookConfigurations(namespace string) AdmissionWebhookConfigurationInterface {
	return newAdmissionWebhookConfigurations(c, namespace)
}

func (c *BkbcsV2Client) Agents(namespace string) AgentInterface {
	return newAgents(c, namespace)
}

func (c *BkbcsV2Client) AgentSchedInfos(namespace string) AgentSchedInfoInterface {
	return newAgentSchedInfos(c, namespace)
}

func (c *BkbcsV2Client) Applications(namespace string) ApplicationInterface {
	return newApplications(c, namespace)
}

func (c *BkbcsV2Client) BcsClusterAgentSettings(namespace string) BcsClusterAgentSettingInterface {
	return newBcsClusterAgentSettings(c, namespace)
}

func (c *BkbcsV2Client) BcsCommandInfos(namespace string) BcsCommandInfoInterface {
	return newBcsCommandInfos(c, namespace)
}

func (c *BkbcsV2Client) BcsConfigMaps(namespace string) BcsConfigMapInterface {
	return newBcsConfigMaps(c, namespace)
}

func (c *BkbcsV2Client) BcsEndpoints(namespace string) BcsEndpointInterface {
	return newBcsEndpoints(c, namespace)
}

func (c *BkbcsV2Client) BcsSecrets(namespace string) BcsSecretInterface {
	return newBcsSecrets(c, namespace)
}

func (c *BkbcsV2Client) BcsServices(namespace string) BcsServiceInterface {
	return newBcsServices(c, namespace)
}

func (c *BkbcsV2Client) Crds(namespace string) CrdInterface {
	return newCrds(c, namespace)
}

func (c *BkbcsV2Client) Crrs(namespace string) CrrInterface {
	return newCrrs(c, namespace)
}

func (c *BkbcsV2Client) Deployments(namespace string) DeploymentInterface {
	return newDeployments(c, namespace)
}

func (c *BkbcsV2Client) Tasks(namespace string) TaskInterface {
	return newTasks(c, namespace)
}

func (c *BkbcsV2Client) TaskGroups(namespace string) TaskGroupInterface {
	return newTaskGroups(c, namespace)
}

func (c *BkbcsV2Client) Versions(namespace string) VersionInterface {
	return newVersions(c, namespace)
}

// NewForConfig creates a new BkbcsV2Client for the given config.
func NewForConfig(c *rest.Config) (*BkbcsV2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BkbcsV2Client{client}, nil
}

// NewForConfigOrDie creates a new BkbcsV2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BkbcsV2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BkbcsV2Client for the given RESTClient.
func New(c rest.Interface) *BkbcsV2Client {
	return &BkbcsV2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BkbcsV2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
