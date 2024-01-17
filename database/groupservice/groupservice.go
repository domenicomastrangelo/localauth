package groupservice

import (
	"context"
	"localauth/database/group"
	"localauth/database/grouprepository"
	"strings"
)

type Service interface {
	Authorize(ctx context.Context, group *group.Group) error
	AddGroup(ctx context.Context, group *group.Group) error
	ValidateGroup(group *group.Group) error
	GetGroups(ctx context.Context) (*[]group.Group, error)
	GetGroup(ctx context.Context, id int) (*group.Group, error)
}

type ServiceImpl struct {
	GroupRepository grouprepository.Repository
}

func New(groupRepository grouprepository.Repository) Service {
	return &ServiceImpl{
		GroupRepository: groupRepository,
	}
}

func (g *ServiceImpl) Authorize(ctx context.Context, group *group.Group) error {
	return nil
}

func (g *ServiceImpl) AddGroup(ctx context.Context, group *group.Group) error {
	if err := g.ValidateGroup(group); err != nil {
		return err
	}
	return g.GroupRepository.AddGroup(group, ctx, group.ID)
}

func (g *ServiceImpl) ValidateGroup(group *group.Group) error {
	if strings.TrimSpace(group.Name) == "" {
		return grouprepository.ErrInvalidGroupName
	}

	return nil
}

func (g *ServiceImpl) GetGroups(ctx context.Context) (*[]group.Group, error) {
	return g.GroupRepository.GetGroups(ctx)
}

func (g *ServiceImpl) GetGroup(ctx context.Context, id int) (*group.Group, error) {
	return g.GroupRepository.GetGroup(ctx, id)
}
