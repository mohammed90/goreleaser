package dist

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/goreleaser/goreleaser/pkg/config"
	"github.com/goreleaser/goreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

func TestDistDoesNotExist(t *testing.T) {
	var folder = t.TempDir()
	var dist = filepath.Join(folder, "dist")
	require.NoError(
		t,
		Pipe{}.Run(
			&context.Context{
				Config: config.Project{
					Dist: dist,
				},
			},
		),
	)
}

func TestPopulatedDistExists(t *testing.T) {
	var folder = t.TempDir()
	var dist = filepath.Join(folder, "dist")
	require.NoError(t, os.Mkdir(dist, 0755))
	f, err := os.Create(filepath.Join(dist, "mybin"))
	require.NoError(t, err)
	require.NoError(t, f.Close())
	var ctx = &context.Context{
		Config: config.Project{
			Dist: dist,
		},
	}
	require.Error(t, Pipe{}.Run(ctx))
	ctx.RmDist = true
	require.NoError(t, Pipe{}.Run(ctx))
	_, err = os.Stat(dist)
	require.False(t, os.IsExist(err))
}

func TestEmptyDistExists(t *testing.T) {
	var folder = t.TempDir()
	var dist = filepath.Join(folder, "dist")
	require.NoError(t, os.Mkdir(dist, 0755))
	var ctx = &context.Context{
		Config: config.Project{
			Dist: dist,
		},
	}
	require.NoError(t, Pipe{}.Run(ctx))
	_, err := os.Stat(dist)
	require.False(t, os.IsNotExist(err))
}

func TestDescription(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}
