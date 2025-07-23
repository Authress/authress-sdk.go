package authress

type BuildInfo struct {
	Version string
}

func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version: "v0.9.23",
	}
}
