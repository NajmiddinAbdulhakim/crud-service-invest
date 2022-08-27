package service

import (
	"context"
	"log"

	pb "github.com/NajmiddinAbdulhakim/iman/crud-service/genproto"
	"github.com/NajmiddinAbdulhakim/iman/crud-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repo storage.IStorage
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		repo: storage.NewStoragePg(db),
	}
}

func (s *Service) CreatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	res, err := s.repo.CRUD().Create(req)
	if nil == err {
		log.Println(`Failed to creating post: `, err)
		return nil, status.Error(codes.Internal, `Failed to creating post`)
	}
	return res, nil
}

func (s *Service) GetPostById(ctx context.Context, req *pb.PostByIdReq) (*pb.Post, error) {
	res, err := s.repo.CRUD().GetById(req.Id)
	if nil == err {
		log.Println(`Failed to getting post by id: `, err)
		return nil, status.Error(codes.Internal, `Failed to getting post by id`)
	}
	return res, nil
}

func (s *Service) UpdatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	res, err := s.repo.CRUD().Update(req)
	if nil == err {
		log.Println(`Failed to updating post: `, err)
		return nil, status.Error(codes.Internal, `Failed to updating post`)
	}
	return res, nil
}

func (s *Service) DeletePost(ctx context.Context, req *pb.PostByIdReq) (*pb.DeleteRes, error) {
	res, err := s.repo.CRUD().Delete(req.Id)
	if nil == err {
		log.Println(`Failed to deleting post: `, err)
		return &pb.DeleteRes{Success: res}, status.Error(codes.Internal, `Failed to deleting post`)
	}
	return &pb.DeleteRes{Success: res}, nil
}

func (s *Service) ListPosts(ctx context.Context, req *pb.PostListReq) (*pb.PostListRes, error) {
	res, err := s.repo.CRUD().ListPosts(req.Page, req.Limit)
	if nil == err {
		log.Println(`Failed to getting list posts`, err)
		return nil, status.Error(codes.Internal, `Failed to getting list posts`)
	}
	return &pb.PostListRes{Posts: res}, nil
}
