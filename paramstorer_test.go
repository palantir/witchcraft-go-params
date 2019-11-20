package wparams_test

import (
	"testing"

	wparams "github.com/palantir/witchcraft-go-params"
	"github.com/stretchr/testify/assert"
)

// Verifies that "NewParamStorer" combines individual storers and adds values from first to last and ensures key
// uniqueness across safe and unsafe values.
func TestNewParamStorer(t *testing.T) {
	safeParamStorer := wparams.NewSafeParamStorer(map[string]interface{}{"key": "safeValue"})
	unsafeParamStorer := wparams.NewUnsafeParamStorer(map[string]interface{}{"key": "unsafeValue"})

	paramStorer := wparams.NewParamStorer(safeParamStorer, unsafeParamStorer)
	assert.Equal(t, map[string]interface{}{}, paramStorer.SafeParams())
	assert.Equal(t, map[string]interface{}{"key": "unsafeValue"}, paramStorer.UnsafeParams())
}

func TestNewSafeParamStorer(t *testing.T) {
	safeVals := map[string]interface{}{"key": "safeValue"}

	paramStorer := wparams.NewSafeParamStorer(safeVals)
	assert.Equal(t, safeVals, paramStorer.SafeParams())
	assert.Equal(t, map[string]interface{}{}, paramStorer.UnsafeParams())
}

func TestNewUnsafeParamStorer(t *testing.T) {
	unsafeVals := map[string]interface{}{"key": "unsafeValue"}

	paramStorer := wparams.NewUnsafeParamStorer(unsafeVals)
	assert.Equal(t, map[string]interface{}{}, paramStorer.SafeParams())
	assert.Equal(t, unsafeVals, paramStorer.UnsafeParams())
}

func TestNewSafeAndUnsafeParamStorer(t *testing.T) {
	safeVals := map[string]interface{}{"safeKey": "safeValue"}
	unsafeVals := map[string]interface{}{"unsafeKey": "unsafeValue"}

	paramStorer := wparams.NewSafeAndUnsafeParamStorer(safeVals, unsafeVals)
	assert.Equal(t, safeVals, paramStorer.SafeParams())
	assert.Equal(t, unsafeVals, paramStorer.UnsafeParams())
}
