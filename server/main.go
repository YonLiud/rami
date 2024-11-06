package server

import (
	"rami/database"
)

func main() {
	database.InitDB("rami.db")

}
