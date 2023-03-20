// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	ai "portal/internal/handler/ai"
	"portal/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: healthHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: ai.ChatHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/portal/ai"),
	)
}
