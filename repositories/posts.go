package repositories

import (
	"database/sql"
	"devbook-api/models"
)

type posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (postsRepository posts) Create(post models.Post) (uint64, error) {
	statements, err := postsRepository.db.Prepare(
		"insert into posts (title, content, author_id) values(?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statements.Close()

	result, err := statements.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

func (postsRepository posts) GetPostByID(postID uint64) (models.Post, error) {
	row, err := postsRepository.db.Query(`
		select p.*, u.nick from
		posts p inner join users u
		on u.id = p.author_id where p.id = ?
	`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post
	if row.Next() {
		if err := row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (postsRepository posts) GetPosts(userID uint64) ([]models.Post, error) {
	rows, err := postsRepository.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		inner join followers s on p.author_id = s.userID
		where u.id = ? or s.followerID = ?
		order by 1 desc
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (postsRepository posts) Update(postID uint64, post models.Post) error {
	statement, err := postsRepository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (postsRepository posts) Delete(postID uint64) error {
	statement, err := postsRepository.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (postsRepository posts) FindByUserID(userID uint64) ([]models.Post, error) {
	rows, err := postsRepository.db.Query(`
		select p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		where p.author_id = ?
		order by 1 desc
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (postsRepository posts) Like(postID uint64) error {
	statement, err := postsRepository.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (postsRepository posts) Unlike(postID uint64) error {
	statement, err := postsRepository.db.Prepare("update posts set likes = likes - 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
