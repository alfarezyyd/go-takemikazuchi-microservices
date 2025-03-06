package discovery

import (
	"context"
	"errors"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"log"
	"strconv"
	"strings"
)

type ConsulServiceRegistry struct {
	consulClient *consul.Client
}

func NewRegistry(serviceAddress, serviceName string) (*ConsulServiceRegistry, error) {
	config := consul.DefaultConfig()
	config.Address = serviceAddress

	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConsulServiceRegistry{client}, nil
}

func (serviceRegistry *ConsulServiceRegistry) Register(ctx context.Context, instanceID, serviceName, hostPort string) error {
	host, portStr, found := strings.Cut(hostPort, ":")
	if !found {
		return errors.New("invalid host:port format. Eg: localhost:8081")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	return serviceRegistry.consulClient.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		ID:      instanceID,
		Address: host,
		Port:    port,
		Name:    serviceName,
		Check: &consul.AgentServiceCheck{
			CheckID:                        instanceID,
			TLSSkipVerify:                  true,
			TTL:                            "5s",
			Timeout:                        "1s",
			DeregisterCriticalServiceAfter: "10s",
		},
	})
}

func (serviceRegistry *ConsulServiceRegistry) Deregister(ctx context.Context, instanceID string, serviceName string) error {
	log.Printf("Deregistering service %s", instanceID)
	return serviceRegistry.consulClient.Agent().CheckDeregister(instanceID)
}

func (serviceRegistry *ConsulServiceRegistry) HealthCheck(instanceID string, serviceName string) error {
	return serviceRegistry.consulClient.Agent().UpdateTTL(instanceID, "online", consul.HealthPassing)
}

func (serviceRegistry *ConsulServiceRegistry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := serviceRegistry.consulClient.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	var instances []string
	for _, entry := range entries {
		instances = append(instances, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
	}

	return instances, nil
}
