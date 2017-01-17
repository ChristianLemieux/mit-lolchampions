package champion

import (
	"../../data"
	"../../objects"
	"database/sql"
	"log"
)

func GetChampion(id string) *objects.Champion {
	e := data.GetQueryEngine()
	rows, err := e.Query("select * from champions where id = ? limit 1;", id)

	if err != nil {
		log.Printf("There was an issue fetching the champion: %q", err)
		return nil
	}
	defer rows.Close()
	return rowToChampion(rows)
}

func GetAllChampions() []*objects.Champion {
	e := data.GetQueryEngine()
	rows, err := e.Query("select * from champions", nil)
	champions := make([]*objects.Champion, 0)

	if err != nil {
		log.Printf("There was an issue fetching the list of champions: %q", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		champions = append(champions, scanChampion(rows))
	}
	return champions
}

func rowToChampion(row *sql.Rows) *objects.Champion {
	if row.Next() {
		return scanChampion(row)
	}
	return nil
}

// Convert a SQL row object into a champion object
func scanChampion(row *sql.Rows) *objects.Champion {
	champion := new(objects.Champion)
	err := row.Scan(
		&champion.Id,
		&champion.Name,
		&champion.Title,
		&champion.Blurb,
		&champion.Attack,
		&champion.Defense,
		&champion.Magic,
		&champion.Difficulty,
		&champion.Image,
		&champion.Tag,
	)
	if err != nil {
		log.Printf("There was an issue scanning the champion results: %q", err)
		return nil
	}
	return champion
}
