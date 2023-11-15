package ds

import (
	"time"
)

type AppInfo struct {
	AppName     string `json:"app_name"`
	Version     string `json:"version"`
	BuildTime   string `json:"build_time"`
	BuildOS     string `json:"build_os"`
	BuildCommit string `json:"build_commit"`
	StartupTime string `json:"startup_time"`
}

func NewAppInfo() *AppInfo {
	return &AppInfo{
		AppName:     "crude-app",
		StartupTime: time.Now().Format(time.RFC3339),
	}
}

func (i *AppInfo) WithVersion(version string) *AppInfo {
	i.Version = version
	return i
}

func (i *AppInfo) WithBuildTime(buildTime string) *AppInfo {
	i.BuildTime = buildTime
	return i
}

func (i *AppInfo) WithBuildOS(buildOS string) *AppInfo {
	i.BuildOS = buildOS
	return i
}

func (i *AppInfo) WithBuildCommit(commit string) *AppInfo {
	i.BuildCommit = commit
	return i
}
