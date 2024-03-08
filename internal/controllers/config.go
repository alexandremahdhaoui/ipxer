package controllers

import (
	"context"
	"errors"
	"github.com/alexandremahdhaoui/ipxe-api/internal/adapter"
	"github.com/alexandremahdhaoui/ipxe-api/internal/types"
	"github.com/google/uuid"
)

// ---------------------------------------------------- INTERFACE --------------------------------------------------- //

type Config interface {
	GetByID(ctx context.Context, profileID, configID uuid.UUID) ([]byte, error)
}

// --------------------------------------------------- CONSTRUCTORS ------------------------------------------------- //

func NewConfig(profile adapter.Profile, mux ResolveTransformerMux) Config {
	return &config{
		profile: profile,
		mux:     mux,
	}
}

// ---------------------------------------------------- CONFIG ------------------------------------------------- //

type config struct {
	profile adapter.Profile
	mux     ResolveTransformerMux
}

func (c *config) GetByID(ctx context.Context, profileID, configID uuid.UUID) ([]byte, error) {
	if configID == uuid.Nil {
		return nil, errors.New("TODO") //TODO: err
	}

	profile, err := c.profile.FindByID(ctx, profileID)
	if err != nil {
		return nil, err //TODO: wrap me
	}

	//TODO(super ineffective): change me
	var (
		content types.Content
		found   bool
	)
	for _, ctt := range profile.AdditionalContent {
		if ctt.ExposedConfigID == content.ExposedConfigID {
			content = ctt
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("TODO") //TODO: err
	}

	res, err := c.mux.ResolveAndTransformBatch(ctx, []types.Content{content})
	if err != nil {
		return nil, err //TODO: wrap me
	}

	return res[content.Name], nil
}
