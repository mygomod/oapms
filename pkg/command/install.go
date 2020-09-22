package command

import (
	"github.com/mygomod/muses"
	"github.com/mygomod/muses/pkg/cache/redis"
	mmysql "github.com/mygomod/muses/pkg/database/mysql"
	"github.com/mygomod/muses/pkg/session/ginsession"
	"github.com/spf13/cobra"
	"oapms/pkg/command/install"
	"oapms/pkg/mus"
	"oapms/pkg/service"
)

var InstallCmd = &cobra.Command{
	Use:  "install",
	Long: `Show install information`,
	Run:  installCmd,
}

var ConfigPath string
var InstallMode string

func init() {
	InstallCmd.PersistentFlags().StringVarP(&ConfigPath, "conf", "c", "conf/conf.toml", "conf path")
	InstallCmd.PersistentFlags().StringVarP(&InstallMode, "mode", "m", "all", "mode type")
}

func installCmd(cmd *cobra.Command, args []string) {
	app := muses.Container(
		mmysql.Register,
		ginsession.Register,
		redis.Register,
	)
	app.SetCfg(ConfigPath)

	app.SetPostRun(mus.Init, service.Init, func() error {
		var err error
		if InstallMode == "all" || InstallMode == "install" {
			err = install.Create()
			if err != nil {
				return err
			}
		}
		if InstallMode == "all" || InstallMode == "mock" {
			install.MockData()
			return nil
		}
		return nil
	})
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
