package services

import (
	"context"
	"database/sql"
	"errors"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/pkg/api"
)

type PanelServer struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewPanelServer(db *sql.DB) *PanelServer {
	return &PanelServer{
		db:      db,
		queries: queries.New(db),
	}
}

func (s *PanelServer) GetPanelById(ctx context.Context, id int) (*queries.Panel, error) {
	panel, err := s.queries.GetPanelByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api.NotFoundError("未找到面板")
		}
		return nil, err
	}
	return panel, nil
}

func (s *PanelServer) ListPanels(ctx context.Context, input api.PaginationParams) ([]*queries.Panel, error) {
	queryParams := queries.ListPanelsParams{Limit: input.Current, Offset: input.Offset()}
	panels, err := s.queries.ListPanels(ctx, queryParams)
	if err != nil {
		return nil, err
	}
	return panels, nil
}

func (s *PanelServer) CreatePanel(ctx context.Context, input queries.CreatePanelParams) (*queries.Panel, error) {
	return s.queries.CreatePanel(ctx, input)
}

func (s *PanelServer) UpdatePanel(ctx context.Context, input queries.UpdatePanelParams) error {
	return s.queries.UpdatePanel(ctx, input)
}

func (s *PanelServer) DeletePanel(ctx context.Context, id int) error {
	return s.queries.DeletePanelByID(ctx, id)
}
