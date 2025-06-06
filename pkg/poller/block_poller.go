package poller

import "context"

type BlockPoller interface {
	Poll(context.Context, chan string, chan error)
}
