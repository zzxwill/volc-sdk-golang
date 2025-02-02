package veen

import (
	"net/url"
	"strconv"
)

func (v *Veen) CreateCloudServer(req *CreateCloudServerReq) (*CreateCloudServerResp, error) {
	resp := &CreateCloudServerResp{}
	if err := v.post("CreateCloudServer", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ListCloudServers(req *ListCloudServersReq) (*ListCloudServersResp, error) {
	resp := &ListCloudServersResp{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(int(req.Page)))
	query.Set("limit", strconv.Itoa(int(req.Limit)))
	query.Set("order_by", strconv.Itoa(int(req.OrderBy)))
	if req.FuzzyName != "" {
		query.Set("fuzzy_name", req.FuzzyName)
	}
	if err := v.get("ListCloudServers", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) GetCloudServer(req *GetCloudServerReq) (*GetCloudServerResp, error) {
	resp := &GetCloudServerResp{}
	query := url.Values{}
	query.Set("cloud_server_id", req.CloudServerID)
	if err := v.get("GetCloudServer", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) StartCloudServer(req *StartCloudServerReq) (*StartCloudServerResp, error) {
	resp := &StartCloudServerResp{}
	if err := v.post("StartCloudServer", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) StopCloudServer(req *StopCloudServerReq) (*StopCloudServerResp, error) {
	resp := &StopCloudServerResp{}
	if err := v.post("StopCloudServer", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) RebootCloudServer(req *RebootCloudServerReq) (*RebootCloudServerResp, error) {
	resp := &RebootCloudServerResp{}
	if err := v.post("RebootCloudServer", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) DeleteCloudServer(req *DeleteCloudServerReq) (*DeleteCloudServerResp, error) {
	resp := &DeleteCloudServerResp{}
	if err := v.post("DeleteCloudServer", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ListInstanceTypes(req *ListInstanceTypesReq) (*ListInstanceTypesResp, error) {
	resp := &ListInstanceTypesResp{}
	query := url.Values{}
	if err := v.get("ListInstanceTypes", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ListAvailableResourceInfo(req *ListAvailableResourceInfoReq) (*ListAvailableResourceInfoResp, error) {
	resp := &ListAvailableResourceInfoResp{}
	query := url.Values{}
	if req.InstanceType != "" {
		query.Set("instance_type", req.InstanceType)
	}
	if req.CloudDiskType != "" {
		query.Set("cloud_disk_type", req.CloudDiskType)
	}
	if err := v.get("ListAvailableResourceInfo", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) CreateInstance(req *CreateInstanceReq) (*CreateInstanceResp, error) {
	resp := &CreateInstanceResp{}
	if err := v.post("CreateInstance", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ListInstances(req *ListInstancesReq) (*ListInstancesResp, error) {
	resp := &ListInstancesResp{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(int(req.Page)))
	query.Set("limit", strconv.Itoa(int(req.Limit)))
	query.Set("order_by", strconv.Itoa(int(req.OrderBy)))
	if len(req.Countries) != 0 {
		query.Set("countries", req.Countries)
	}
	if len(req.Regions) != 0 {
		query.Set("regions", req.Regions)
	}
	if len(req.Cities) != 0 {
		query.Set("cities", req.Cities)
	}
	if len(req.Isps) != 0 {
		query.Set("isps", req.Isps)
	}
	if len(req.ClusterNames) != 0 {
		query.Set("cluster_names", req.ClusterNames)
	}
	if len(req.CloudServerIdentities) != 0 {
		query.Set("cloud_server_identities", req.CloudServerIdentities)
	}
	if len(req.Status) != 0 {
		query.Set("status", req.Status)
	}
	if len(req.SpecNames) != 0 {
		query.Set("spec_names", req.SpecNames)
	}
	if len(req.InstanceUuids) != 0 {
		query.Set("instance_uuids", req.InstanceUuids)
	}
	if len(req.InstanceNames) != 0 {
		query.Set("instance_names", req.InstanceNames)
	}
	if len(req.Ips) != 0 {
		query.Set("ips", req.Ips)
	}
	if len(req.Cidrs) != 0 {
		query.Set("cidrs", req.Cidrs)
	}
	if len(req.InstanceIdentities) != 0 {
		query.Set("instance_identities", req.InstanceIdentities)
	}
	if len(req.VPCIdentities) != 0 {
		query.Set("vpc_identities", req.VPCIdentities)
	}
	if len(req.FuzzyIPWithDots) != 0 {
		query.Set("fuzzy_ip_with_dots", req.FuzzyIPWithDots)
	}

	if err := v.get("ListInstances", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) GetInstance(req *GetInstanceReq) (*GetInstanceResp, error) {
	resp := &GetInstanceResp{}
	query := url.Values{}
	if req.InstanceIdentity != "" {
		query.Set("instance_identity", req.InstanceIdentity)
	}
	if err := v.get("GetInstance", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) StartInstances(req *StartInstancesReq) (*StartInstancesResp, error) {
	resp := &StartInstancesResp{}
	if err := v.post("StartInstances", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) StopInstances(req *StopInstancesReq) (*StopInstancesResp, error) {
	resp := &StopInstancesResp{}
	if err := v.post("StopInstances", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) RebootInstances(req *RebootInstancesReq) (*RebootInstancesResp, error) {
	resp := &RebootInstancesResp{}
	if err := v.post("RebootInstances", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) OfflineInstances(req *OfflineInstancesReq) (*OfflineInstancesResp, error) {
	resp := &OfflineInstancesResp{}
	if err := v.post("OfflineInstances", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) SetInstanceName(req *SetInstanceNameReq) (*SetInstanceNameResp, error) {
	resp := &SetInstanceNameResp{}
	if err := v.post("SetInstanceName", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ResetLoginCredential(req *ResetLoginCredentialReq) (*ResetLoginCredentialResp, error) {
	resp := &ResetLoginCredentialResp{}
	if err := v.post("ResetLoginCredential", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) GetInstanceCloudDiskInfo(req *GetInstanceCloudDiskInfoReq) (*GetInstanceCloudDiskInfoResp, error) {
	resp := &GetInstanceCloudDiskInfoResp{}
	query := url.Values{}
	if req.InstanceIdentity != "" {
		query.Set("instance_identity", req.InstanceIdentity)
	}
	if err := v.get("GetInstanceCloudDiskInfo", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ScaleInstanceCloudDiskCapacity(req *ScaleInstanceCloudDiskCapacityReq) (*ScaleInstanceCloudDiskCapacityResp, error) {
	resp := &ScaleInstanceCloudDiskCapacityResp{}
	if err := v.post("ScaleInstanceCloudDiskCapacity", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
