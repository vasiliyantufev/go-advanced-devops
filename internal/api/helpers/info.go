package helpers

import log "github.com/sirupsen/logrus"

func PrintInfo(buildVersion string, buildDate string, buildCommit string) {
	log.Infof("Build version: %s", buildVersion)
	log.Infof("Build date: %s", buildDate)
	log.Infof("Build commit: %s", buildCommit)
}
