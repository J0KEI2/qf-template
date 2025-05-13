package repository

import (
	"log"
	"os"
	"os/exec"

	"github.com/AlekSi/pointer"
	"github.com/spf13/viper"
	helpers "github.com/zercle/gofiber-helpers"
)

func (u reportRepository) ExportReport(QFMainID, qfSubplanId *string) (string, error) {
	err := os.MkdirAll("/app/public/documents/" + pointer.GetString(QFMainID) + "/report/report/", 0755)
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return "", err
	}
	host := viper.GetString("db.postgres.username") + "@" + viper.GetString("db.postgres.host") + ":" + viper.GetString("db.postgres.port")
	password := viper.GetString("db.postgres.password")
	database := viper.GetString("db.postgres.db_name")
	command := exec.Command("python3", "/app/lib/main.py",
		"--host", host,
		"--password", password,
		"--database", database,
		"-qf", pointer.GetString(QFMainID),
		"-sp", pointer.GetString(qfSubplanId))
	out, err := command.CombinedOutput()
	if err != nil {
		log.Printf("out = %v", string(out))
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return "", err
	}
	return string(out), nil
}
