package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	var err error
	databaseUrl := "postgres://postgres:12345gendi@localhost:5432/personal_web"

	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
		os.Exit(1)
	}
	fmt.Println("Success connect to database")
}

// postgres://{user}:{password}@{host}:{port}/{database}
// user = user postgres
// password = password postgres
// host = host postgres
// port = port postgres
// database = database postgres
// INSERT INTO public.tb_projects(
// 	id, projectname, sdate, edate, description, technologies, image, author_id)
// 	VALUES (?, ?, ?, ?, ?, ?, ?, ?);

// 	SELECT id, projectname, sdate, edate, description, technologies, image, author_id
// 	FROM public.tb_projects
// 	SELECT tb_projects.id, projectname, sdate, edate, description, technologies, image, tb_user.name as author
// 	FROM tb_user
// 	LEFT JOIN tb_projects ON tb_user.id = tb_projects.author_id ORDER BY id DESC

// 	SELECT tb_projects.id, projectname, tb_user.name as author
// 	FROM tb_projects
// 	INNER JOIN tb_user ON tb_projects.author_id = tb_user.id ORDER BY id DESC

// 	SELECT *
// 	FROM tb_projects
// 	FULL OUTER JOIN tb_user
// 	ON tb_projects.author_id = tb_user.id
