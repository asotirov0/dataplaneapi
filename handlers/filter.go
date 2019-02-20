package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/haproxytech/client-native"
	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/filter"
)

//CreateFilterHandlerImpl implementation of the CreateFilterHandler interface using client-native client
type CreateFilterHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent *haproxy.ReloadAgent
}

//DeleteFilterHandlerImpl implementation of the DeleteFilterHandler interface using client-native client
type DeleteFilterHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent *haproxy.ReloadAgent
}

//GetFilterHandlerImpl implementation of the GetFilterHandler interface using client-native client
type GetFilterHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//GetFiltersHandlerImpl implementation of the GetFiltersHandler interface using client-native client
type GetFiltersHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//ReplaceFilterHandlerImpl implementation of the ReplaceFilterHandler interface using client-native client
type ReplaceFilterHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent *haproxy.ReloadAgent
}

//Handle executing the request and returning a response
func (h *CreateFilterHandlerImpl) Handle(params filter.CreateFilterParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	err := h.Client.Configuration.CreateFilter(params.ParentType, params.ParentName, params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return filter.NewCreateFilterDefault(int(*e.Code)).WithPayload(e)
	}
	h.ReloadAgent.Reload()
	return filter.NewCreateFilterCreated().WithPayload(params.Data)
}

//Handle executing the request and returning a response
func (h *DeleteFilterHandlerImpl) Handle(params filter.DeleteFilterParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	err := h.Client.Configuration.DeleteFilter(params.ID, params.ParentType, params.ParentName, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return filter.NewDeleteFilterDefault(int(*e.Code)).WithPayload(e)
	}
	h.ReloadAgent.Reload()
	return filter.NewDeleteFilterNoContent()
}

//Handle executing the request and returning a response
func (h *GetFilterHandlerImpl) Handle(params filter.GetFilterParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	f, err := h.Client.Configuration.GetFilter(params.ID, params.ParentType, params.ParentName, t)
	if err != nil {
		e := misc.HandleError(err)
		return filter.NewGetFilterDefault(int(*e.Code)).WithPayload(e)
	}
	return filter.NewGetFilterOK().WithPayload(f)
}

//Handle executing the request and returning a response
func (h *GetFiltersHandlerImpl) Handle(params filter.GetFiltersParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	fs, err := h.Client.Configuration.GetFilters(params.ParentType, params.ParentName, t)
	if err != nil {
		e := misc.HandleError(err)
		return filter.NewGetFiltersDefault(int(*e.Code)).WithPayload(e)
	}
	return filter.NewGetFiltersOK().WithPayload(fs)
}

//Handle executing the request and returning a response
func (h *ReplaceFilterHandlerImpl) Handle(params filter.ReplaceFilterParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	err := h.Client.Configuration.EditFilter(params.ID, params.ParentType, params.ParentName, params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return filter.NewReplaceFilterDefault(int(*e.Code)).WithPayload(e)
	}
	h.ReloadAgent.Reload()
	return filter.NewReplaceFilterOK().WithPayload(params.Data)
}
