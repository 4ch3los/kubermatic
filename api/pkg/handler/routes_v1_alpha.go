package handler

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

// RegisterV1Alpha declares all HTTP paths that are experimental
// and may change in the future in non compatible way
func (r Routing) RegisterV1Alpha(mux *mux.Router) {
	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/metrics").
		Handler(r.getClusterMetrics())
}

// swagger:route GET /api/v1alpha/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/metrics project getClusterMetrics
//
//    Gets cluster metrics
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: []ClusterMetric
//       401: empty
//       403: empty
func (r Routing) getClusterMetrics() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			r.authenticator.Verifier(),
			r.userSaverMiddleware(),
			r.datacenterMiddleware(),
			r.userInfoMiddleware(),
		)(getClusterMetrics(r.projectProvider, r.prometheusClient)),
		decodeGetClusterReq,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}
