package webhooksecret

import (
	"context"
	"net/url"

	"k8s.io/apimachinery/pkg/types"
)

// HookClient implementations provide functionality for creating hooks in a Git
// Hosting Service.
type HookClient interface {
	CreateHook(context.Context, *url.URL, string) error
}

// RouteGetter implementations get the URL for OpenShift Routes.
type RouteGetter interface {
	RouteURL(types.NamespacedName) (*url.URL, error)
}