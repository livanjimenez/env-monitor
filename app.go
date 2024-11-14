package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {}
func (a *App) Run(addr string)                          {}
