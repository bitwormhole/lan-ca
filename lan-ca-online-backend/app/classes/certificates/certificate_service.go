package certificates

import (
	"context"

	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
)

type Service interface {
	Find(ctx context.Context, id dxo.CertificateID) (*dto.Certificate, error)

	List(ctx context.Context, q *Query) ([]*dto.Certificate, error)
}
