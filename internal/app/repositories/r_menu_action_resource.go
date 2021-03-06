package repositories

import (
	"context"

	"github.com/chunganhbk/gin-go/internal/app/schema"
)

// IMenuActionResource
type IMenuActionResource interface {

	Query(ctx context.Context, params schema.MenuActionResourceQueryParam, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResourceQueryResult, error)

	Get(ctx context.Context, id string, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResource, error)

	Create(ctx context.Context, item schema.MenuActionResource) error

	Update(ctx context.Context, id string, item schema.MenuActionResource) error

	Delete(ctx context.Context, id string) error

	DeleteByActionID(ctx context.Context, actionID string) error

	DeleteByMenuID(ctx context.Context, menuID string) error
}
