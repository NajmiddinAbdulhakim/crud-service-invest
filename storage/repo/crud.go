package repo

import (
	pb "github.com/NajmiddinAbdulhakim/iman/crud-service/genproto"

)

type CRUDStorage interface {
	Create(p *pb.Post) (*pb.Post, error)
	GetById(id int64) (*pb.Post, error)
	Update(p *pb.Post) (*pb.Post, error)
	Delete(id int64) (bool, error)
	ListPosts(page, limit int64) ([]*pb.Post, error)
}