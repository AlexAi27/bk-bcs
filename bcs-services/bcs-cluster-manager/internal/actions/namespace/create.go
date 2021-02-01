/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package namespace

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	cmproto "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/store"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/types"
)

// CreateAction action for create namespace
type CreateAction struct {
	ctx   context.Context
	model store.ClusterManagerModel
	req   *cmproto.CreateNamespaceReq
	resp  *cmproto.CreateNamespaceResp
}

// NewCreateAction create namespace action
func NewCreateAction(model store.ClusterManagerModel) *CreateAction {
	return &CreateAction{
		model: model,
	}
}

func (ca *CreateAction) validate() error {
	if err := ca.req.Validate(); err != nil {
		return err
	}
	return nil
}

func (ca *CreateAction) queryFederationCluster(clusterID string) error {
	_, err := ca.model.GetCluster(ca.ctx, clusterID)
	return err
}

func (ca *CreateAction) createNamespace() error {
	newNs := &types.Namespace{
		Name:                ca.req.Name,
		FederationClusterID: ca.req.FederationClusterID,
		ProjectID:           ca.req.ProjectID,
		BusinessID:          ca.req.BusinessID,
		Labels:              ca.req.Labels,
	}
	return ca.model.CreateNamespace(ca.ctx, newNs)
}

func (ca *CreateAction) setResp(code uint64, msg string) {
	ca.resp.Seq = ca.req.Seq
	ca.resp.ErrCode = code
	ca.resp.ErrMsg = msg
}

// Handle create namespace request
func (ca *CreateAction) Handle(ctx context.Context,
	req *cmproto.CreateNamespaceReq, resp *cmproto.CreateNamespaceResp) {
	if req == nil || resp == nil {
		blog.Errorf("create namespace failed, req or resp is empty")
		return
	}
	ca.ctx = ctx
	ca.req = req
	ca.resp = resp

	if err := ca.validate(); err != nil {
		ca.setResp(types.BcsErrClusterManagerInvalidParameter, err.Error())
		return
	}
	if len(ca.req.FederationClusterID) != 0 {
		if err := ca.queryFederationCluster(ca.req.FederationClusterID); err != nil {
			ca.setResp(types.BcsErrClusterManagerDBOperation, err.Error())
			return
		}
	}

	if err := ca.createNamespace(); err != nil {
		ca.setResp(types.BcsErrClusterManagerDBOperation, err.Error())
		return
	}
	ca.setResp(types.BcsErrClusterManagerSuccess, types.BcsErrClusterManagerSuccessStr)
	return
}