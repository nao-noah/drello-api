package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/workspace"
)

func Update(ctx context.Context, workspaceRepo repository.Workspace, input *UpdateInput) (*UpdateOutput, error) {
	wdomain, err := workspaceRepo.Update(ctx, input.id, input.title)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Workspace: *wdomain}, nil
}

type UpdateInput struct {
	id    int
	title string
}

func NewUpdateInput(id int, title string) *UpdateInput {
	return &UpdateInput{id: id, title: title}
}

type UpdateOutput struct {
	Workspace workspace.Workspace
}
