# Automation API

Programmatic infrastructure.

## Godocs
See the full godocs for the most extensive and up to date information including full examples coverage: 

https://pkg.go.dev/github.com/khulnasoft/sdk/go/auto?tab=doc

## Examples

Multiple full working examples with detailed walkthroughs can be found in this repo:

https://github.com/khulnasoft/automation-api-examples


## Overview

Package auto contains the Khulnasoft Automation API, the programmatic interface for driving Khulnasoft programs
without the CLI.
Generally this can be thought of as encapsulating the functionality of the CLI (`khulnasoft up`, `khulnasoft preview`,
`khulnasoft destroy`, `khulnasoft stack init`, etc.) but with more flexibility. This still requires a
CLI binary to be installed and available on your $PATH.

In addition to fine-grained building blocks, Automation API provides three out of the box ways to work with Stacks:

1. Programs locally available on-disk and addressed via a filepath (NewStackLocalSource)
```go
    stack, err := NewStackLocalSource(ctx, "myOrg/myProj/myStack", filepath.Join("..", "path", "to", "project"))
```
2. Programs fetched from a Git URL (NewStackRemoteSource)
```go
	stack, err := NewStackRemoteSource(ctx, "myOrg/myProj/myStack", GitRepo{
		URL:         "https:github.com/khulnasoft/test-repo.git",
		ProjectPath: filepath.Join("project", "path", "repo", "root", "relative"),
    })
```
3. Programs defined as a function alongside your Automation API code (NewStackInlineSource)
```go
	 stack, err := NewStackInlineSource(ctx, "myOrg/myProj/myStack", "myProj", func(pCtx *khulnasoft.Context) error {
		bucket, err := s3.NewBucket(pCtx, "bucket", nil)
		if err != nil {
			return err
		}
		pCtx.Export("bucketName", bucket.Bucket)
		return nil
     })
```

Each of these creates a stack with access to the full range of Khulnasoft lifecycle methods
(up/preview/refresh/destroy), as well as methods for managing config, stack, and project settings.

```go
	 err := stack.SetConfig(ctx, "key", ConfigValue{ Value: "value", Secret: true })
	 preRes, err := stack.Preview(ctx)
	 detailed info about results
     fmt.Println(preRes.prev.Steps[0].URN)
```

The Automation API provides a natural way to orchestrate multiple stacks,
feeding the output of one stack as an input to the next as shown in the package-level example below.
The package can be used for a number of use cases:

- Driving khulnasoft deployments within CI/CD workflows
- Integration testing
- Multi-stage deployments such as blue-green deployment patterns
- Deployments involving application code like database migrations
- Building higher level tools, custom CLIs over khulnasoft, etc
- Using khulnasoft behind a REST or GRPC API
- Debugging Khulnasoft programs (by using a single main entrypoint with "inline" programs)

To enable a broad range of runtime customization the API defines a `Workspace` interface.
A Workspace is the execution context containing a single Khulnasoft project, a program, and multiple stacks.
Workspaces are used to manage the execution environment, providing various utilities such as plugin
installation, environment configuration ($KHULNASOFT_HOME), and creation, deletion, and listing of Stacks.
Every Stack including those in the above examples are backed by a Workspace which can be accessed via:
```go
	 w = stack.Workspace()
     err := w.InstallPlugin("aws", "v3.2.0")
```
Workspaces can be explicitly created and customized beyond the three Stack creation helpers noted above:
```go
	 w, err := NewLocalWorkspace(ctx, WorkDir(filepath.Join(".", "project", "path"), KhulnasoftHome("~/.khulnasoft"))
     s := NewStack(ctx, "org/proj/stack", w)
```
A default implementation of workspace is provided as `LocalWorkspace`. This implementation relies on Khulnasoft.yaml
and Khulnasoft.<stack>.yaml as the intermediate format for Project and Stack settings. Modifying ProjectSettings will
alter the Workspace Khulnasoft.yaml file, and setting config on a Stack will modify the Khulnasoft.<stack>.yaml file.
This is identical to the behavior of Khulnasoft CLI driven workspaces. Custom Workspace
implementations can be used to store Project and Stack settings as well as Config in a different format,
such as an in-memory data structure, a shared persistent SQL database, or cloud object storage. Regardless of
the backing Workspace implementation, the Khulnasoft SaaS Console will still be able to display configuration
applied to updates as it does with the local version of the Workspace today.

The Automation API also provides error handling utilities to detect common cases such as concurrent update
conflicts:

```go
	uRes, err :=stack.Up(ctx)
	if err != nil && IsConcurrentUpdateError(err) { /* retry logic here */ }
```

## Developing the Godocs
This repo has extensive examples and godoc content. To test out your changes locally you can do the following:

1. enlist in the appropriate khulnasoft branch:
2. cd $GOPATH/src/github.com/khulnasoft/sdk/go/auto
3. godoc -http=:6060
4. Navigate to http://localhost:6060/pkg/github.com/khulnasoft/sdk/go/auto/

## Known Issues

Please upvote issues, add comments, and open new ones to help prioritize our efforts:
https://github.com/khulnasoft/khulnasoft/issues?q=is%3Aissue+is%3Aopen+label%3Aarea%2Fautomation-api
