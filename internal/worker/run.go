package worker

import "context"

// Run keeps the node link alive and folds every decoded event into the registry
// until ctx is cancelled.
func (w *Worker) Run(ctx context.Context) {
	w.mesh.Connect(ctx)

	for {
		ev, err := w.mesh.Read(ctx)
		if err != nil {
			return
		}

		w.apply(ev)
	}
}
