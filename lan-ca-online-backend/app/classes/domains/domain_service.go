package domains

import (
	"context"

	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
)

type Service interface {
	Find(ctx context.Context, id dxo.DomainID) (*dto.Domain, error)

	List(ctx context.Context, q *Query) ([]*dto.Domain, error)

	Insert(ctx context.Context, item *dto.Domain) (*dto.Domain, error)
}
