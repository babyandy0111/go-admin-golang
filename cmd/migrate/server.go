package migrate

import (
	"bytes"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"strconv"
	"text/template"
	"time"

	"github.com/go-admin-team/go-admin-core/config/source/file"
	"github.com/spf13/cobra"

	"github.com/go-admin-team/go-admin-core/sdk/config"
	"go-admin/cmd/migrate/migration"
	_ "go-admin/cmd/migrate/migration/version"
	_ "go-admin/cmd/migrate/migration/version-local"
	"go-admin/common/database"
	"go-admin/common/models"
)

var (
	configYml string
	generate  bool
	goAdmin   bool
	host      string
	StartCmd  = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "go-admin migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
	StartCmd.PersistentFlags().BoolVarP(&goAdmin, "goAdmin", "a", false, "generate go-admin migration file")
	StartCmd.PersistentFlags().StringVarP(&host, "domain", "d", "*", "select tenant host")
}

func run() {

	if !generate {
		fmt.Println(`start init`)
		//1. 读取配置
		config.Setup(
			file.NewSource(file.WithPath(configYml)),
			initDB,
		)
	} else {
		fmt.Println(`generate migration file`)
		_ = genFile()
	}
}

func migrateModel() error {
	if host == "" {
		host = "*"
	}
	db := sdk.Runtime.GetDbByKey(host)
	if config.DatabasesConfig[host].Driver == "mysql" {
		//初始化資料库時候用
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	err := db.Debug().AutoMigrate(&models.Migration{})
	if err != nil {
		return err
	}
	migration.Migrate.SetDb(db.Debug())
	migration.Migrate.Migrate()
	return err
}
func initDB() {
	//3. 初始化資料库链接
	database.Setup()
	//4. 資料库迁移
	fmt.Println("migrate start")
	_ = migrateModel()
	fmt.Println(`migrate init ok`)
}

func genFile() error {
	t1, err := template.ParseFiles("template/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version_local"
	if goAdmin {
		m["Package"] = "version"
	}
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if goAdmin {
		pkg.FileCreate(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
	} else {
		pkg.FileCreate(b1, "./cmd/migrate/migration/version-local/"+m["GenerateTime"]+"_migrate.go")
	}
	return nil
}
