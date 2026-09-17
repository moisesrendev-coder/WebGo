package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}

	sqlScript, err := os.ReadFile("script/table.sql")
	if err != nil {
		t.Fatalf("failed to read table.sql: %v", err)
	}

	_, err = db.Exec(string(sqlScript))
	if err != nil {
		t.Fatalf("failed to execute table.sql: %v", err)
	}

	return db
}

func TestUserRepository(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	userRepo := NewSQLUserRepository(db)

	userID, err := userRepo.CreateUser("Moises", "moises@example.com", "password123", "http://avatar.png")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if userID <= 0 {
		t.Errorf("expected valid userID, got %d", userID)
	}

	user, err := userRepo.GetUserByEmail("moises@example.com")
	if err != nil {
		t.Fatalf("failed to get user by email: %v", err)
	}

	if user.ID != userID {
		t.Errorf("expected user ID %d, got %d", userID, user.ID)
	}
	if user.Profile.UserID != userID {
		t.Errorf("expected profile user_id %d, got %d", userID, user.Profile.UserID)
	}

	authID, err := userRepo.Authenticate("moises@example.com", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}
	if authID != userID {
		t.Errorf("expected authID %d, got %d", userID, authID)
	}

	users, err := userRepo.GetUsers()
	if err != nil {
		t.Fatalf("failed to get users: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}
}

func TestPostRepository(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	userRepo := NewSQLUserRepository(db)
	postRepo := NewSQLPostRepository(db)

	userID, err := userRepo.CreateUser("Moises", "moises@example.com", "password123", "avatar.png")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	postID, err := postRepo.CreatePost("First Post", "https://golang.org", userID)
	if err != nil {
		t.Fatalf("failed to create post: %v", err)
	}

	err = postRepo.AddVote(userID, postID)
	if err != nil {
		t.Fatalf("failed to add vote: %v", err)
	}

	commentID, err := postRepo.AddComment(userID, postID, "Great post!")
	if err != nil {
		t.Fatalf("failed to add comment: %v", err)
	}
	if commentID <= 0 {
		t.Errorf("expected valid commentID, got %d", commentID)
	}

	post, err := postRepo.GetByID(postID)
	if err != nil {
		t.Fatalf("failed to get post by id: %v", err)
	}

	if post.ID != postID {
		t.Errorf("expected post ID %d, got %d", postID, post.ID)
	}
	if post.UserID != userID {
		t.Errorf("expected post UserID %d, got %d", userID, post.UserID)
	}
	if post.UserName != "Moises" {
		t.Errorf("expected post UserName Moises, got %s", post.UserName)
	}
	if post.VoteCount != 1 {
		t.Errorf("expected VoteCount 1, got %d", post.VoteCount)
	}
	if post.CommentCount != 1 {
		t.Errorf("expected CommentCount 1, got %d", post.CommentCount)
	}

	comments, err := postRepo.GetComments(postID)
	if err != nil {
		t.Fatalf("failed to get comments: %v", err)
	}
	if len(comments) != 1 {
		t.Fatalf("expected 1 comment, got %d", len(comments))
	}
	if comments[0].UserName != "Moises" {
		t.Errorf("expected comment author Moises, got %s", comments[0].UserName)
	}
	if comments[0].Body != "Great post!" {
		t.Errorf("expected body 'Great post!', got '%s'", comments[0].Body)
	}

	posts, meta, err := postRepo.GetAll(Filter{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatalf("failed to get all posts: %v", err)
	}
	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(posts))
	}
	if meta.TotalRecords != 1 {
		t.Errorf("expected TotalRecords 1, got %d", meta.TotalRecords)
	}

	// Test PageSize 100 and search query
	_, meta100, err := postRepo.GetAll(Filter{
		Page:     1,
		PageSize: 100,
		Query:    "First",
	})
	if err != nil {
		t.Fatalf("failed to get posts with PageSize 100 and query: %v", err)
	}
	if meta100.TotalRecords != 1 {
		t.Errorf("expected TotalRecords 1 for query 'First', got %d", meta100.TotalRecords)
	}
}
