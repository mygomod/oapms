package main

import (
	"github.com/mygomod/muses"
	"github.com/mygomod/muses/pkg/cache/redis"
	"github.com/mygomod/muses/pkg/cmd"
	"github.com/mygomod/muses/pkg/database/mysql"
	musgin "github.com/mygomod/muses/pkg/server/gin"
	"github.com/mygomod/muses/pkg/server/stat"
	"github.com/mygomod/muses/pkg/session/ginsession"
	"github.com/spf13/cobra"
	"oapms/pkg/command"
	"oapms/pkg/mus"
	"oapms/pkg/router"
	"oapms/pkg/service"
)

func main() {
	app := muses.Container(
		cmd.Register,
		stat.Register,
		redis.Register,
		mysql.Register,
		musgin.Register,
		ginsession.Register,
	)
	app.SetRootCommand(func(cobraCommand *cobra.Command) {
		cobraCommand.AddCommand(command.InstallCmd)
	})
	app.SetGinRouter(router.InitRouter)
	app.SetPostRun(mus.Init, service.Init)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
