package folder

import (
	"time"
	"github.com/goshepherd/goshepherd/app/model"
	"database/sql"
	"log"
)


var db = model.InitDb()

func init() {
	db.AddTableWithName(Folder{}, "folders").SetKeys(true, "Id")
}

type Folder struct {
	Id int `db:"folder_id"`
	Name string `db:"folder_name"`
	Created time.Time `db:"created_at"`
	Updated time.Time `db:"updated_at"`
}

type PostFolder struct {
	Folder
	PostCount sql.NullInt64 `db:"count"`
}

func All() []Folder {
	var folders []Folder

	db.Select(&folders, `SELECT folder_id, folder_name, created_at, updated_at FROM folders
	ORDER BY updated_at, folder_id`)


	return folders
}

func AllWithCount() []PostFolder {
	var folders []PostFolder

	db.Select(&folders, `SELECT f.folder_id, folder_name, created_at, updated_at, count FROM folders AS f
	LEFT JOIN (SELECT count(*) as count, folder_id FROM post_folders GROUP BY folder_id) AS pf USING (folder_id)
	ORDER BY pf.count, updated_at, folder_id`)

	return folders
}

func search(word string) []Folder  {
	var folders []Folder
	db.Select(&folders, `SELECT folder_id FROM folders WHERE folder_name LIKE ?`, word)

	return folders
}

func Add(name string) error {
	folder := Folder{
		Name: name,
		Created: time.Now(),
		Updated: time.Now(),
	}

	err := db.Insert(&folder)
	return err
}

func PickOne(folder_id int) Folder {
	var folder Folder

	db.SelectOne(&folder, `SELECR folder_id, folder_name, created_at, updated_at FROM folders
	WHERE folder_id = ?
	ORDER BY updated_at DESC, folder_id DESC`, folder_id)

	return folder
}

func DeleteOne(folder_id int) error {
	tx, _ := db.Begin()
	_, err := tx.Exec(`DELETE FROM post_folders WHERE folder_id = ?`, folder_id)
	if err != nil {
		log.Fatal(err)
	}

	folder := Folder{
		Id: folder_id,
	}

	_, err = tx.Delete(&folder)
	if err != nil {
		log.Fatal(err)
	}

	return tx.Commit()
}
