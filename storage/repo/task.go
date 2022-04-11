package repo

import (
	pb "github.com/Shahboz4131/7-project/genproto"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id string) (pb.Task, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id string) error
}
