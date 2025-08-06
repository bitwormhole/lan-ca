package scenes

import (
	"context"

	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
)

type Service interface {
	Find(ctx context.Context, id dxo.SceneID) (*dto.Scene, error)

	List(ctx context.Context, q *Query) ([]*dto.Scene, error)

	Insert(ctx context.Context, item *dto.Scene) (*dto.Scene, error)
}
