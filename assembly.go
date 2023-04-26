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
	CommitUnix   string            // commit unix, timestamp
	CommitTime   string            // commit time
	BuildTime    string            // build time
	BuildNumber  string            // build num
	BuildJob     string            // build job
	Args         string            // other args
	HideArgs     string            // hide args
)

func Info() map[string]string {
	if len(CommitUnix) > 0 {
		commitTime, _ := strconv.ParseInt(CommitUnix, 10, 64)
		CommitTime = time.Unix(commitTime, 0).Format("2006-01-02 15:04:05")
	}
	buildTime, _ := strconv.ParseInt(BuildTime, 10, 64)
	return map[string]string{
		"name":         Name,
		"version":      Version,
		"comit_id":     CommitID,
		"comit_author": CommitAuthor,
		"comit_time":   CommitTime,
		"build_time":   time.Unix(buildTime, 0).Format("2006-01-02 15:04:05"),
		"build_number": BuildNumber,
		"build_job":    BuildJob,
		"args":         Args,
	}
}
