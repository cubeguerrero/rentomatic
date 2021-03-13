package use_cases

import "github.com/cubeguerrero/rentomatic/domain"

type roomRepoInterface interface {
	List() []*domain.Room
}
