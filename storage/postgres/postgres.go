package postgres

import (
	pb "github.com/NajmiddinAbdulhakim/iman/crud-service/genproto"
	"github.com/jmoiron/sqlx"
)

type crudRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *crudRepo {
	return &crudRepo{db: db}
}

func (r *crudRepo) Create(p *pb.Post) (*pb.Post, error) {
	query := `INSERT INTO posts VALUES($1, $2, $3, $4) RETURNING *`
	row := r.db.QueryRow(query, p.Id, p.UserId, p.Title, p.Body)

	var post *pb.Post
	err := row.Scan(
		post.Id,
		post.UserId,
		post.Title,
		post.Body,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *crudRepo) GetById(id int64) (*pb.Post, error) {
	query := `SELECT * from posts WHERE id = $1`
	row := r.db.QueryRow(query, id)
	var post *pb.Post
	err := row.Scan(
		post.Id,
		post.UserId,
		post.Title,
		post.Body,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *crudRepo) Update(p *pb.Post) (*pb.Post, error) {
	query := `UPDATE posts SET user_id = $1, title = $2, body = $3 WHERE id = $4`
	_, err := r.db.Exec(query, p.UserId, p.Title, p.Body, p.Id)
	if err != nil {
		return nil, err
	}

	getQuery := `SELECT * from posts WHERE id = $1`
	row := r.db.QueryRow(getQuery, p.Id)

	var post *pb.Post
	err = row.Scan(
		post.Id,
		post.UserId,
		post.Title,
		post.Body,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *crudRepo) Delete(id int64) (bool, error) {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return false, err
	}
	return true, err
}

func (r *crudRepo) ListPosts(page, limit int64) ([]*pb.Post, error) {
	query := `SELECT * FROM posts OFFSET $1 LIMIT $2`
	rows, err := r.db.Query(query, page, limit)
	if err != nil {
		return nil, err
	}
	
	var posts []*pb.Post
	for rows.Next() {
		var p pb.Post
		err := rows.Scan(
			&p.Id,
			&p.UserId,
			&p.Title,
			&p.Body,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &p)
	}
	return posts, nil
}
