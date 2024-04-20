package service

import (
	"context"
	org "github.com/go-atreus/atreus-server/app/admin/api/organization"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

var _ org.OrganizationServer = (*OrganizationService)(nil)

type OrganizationService struct {
	org.OrganizationDefaultServer
}

func (m *OrganizationService) QueryOrganization(ctx context.Context, organization *org.SysOrganization) (*org.SysOrganization, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrganization(logger log.Logger, data *data.Data) *OrganizationService {
	return &OrganizationService{
		OrganizationDefaultServer: org.OrganizationDefaultServer{
			DB: data.ORM,
		},
	}
}

func (m *OrganizationService) OrganizationTree(ctx context.Context, in *org.SysOrganization) (res *org.ListSysOrganization, err error) {
	out := &org.ListSysOrganization{}
	tree, err := m.OrganizationChildren(ctx, &org.SysOrganization{Id: 0})
	if err != nil {
		return nil, err
	}
	out.Results = tree
	return out, nil
}

func (m *OrganizationService) OrganizationChildren(ctx context.Context, in *org.SysOrganization) (res []*org.SysOrganization, err error) {
	var orgs []org.SysOrganizationORM
	if err := m.DB.Where("parent_id=?", in.Id).Find(&orgs).Error; err != nil {
		return nil, err
	}
	var pbResponse []*org.SysOrganization

	for _, responseEntry := range orgs {
		children, _ := m.OrganizationChildren(ctx, &org.SysOrganization{Id: responseEntry.Id})
		var pb org.SysOrganization
		pb, err = responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pb.Children = children
		pbResponse = append(pbResponse, &pb)
	}

	res = pbResponse
	return
}
