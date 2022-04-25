/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type student struct {
	Name string
	Dept string
}

// startServerxCmd represents the startServerx command
var startServerxCmd = &cobra.Command{
	Use:   "startServerx",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Serverx Started..................")
		http.HandleFunc("/", get_student)
		http.ListenAndServe(":8010", nil)

	},
}

func init() {
	rootCmd.AddCommand(startServerxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startServerxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startServerxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func get_student(w http.ResponseWriter, req *http.Request) {
	fmt.Println("called student")
	switch req.Method {
	case "GET":
		st := student{}
		fmt.Fprintln(w, "got get method")
		db := dbConn()
		l := len(req.URL.Query())

		if l < 1 {
			sqlDB, err := db.Query("SELECT * FROM students")
			if err != nil {
				panic(err)
			}
			for sqlDB.Next() {
				var name, dept string
				err := sqlDB.Scan(&name, &dept)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(name, dept)
				fmt.Fprintln(w, name, dept)

				st.Name = name
				st.Dept = dept

			}

		} else {
			variable := req.URL.Query().Get("name")

			query := fmt.Sprintf("Select * from students where name='%v'or dept='%v'", variable, variable)

			sqlDB, err := db.Query(query)
			if err != nil {
				panic(err)
			}
			for sqlDB.Next() {
				var name, dept string
				err := sqlDB.Scan(&name, &dept)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(name, dept)
				fmt.Fprintln(w, name, dept)

				st.Name = name
				st.Dept = dept

			}

		}

	case "POST":
		db := dbConn()
		l := len(req.URL.Query())
		if l == 2 {
			name := req.URL.Query().Get("name")
			dept := req.URL.Query().Get("dept")
			fmt.Fprintln(w, "got post method with parameters", name, dept)
			query := fmt.Sprintf("INSERT into students values('%v','%v')", name, dept)
			_, err := db.Query(query)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Fprintln(w, "Invalid number of parameters to insert data")
		}

	case "PUT":
		db := dbConn()
		l := len(req.URL.Query())
		if l < 2 {
			fmt.Fprintln(w, "Invalid number of parameters to insert data")
		} else {
			name := req.URL.Query().Get("name")
			dept := req.URL.Query().Get("dept")

			query := fmt.Sprintf("update students set dept='%v' where name='%v')", dept, name)
			_, err := db.Query(query)
			if err != nil {
				panic(err)
			}

		}
		fmt.Fprintln(w, "got put method")

	}

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "students"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
