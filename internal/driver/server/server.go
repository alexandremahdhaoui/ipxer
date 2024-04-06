package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/alexandremahdhaoui/ipxer/internal/controller"
	"github.com/alexandremahdhaoui/ipxer/internal/types"
	"github.com/labstack/echo/v4"
)

var (
	ErrGetIPXEBoostrap    = errors.New("getting ipxe bootstrap")
	ErrGetConfigByID      = errors.New("getting config by id")
	ErrGetIPXEBySelectors = errors.New("getting ipxe by labels")
)

func New(ipxe controller.IPXE, config controller.Config) ServerInterface {
	return &server{
		ipxe:   ipxe,
		config: config,
	}
}

type server struct {
	ipxe   controller.IPXE
	config controller.Config
}

func (s *server) GetIPXEBootstrap(c echo.Context) error {
	// call controller
	b := s.ipxe.Boostrap()

	// write response
	if _, err := c.Response().Write(b); err != nil {
		return writeErr(c, 500, errors.Join(err, ErrGetIPXEBoostrap))
	}

	c.Response().Status = 200

	return nil
}

func (s *server) GetConfigByID(c echo.Context, profileName string, configID UUID, _ GetConfigByIDParams) error {
	// call controller
	b, err := s.config.GetByID(context.Background(), profileName, configID)
	if err != nil {
		return writeErr(c, 500, errors.Join(err, ErrGetConfigByID))
	}

	// write response
	if _, err := c.Response().Write(b); err != nil {
		return writeErr(c, 500, errors.Join(err, ErrGetConfigByID))
	}

	c.Response().Status = 200

	return nil
}

func (s *server) GetIPXEBySelectors(c echo.Context, params GetIPXEBySelectorsParams) error {
	// convert into type
	//TODO: use params instead of converting the echo context?
	selectors, err := types.NewIpxeSelectorsFromContext(c)
	if err != nil {
		return writeErr(c, 500, errors.Join(err, ErrGetIPXEBySelectors))
	}

	// call controller
	b, err := s.ipxe.FindProfileAndRender(context.Background(), selectors)
	if err != nil {
		return writeErr(c, 500, errors.Join(err, ErrGetIPXEBySelectors))
	}

	// write response
	if _, err := c.Response().Write(b); err != nil {
		return writeErr(c, 500, errors.Join(err, ErrGetIPXEBySelectors))
	}

	c.Response().Status = 200

	return nil
}

func writeErr(c echo.Context, code int, err error) error {
	b, marshallErr := json.Marshal(Error{
		Code:    int32(code),
		Message: err.Error(),
	})
	if marshallErr != nil {
		c.Response().Status = code
		b = []byte(marshallErr.Error())
		err = errors.Join(marshallErr, err)
	}

	_, writeErr := c.Response().Write(b)
	if writeErr != nil {
		c.Response().Status = code
		err = errors.Join(writeErr, err)
	}

	c.Response().Status = code

	return err
}
