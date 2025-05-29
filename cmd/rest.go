package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arvinpaundra/go-boilerplate/api/route"
	"github.com/arvinpaundra/go-boilerplate/config"
	"github.com/arvinpaundra/go-boilerplate/core"
	"github.com/arvinpaundra/go-boilerplate/core/validator"
	"github.com/arvinpaundra/go-boilerplate/database/nosqlpkg"
	"github.com/arvinpaundra/go-boilerplate/database/sqlpkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var port string

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Start rest server",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".", ".env", "env")

		pgsql := sqlpkg.NewPostgres()

		sqlpkg.NewConnection(pgsql)

		rdb := nosqlpkg.NewRedisDB()

		nosqlpkg.NewInMemoryConection(rdb)

		g := gin.New()

		route.NewRoutes(
			g,
			sqlpkg.GetConnection(),
			nosqlpkg.GetInMemoryConnection(),
			validator.NewValidator(),
		).GatherRoutes()

		srv := http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: g,
		}

		go func() {
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("failed to start server: %s", err.Error())
			}
		}()

		wait := core.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"rest-server": func(_ context.Context) error {
				return srv.Close()
			},
			"postgres": func(_ context.Context) error {
				return pgsql.Close()
			},
			"redis": func(_ context.Context) error {
				return rdb.Close()
			},
		})

		_ = <-wait
	},
}

func init() {
	restCmd.Flags().StringVarP(&port, "port", "p", "8080", "bind server to port. default: 8080")
	rootCmd.AddCommand(restCmd)
}
