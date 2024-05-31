package migrations

import (
	"context"

	"github.com/steevedevops/simplecruda/src/app/post"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*post.Post)(nil)).
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
			Model((*post.Post)(nil)).
			IfExists().
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
