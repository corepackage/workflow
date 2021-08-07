# Workflow

An automation engine to run different workflows. The workflow engine is a stand alone CLI engine which takes workflow configurations and expose respective API to run workflows.

Please note, the project is currently in development.

## Installing dependencies

To install the necessary dependencies, run the following command on root of the folder

```
go get ./...
```

## Installing the engine

To install the engine for use case, run the following command

```
go install ./...
```

## Creating a workflow config

To create a workflow config refer the [sample.yml](example/config/sample.yml) for proper format.

To use dynamic data for the configurations use `$$` prefix before accessing any data from the preceeding steps or request data. This dynamic data can be used in payload, params or api endpoint.

Current version supports the following configurations :

1. Async steps
2. Delay in executing
3. Execution timeout
4. Breakpoint

**NOTE** : Currently, data can be set from query params, inpuy body json or preceeding steps (using step id).

## Types of steps

Current version supports two types of steps :

1. API step : This type of step can be configured to access API endpoints. Workflow automatically calls the API using the configured set in yaml file.

2. LOGIC step : To execute a binary via workflow, setup the path for the exe in workflow and setup the input parameters for the exe. To integrate binaries with workflow use the [workflow-sdk](https://github.com/corepackage/workflow-sdk).

**NOTE:** Currently the workflow supports only go binaries.

## Pushing the config to workflow

To push a new created config, use the following command :

```
workflow push <file-path>:<version>
```

If version not specified, default version taken is `latest`.
You can push multiple files at a time by space separating each path.

**NOTE:** Currently we support only yaml files

## Starting the engine

To start the engine, run the following command:

```
workflow run [-options]
```

This command will run all the active configs present within the workflow. The default port at which the workflow engine runs is `:7200`.

All the workflow configs will expose the respective API endpoint to execute.

For e.g: a workflow with id `sampleid` can be accessed on `http://localhost:7200/sampleid`

Following are the options which are supported with the run command:

`--port` or `-p` : To specify a port to run the workflow

`--path` : To specify a prefix path for the APIs

`--detach` or `-d` : To run the workflow in background mode

## Stopping the workflow

To stop the engine running in background, run the following command:

```
workflow stop
```

To stop a specific config, run the following command:

```
workflow stop <workflow-id/workflow-name>:<version>
```

You can specify multiple workflows by space separating the names.

**NOTE:** If version is not specified, default `latest` version will taken

## To get list of the configs present in workflow

To get the list of the active configs, run the following command:

```
workflow list [-options]
```

Following are the options which are supported with the run command:

`--all` or `-a` : To list all the configs

## Remove a config

To remove the config from the workflow, run the following command:

```
workflow rm <workflow-id/workflow-name>:<version>
```

## To get help

```
workflow --help
```
