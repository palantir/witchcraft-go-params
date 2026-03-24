package wparams_test

import (
	"testing"

	wparams "github.com/palantir/witchcraft-go-params"
	"github.com/stretchr/testify/assert"
)

// Verifies that "NewParamStorer" combines individual storers and adds values from first to last and ensures key
// uniqueness across safe and unsafe values.
func TestNewParamStorer(t *testing.T) {
	safeParamStorer := wparams.NewSafeParamStorer(map[string]any{"key": "safeValue"})
	unsafeParamStorer := wparams.NewUnsafeParamStorer(map[string]any{"key": "unsafeValue"})

	paramStorer := wparams.NewParamStorer(safeParamStorer, unsafeParamStorer)
	assert.Equal(t, map[string]any{}, paramStorer.SafeParams())
	assert.Equal(t, map[string]any{"key": "unsafeValue"}, paramStorer.UnsafeParams())
}

func TestNewSafeParamStorer(t *testing.T) {
	safeVals := map[string]any{"key": "safeValue"}

	paramStorer := wparams.NewSafeParamStorer(safeVals)
	assert.Equal(t, safeVals, paramStorer.SafeParams())
	assert.Equal(t, map[string]any{}, paramStorer.UnsafeParams())
}

func TestNewUnsafeParamStorer(t *testing.T) {
	unsafeVals := map[string]any{"key": "unsafeValue"}

	paramStorer := wparams.NewUnsafeParamStorer(unsafeVals)
	assert.Equal(t, map[string]any{}, paramStorer.SafeParams())
	assert.Equal(t, unsafeVals, paramStorer.UnsafeParams())
}

func TestNewSafeAndUnsafeParamStorer(t *testing.T) {
	safeVals := map[string]any{"safeKey": "safeValue"}
	unsafeVals := map[string]any{"unsafeKey": "unsafeValue"}

	paramStorer := wparams.NewSafeAndUnsafeParamStorer(safeVals, unsafeVals)
	assert.Equal(t, safeVals, paramStorer.SafeParams())
	assert.Equal(t, unsafeVals, paramStorer.UnsafeParams())
}

func TestParamStorerSideEffects(t *testing.T) {
	safeVals := map[string]any{"safeKey": "safeValue"}
	unsafeVals := map[string]any{"unsafeKey": "unsafeValue"}
	paramStorer := wparams.NewSafeAndUnsafeParamStorer(safeVals, unsafeVals)

	// edit maps and expect no change to params
	safeVals["safeKey1"] = "safeValue1"
	delete(unsafeVals, "unsafeKey")
	assert.NotContains(t, paramStorer.SafeParams(), "safeKey1")
	assert.Contains(t, paramStorer.UnsafeParams(), "unsafeKey")

	// inherit storer and expect no change to maps
	newStorer := wparams.NewParamStorer(paramStorer, wparams.NewSafeParam("newParam", "newValue"))
	assert.NotContains(t, paramStorer.SafeParams(), "newParam")
	assert.Contains(t, newStorer.SafeParams(), "newParam")

	// overwrite key and ensure it does not affect parents
	newStorer = wparams.NewParamStorer(newStorer, wparams.NewUnsafeParam("newParam", "unsafeValue"))
	assert.NotContains(t, newStorer.SafeParams(), "newParam")
	assert.Contains(t, newStorer.UnsafeParams(), "newParam")

	// overwrite returned map, verify it does overwrite underlying value
	// this is not desirable, but we do it so we don't allocate a new map on lookups
	sp := newStorer.SafeParams()
	sp["foo"] = "bar"
	assert.Contains(t, newStorer.SafeParams(), "foo")
}
