package post

import(
	"github.com/youkyll/goshepherd/app/model"
	"time"
	"database/sql"
	"log"
)

var db = model.InitDb()

func init() {
	db.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
}

type Post struct {
	Id int `db:"post_id"`
	Title string `db:"title"`
	Content string `db:"content"`
	Created time.Time `db:"created_at"`
	Updated time.Time `db:"updated_at"`
}

type DetailPost struct {
	Post
	FolderId sql.NullInt64 `db:"folder_id"`
	FolderName sql.NullString `db:"folder_name" `
}

func TakeByFolder(tag_id int, page int, per_page int) []Post {
	var posts []Post

	db.Select(&posts ,`SELECT post_id FROM posts
	JOIN post_folders AS pt USING(post_id)
	WHERE pt.folder_id = ?
	OFFSET (page - 1) * per_page LIMIT per_page
	ORDER BY updated_at DESC, post_id DESC`)

	return posts
}

func Pick(post_ids []int) []Post {
	var posts []Post

	db.Select(&posts, `SELECT post_id, title, content, updated_at FROM posts
	WHERE post_id IN ?
	ORDER BY updated_at DESC, post_id DESC`)

	return posts
}

func PickOne(post_id int) Post {
	var post Post

	db.SelectOne(&post, `SELECT post_id, title, content, created_at, updated_at FROM posts
	WHERE post_id = ?
	ORDER BY updated_at DESC, post_id DESC`, post_id)

	return post
}

func DetailOne(post_id int) DetailPost {
	var post DetailPost

	db.SelectOne(&post, `SELECT p.post_id, title, content, p.created_at, p.updated_at, folder_id, folder_name FROM posts AS p
	LEFT JOIN post_folders AS pf USING(post_id) LEFT JOIN folders AS f USING(folder_id) WHERE p.post_id = ?
	ORDER BY p.updated_at DESC, post_id DESC`, post_id)

	return post
}


func All() []Post {
	var posts []Post

	db.Select(&posts, `SELECT * FROM posts ORDER BY updated_at DESC, post_id DESC`)

	return posts
}

func SearchTitle(title string) []Post {
	var posts []Post

	db.Select(&posts, `SELECT * FROM posts WHERE title like ? ORDER BY updated_at DESC, post_id DESC`, title + `%`)

	return posts
}

func Add(title string, content string, folder_id int) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	post := Post{
		Title: title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	if err := tx.Insert(&post); err != nil {
		tx.Rollback()
		return 0, err
	}

	if folder_id != 0 {
		if _, err = tx.Exec(`INSERT INTO post_folders (post_id, folder_id) VALUES (?, ?)`, post.Id, folder_id); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit();
	return post.Id, nil
}

func Update(post_id int, title string, content string, folder_id int) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	post := Post{
		Id: post_id,
		Title: title,
		Content: content,
		Updated: time.Now(),
	}

	if _, err := tx.Update(&post); err != nil {
		tx.Rollback()
		return 0, err
	}

	if folder_id != 0 {
		if !CheckFolder(post_id) {
			_, err = tx.Exec(`INSERT INTO post_folders (post_id, folder_id) VALUES (?, ?)`, post_id, folder_id)
		}

		_, err = tx.Exec(`UPDATE post_folders SET folder_id = ? WHERE post_id = ?`, folder_id, post_id)
		if err !=nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit();
	return post.Id, nil
}

func DeleteOne(post_id int) error {
	tx, _ := db.Begin()
	post := Post{
		Id: post_id,
	}

	_, err := tx.Delete(&post)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(`DELETE FROM post_folders WHERE post_id = ?`, post_id)
	if err != nil {
		log.Fatal(err)
	}

	return tx.Commit()
}

func CheckFolder(post_id int) bool {
	var select_id,_ = db.SelectInt(`SELECT post_id FROM post_folders WHERE post_id = ?`, post_id)

	return select_id != 0
}

func SearchForWord(word string) []Post {
	var posts []Post
	search_word := `%` + word + `%`

	db.Select(&posts, `SELECT post_id, title, content, created_at, updated_at FROM posts
	WHERE title LIKE ? OR content LIKE ? ORDER BY post_id`,
		search_word, search_word,
	)

	return posts
}

func ByFolderId(folder_id int) []Post {
	var posts []Post

	db.Select(&posts, `SELECT p.post_id, title, content, created_at, updated_at FROM posts AS p
	JOIN post_folders USING(post_id) WHERE folder_id = ?`, folder_id)

	return posts
}

