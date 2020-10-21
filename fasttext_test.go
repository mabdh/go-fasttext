package fasttext

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFoo(t *testing.T) {
	m, err := New("dbpedia.ftz")
	require.Nil(t, err)
	result := m.Predict("this is testing", 5, float32(0.01))

	expectedResult := []Result{
		{
			Label: "__label__12",
			Prob:  0.63721514,
		},
		{

			Label: "__label__10",
			Prob:  0.3319999,
		},
		{

			Label: "__label__9",
			Prob:  0.027013373,
		},
	}
	m.Free()
	assert.Equal(t, result, expectedResult)
}
