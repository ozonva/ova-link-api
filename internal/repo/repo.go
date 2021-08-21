package repo

import "github.com/ozonva/ova-link-api/internal/link"

type Repo interface {
	AddEntities(entities []link.Link) error
	ListEntities(limit uint64, offset uint64) ([]link.Link, error)
	DescribeEntity(entityId uint64) (*link.Link, error)
}
