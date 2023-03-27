package dao

import "tieu-minh/pkg/db"

type Repo struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (repo *Repo) GetList() error {
	query := `select id, name from public.todo;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var todo = &Todo{}
	for rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Name)
		if err != nil {
			return err
		}
		repo.Todos = append(repo.Todos, *todo)
	}
	return nil
}

func (todo *Todo) Insert() error {
	query := `INSERT INTO public.todo (id, "name") VALUES($1, $2);`
	_, err := db.PSQL.Exec(query, todo.Id, todo.Name)
	return err
}

func (todo *Todo) Delete() error {
	query := `delete from public.todo where id = $1;`
	_, err := db.PSQL.Exec(query, todo.Id)
	return err
}
