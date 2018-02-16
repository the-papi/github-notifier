package settings

var HomeDir = ""

const (
	// Relative paths to home directory
	DataPath   = "/.local/share/github-notifier" // Data path
	ConfigPath = "/.config/github-notifier"      // Config directory path

	SrcRoot        = "/src/github.com/PapiCZ/github-notifier/" // Relative to GOPATH
	PidFileName    = ".github-notifier.pid"
	ConfigFileName = "config.json"
	IconFileName   = "octocat.png"
)
