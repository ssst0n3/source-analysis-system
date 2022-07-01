package db

import (
	"encoding/json"
	"fmt"
	"github.com/ssst0n3/source-analysis-system/model"
	"github.com/ssst0n3/source-analysis-system/test/test_data"
	"github.com/ssst0n3/source-analysis-system/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_NextRecursive(t *testing.T) {
	m := model.NewMatrix(1, test_data.MapNodeRelations)
	m.ChildRecursive(0)
	m.Print()
}

func TestNodeMatrix(t *testing.T) {
	assert.NoError(t, test_db.InitNodes(DB))
	assert.NoError(t, test_db.InitNodeRelations(DB))
	matrix, err := NodeMatrix(1)
	assert.NoError(t, err)
	marshaled, err := json.Marshal(matrix)
	assert.NoError(t, err)
	fmt.Printf("%s\n", marshaled)
}
