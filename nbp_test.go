package nbp_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tzdybal/go-nbp"
)

func TestGetCurrentRates(t *testing.T) {
	c := nbp.New()
	require.NotNil(t, c)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	results, err := c.GetCurrentRates(ctx, nbp.TableA)
	require.NoError(t, err)
	require.NotNil(t, results)
}

func Test_GetNLastRates(t *testing.T) {
	c := nbp.New()
	require.NotNil(t, c)

	cases := []struct {
		name        string
		n           uint64
		expectError bool
	}{
		{
			"valid n value", 14, false,
		},
		{
			"invalid n value", 15, true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()

			results, err := c.GetNLastRates(ctx, nbp.TableB, tc.n)
			if tc.expectError {
				require.Error(t, err)
				require.Nil(t, results)
			} else {
				require.NoError(t, err)
				require.NotNil(t, results)
				require.Len(t, results, int(tc.n))
			}
		})
	}
}
