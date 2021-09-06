package datasource

import (
	"context"
	domainWorkspace "drello-api/pkg/domain/workspace"
	"drello-api/pkg/infrastracture/mysql"
	"fmt"
)

type Workspace struct{}

func (w Workspace) GetAll(ctx context.Context) (*[]*domainWorkspace.Workspace, error) {
	db := mysql.DBPool()
	rows, err := db.Query("SELECT * FROM workspaces")
	if err != nil {
		return nil, fmt.Errorf("failed querying workspaces: %w", err)
	}
	defer rows.Close()

	workspaces := []*domainWorkspace.Workspace{}

	for rows.Next() {
		var id int
		var title string

		err := rows.Scan(&id, &title)
		if err != nil {
			return nil, fmt.Errorf("failed querying workspaces: %w", err)
		}

		workspaces = append(workspaces, domainWorkspace.New(id, title))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed querying workspaces: %w", err)
	}

	return &workspaces, nil
}

func (w Workspace) GetOne(ctx context.Context, id int) (*domainWorkspace.Workspace, error) {
	db := mysql.DBPool()
	rows, err := db.Query("SELECT id, title FROM workspaces WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("failed querying workspaces: %w", err)
	}
	defer rows.Close()

	workspaces := []*domainWorkspace.Workspace{}

	for rows.Next() {
		var id int
		var title string

		err := rows.Scan(&id, &title)
		if err != nil {
			return nil, fmt.Errorf("failed querying workspaces: %w", err)
		}

		workspaces = append(workspaces, domainWorkspace.New(id, title))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed querying workspaces: %w", err)
	}

	if len := len(workspaces); len != 1 {
		return nil, fmt.Errorf("failed querying workspaces: found invalid number of rows: Expected 1. Got %d", len)
	}

	return workspaces[0], nil
}

func (w Workspace) Create(ctx context.Context, title string) (*domainWorkspace.Workspace, error) {
	// wNode, err := mysql.DBPool().Workspace.Create().SetTitle(title).Save(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed creating workspace: %w", err)
	// }

	// return workspace.New(wNode.ID, wNode.Title), nil

	return nil, nil
}

func (w Workspace) Update(ctx context.Context, id int, title string) (*domainWorkspace.Workspace, error) {
	// wNode, err := mysql.DBPool().Workspace.UpdateOneID(id).SetTitle(title).Save(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed updating workspace: %w", err)
	// }

	// return workspace.New(wNode.ID, wNode.Title), nil

	return nil, nil
}

func (w Workspace) Delete(ctx context.Context, id int) error {
	// err := mysql.DBPool().Workspace.DeleteOneID(id).Exec(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed deleting workspace: %w", err)
	// }

	// return nil

	return nil
}
