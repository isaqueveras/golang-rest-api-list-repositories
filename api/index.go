package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Repository contains the details of a repository
type repositorySymmary struct {
	ID         int
	Name       string
	Owner      string
	TotalStars int
}

type repositories struct {
	Repositories []repositorySymmary
}

var db *sql.DB

// IndexHandler calls `queryRepos()` and marshals the result as JSON
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	repos := repositories{}

	err := queryRepos(&repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

// queryRepos first fetches the repositories data from the db
func queryRepos(repos *repositories) error {
	rows, err := db.Query(`SELECT id, repository_owner, repository_name, total_stars FROM repositories ORDER BY total_starts DESC`)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		repo := repositorySymmary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Owner,
			&repo.Name,
			&repo.TotalStars,
		)

		if err != nil {
			return err
		}

		repos.Repositories = append(repos.Repositories, repo)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
