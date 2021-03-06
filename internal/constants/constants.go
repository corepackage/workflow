package constants

// Constants to specify runtime
type Runtime string

const (
	PYTHON_RUNTIME Runtime = "python"
	GO_RUNTIME             = "go"
	JAVA_RUNTIME           = "java"
)

// To specify step types
type StepType string

const (
	LOGIC_STEP StepType = "logic"
	API_STEP            = "api"
	GRPC_STEP           = "grpc"
)

// To specify CLI commands type

const (
	RUN_ENGINE       string = "run"
	STOP_ENGINE             = "stop"
	PUSH_CONFIG             = "push"
	LIST_ALL_CONFIGS        = "list"
	REMOVE                  = "rm"
)

// To specify various flags
const (
	HELP         string = "--help"
	ALL                 = "all"
	ALL_SHORT           = "a"
	PORT                = "port"
	PORT_SHORT          = "p"
	PATH                = "path"
	DETACH              = "detach"
	DETACH_SHORT        = "d"
)

// To specify 32 byte encrypt-decrypt key
const (
	ENC_DEC_KEY  = "D*F-JaNdRgUkXp2s5v8y/B?E(H+KbPeS"
	ENC_BASE_DIR = "./configs/workflow-configs"
)

// PID file to store process id for the workflow engine
const PID_FILE = "/tmp/workflow.pid"

// Database file path
const DB_PATH = "./configs/engine-configs/workflow.db"

// Regex patterns

const (
	DATA_REGEX  string = `\$\$[[key]](\.[a-zA-Z0-9]*)*`
	QUERY_REGEX        = `\$\$queryParams(\.[a-zA-Z0-9]*)*`
)
