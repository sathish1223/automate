package infra_proxy

import (
	"context"

	gwreq "github.com/chef/automate/api/external/infra_proxy/request"
	gwres "github.com/chef/automate/api/external/infra_proxy/response"
	infra_req "github.com/chef/automate/api/interservice/infra_proxy/request"
	infra_res "github.com/chef/automate/api/interservice/infra_proxy/response"
)

// GetOrgs fetches an array of existing orgs
func (a *InfraProxyServer) GetOrgs(ctx context.Context, r *gwreq.GetOrgs) (*gwres.GetOrgs, error) {
	req := &infra_req.GetOrgs{
		ServerId: r.ServerId,
	}
	res, err := a.client.GetOrgs(ctx, req)
	if err != nil {
		return nil, err
	}

	return &gwres.GetOrgs{
		Orgs: fromUpstreamOrgs(res.Orgs),
	}, nil
}

// GetOrg fetches a single org by ID
func (a *InfraProxyServer) GetOrg(ctx context.Context, r *gwreq.GetOrg) (*gwres.GetOrg, error) {
	req := &infra_req.GetOrg{
		Id:       r.Id,
		ServerId: r.ServerId,
	}
	res, err := a.client.GetOrg(ctx, req)
	if err != nil {
		return nil, err
	}
	return &gwres.GetOrg{
		Org: fromUpstreamOrg(res.Org),
	}, nil
}

// CreateOrg posts an org upstream
func (a *InfraProxyServer) CreateOrg(ctx context.Context, r *gwreq.CreateOrg) (*gwres.CreateOrg, error) {
	req := &infra_req.CreateOrg{
		Id:        r.Id,
		Name:      r.Name,
		AdminUser: r.AdminUser,
		AdminKey:  r.AdminKey,
		ServerId:  r.ServerId,
		Projects:  r.Projects,
	}
	res, err := a.client.CreateOrg(ctx, req)
	if err != nil {
		return nil, err
	}
	return &gwres.CreateOrg{
		Org: fromUpstreamOrg(res.Org),
	}, nil
}

// UpdateOrg updates an org upstream
func (a *InfraProxyServer) UpdateOrg(ctx context.Context, r *gwreq.UpdateOrg) (*gwres.UpdateOrg, error) {
	req := &infra_req.UpdateOrg{
		Id:        r.Id,
		Name:      r.Name,
		AdminUser: r.AdminUser,
		ServerId:  r.ServerId,
		Projects:  r.Projects,
	}
	res, err := a.client.UpdateOrg(ctx, req)
	if err != nil {
		return nil, err
	}
	return &gwres.UpdateOrg{
		Org: fromUpstreamOrg(res.Org),
	}, nil
}

// ResetOrgAdminKey resets the org admin key
func (a *InfraProxyServer) ResetOrgAdminKey(ctx context.Context, r *gwreq.ResetOrgAdminKey) (*gwres.ResetOrgAdminKey, error) {
	req := &infra_req.ResetOrgAdminKey{
		Id:       r.Id,
		ServerId: r.ServerId,
		AdminKey: r.AdminKey,
	}
	res, err := a.client.ResetOrgAdminKey(ctx, req)
	if err != nil {
		return nil, err
	}

	return &gwres.ResetOrgAdminKey{
		Org: fromUpstreamOrg(res.Org),
	}, nil
}

// DeleteOrg deletes an org upstream
func (a *InfraProxyServer) DeleteOrg(ctx context.Context, r *gwreq.DeleteOrg) (*gwres.DeleteOrg, error) {
	req := &infra_req.DeleteOrg{
		Id:       r.Id,
		ServerId: r.ServerId,
	}
	res, err := a.client.DeleteOrg(ctx, req)
	if err != nil {
		return nil, err
	}
	return &gwres.DeleteOrg{
		Org: fromUpstreamOrg(res.Org),
	}, nil
}

//GetInfraServerOrgs: Fetches the list of automate infra server's organisations from the chef server and save it into the automate back end DB
func (c *InfraProxyServer) GetInfraServerOrgs(ctx context.Context, r *gwreq.GetInfraServerOrgs) (*gwres.GetInfraServerOrgs, error) {
	req := &infra_req.GetInfraServerOrgs{
		ServerId: r.ServerId,
	}
	res, err := c.client.GetInfraServerOrgs(ctx, req)
	if err != nil {
		return nil, err
	}
	return &gwres.GetInfraServerOrgs{
		MigrationId: res.MigrationId,
	}, nil
}

//CreateInfraServerOrgs: Creates the organisation on the chef server
func (c *InfraProxyServer) CreateInfraServerOrgs(ctx context.Context, r *gwreq.CreateInfraServerOrgs) (*gwres.CreateInfraServerOrgs, error) {
	req := &infra_req.CreateInfraServerOrgs{
		ServerId: r.ServerId,
		Name:     r.Name,
		FullName: r.FullName,
	}
	_, err := c.client.CreateInfraServerOrgs(ctx, req)
	if err != nil {
		return nil, err
	}
	return &gwres.CreateInfraServerOrgs{}, nil
}

func fromUpstreamOrg(t *infra_res.Org) *gwres.Org {
	return &gwres.Org{
		Id:           t.GetId(),
		Name:         t.GetName(),
		AdminUser:    t.GetAdminUser(),
		CredentialId: t.GetCredentialId(),
		ServerId:     t.GetServerId(),
		Projects:     t.GetProjects(),
	}
}

func fromUpstreamOrgs(orgs []*infra_res.Org) []*gwres.Org {
	ts := make([]*gwres.Org, len(orgs))

	for i, org := range orgs {
		ts[i] = fromUpstreamOrg(org)
	}

	return ts
}
