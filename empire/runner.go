package empire

import "golang.org/x/net/context"

// RunOpts are options that can be provided when running a container.
type RunOpts struct {
	// If provided, overrides the default command.
	Command *string

	// If true, attaches stdin and stdout to the container.
	Attach bool

	// Any extra environment variables to run.
	Env map[string]string
}

// runner is something that's capable of running an App.
type runner interface {
	Run(context.Context, *App, RunOpts) error
}

// nullRunner is a runner implementation that does nothing.
type nullRunner struct{}

func (r *nullRunner) Run(ctx context.Context, app *App, opts RunOpts) error {
	return nil
}
