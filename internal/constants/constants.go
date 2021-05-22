package constants

type Runtime string

const (
	PYTHON_RUNTIME Runtime = "python"
	GO_RUNTIME             = "go"
	JAVA_RUNTIME           = "java"
)

type StepType string

const (
	LOGIC_STEP StepType = "logic"
	API_STEP            = "api"
	GRPC_STEP           = "grpc"
)
