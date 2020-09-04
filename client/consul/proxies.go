package consul

import "github.com/hashicorp/go-hclog"

// Implementation of SupportedProxiesAPI used to interact with Consul
type proxiesClient struct {
	logger hclog.Logger
}
