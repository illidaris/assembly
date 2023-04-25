package assembly

import (
	"strconv"
	"time"
)

var (
	Name         string = "demo"   // name
	Version      string = "v1.0.0" // verion
	CommitID     string            // code commit
	CommitAuthor string            // author
	CommitTime   string            // commit time
	BuildTime    string            // build time
	BuildNumber  string            // build num
	BuildJob     string            // build job
	Args         string            // other args
	HideArgs     string            // hide args
)

func Info() map[string]string {
	commitTime, _ := strconv.ParseInt(CommitTime, 10, 64)
	buildTime, _ := strconv.ParseInt(BuildTime, 10, 64)
	return map[string]string{
		"name":         Name,
		"version":      Version,
		"comit_id":     CommitID,
		"comit_author": CommitAuthor,
		"comit_time":   time.Unix(commitTime, 0).Format("2006-01-02 15:04:05"),
		"build_time":   time.Unix(buildTime, 0).Format("2006-01-02 15:04:05"),
		"build_number": BuildNumber,
		"build_job":    BuildJob,
		"args":         Args,
	}
}
