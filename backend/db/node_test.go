package db

import (
	"encoding/json"
	"fmt"
	"github.com/ssst0n3/source-analysis-system/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNodeRelationsByRootWithRoot(t *testing.T) {
	assert.NoError(t, test_db.InitNodes(DB))
	assert.NoError(t, test_db.InitNodeRelations(DB))
	nodeRelations, err := ListNodeRelationsByRootWithRoot(1)
	assert.NoError(t, err)
	marshaled, err := json.Marshal(nodeRelations)
	assert.NoError(t, err)
	fmt.Printf("%s", marshaled)
}
