package database

import "github.com/jmoiron/sqlx"

const max_db_conns = 20

var Connection = &connection{}

type connection struct {
	userDb     *sqlx.DB
	messageDb  *sqlx.DB
	stepRankDb *sqlx.DB
}

func (c *connection) OpenUserDb() *sqlx.DB {
	db, err := sqlx.Open(
		"mysql",
		"dummy",
		// env.Values.DbUser+":"+env.Values.DbPass+"@tcp("+env.Values.DbHost+":"+env.Values.DbPort+")/"+env.Values.DbName,
	)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(max_db_conns)
	c.userDb = db

	return db
}

func (c *connection) OpenMessageDb() *sqlx.DB {
	db, err := sqlx.Open(
		"mysql",
		"dummy",
		// env.Values.MsgDbUser+":"+env.Values.MsgDbPass+"@tcp("+env.Values.MsgDbHost+":"+env.Values.MsgDbPort+")/"+env.Values.MsgDbName,
	)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(max_db_conns)
	c.messageDb = db

	return db
}

func (c *connection) OpenStepRankeDb() *sqlx.DB {
	db, err := sqlx.Open(
		"mysql",
		"dummy",
		// env.Values.StepRankDbUser+":"+env.Values.StepRankDbPass+"@tcp("+env.Values.StepRankDbHost+":"+env.Values.StepRankDbPort+")/"+env.Values.StepRankDbName,
	)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(max_db_conns)
	c.stepRankDb = db

	return db
}
