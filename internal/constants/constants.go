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
	RUN_ENGINE  string = "run"
	STOP_ENGINE        = "stop"
	PUSH_CONFIG        = "push"
	HELP               = "--help"
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
