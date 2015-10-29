package db

import "github.com/pdxjohnny/sysnet/db/dbtypes"

// Open returns a collection which alows the user to access documents
func Open(dbType, collectionName, addr string) (dbtypes.Collection, error) {
	return nil, nil
}
