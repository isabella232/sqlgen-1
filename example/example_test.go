package example

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func initType(t *testing.T, db *sql.DB, filename string) {
	stmt, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(stmt))
	if err != nil {
		t.Logf("IGNORE: %s", err)
	}
}

func TestCustom(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres dbname=custom sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	initType(t, db, "money.sql")
	initType(t, db, "job.sql")

	_, err = db.Exec(`create table if not exists transactions
		(id text, job job_type)`)
	if err != nil {
		t.Fatal(err)
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	j := Job{
		ID: id,
		Amount: &Money{
			Amount:   100,
			Currency: `hello "world"`,
			Rounded:  true,
		},
		Name: "Test User",
	}
	_, err = db.Exec(`insert into transactions (id, job) values ($1, $2)`, id, &j)
	if err != nil {
		resp, _ := j.Value()
		t.Fatalf("error: %s\n%s", err, resp)
	}

	j2 := Job{}
	err = db.QueryRow(`select job from transactions where id=$1`, id).Scan(&j2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, j, j2)
}

func TestJobArray(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres dbname=custom sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	initType(t, db, "money.sql")
	initType(t, db, "job.sql")

	_, err = db.Exec(`create table if not exists jobs (id text, jobs job_type[])`)
	if err != nil {
		t.Fatal(err)
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	jobs := JobArray{
		{id, &Money{12, "EUR", false}, "Test Job", 1},
		{id + "1", &Money{10, "EUR", false}, "Test Job", 2},
	}

	_, err = db.Exec(`insert into jobs (id, jobs) values ($1, $2)`, id, jobs)
	if err != nil {
		t.Fatal(err)
	}

	var other JobArray
	err = db.QueryRow(`select jobs from jobs where id=$1`, id).Scan(&other)
	if err != nil {
		t.Fatal(err)
	}

	for i, m := range other {
		assert.Equal(t, jobs[i], m)
	}
}

func TestComplete(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres dbname=custom sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	initType(t, db, "complete.sql")

	_, err = db.Exec(`create table if not exists alltypes
		(id text, complete complete_type)`)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	id := fmt.Sprintf("%d", now.UnixNano())
	c := Complete{
		CustomString:  "hello world",
		CustomInt16:   1,
		CustomInt32:   2,
		CustomInt:     3,
		CustomInt64:   4,
		CustomUint16:  5,
		CustomUint32:  6,
		CustomUint:    9,
		CustomUint64:  8,
		CustomFloat32: 12.1,
		CustomFloat64: 12.12,
	}
	_, err = db.Exec(`insert into alltypes (id, complete) values ($1, $2)`, id, &c)
	if err != nil {
		resp, _ := c.Value()
		t.Fatalf("error: %s\n%s", err, resp)
	}

	c2 := Complete{}
	err = db.QueryRow(`select complete from alltypes where id=$1`, id).Scan(&c2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, c, c2)

}

func TestTimes(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres dbname=custom sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	initType(t, db, "times.sql")

	_, err = db.Exec(`create table if not exists times
		(id text, times times_type)`)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	id := fmt.Sprintf("%d", now.UnixNano())
	c := Times{
		CustomTime:  now,
		CustomTimep: &now,
	}
	_, err = db.Exec(`insert into times (id, times) values ($1, $2)`, id, &c)
	if err != nil {
		resp, _ := c.Value()
		t.Fatalf("error: %s\n%s", err, resp)
	}

	c2 := Times{}
	err = db.QueryRow(`select times from times where id=$1`, id).Scan(&c2)
	if err != nil {
		t.Fatal(err)
	}

	// Time has a small problem as the parsed time is always in utc
	if c2.CustomTime.Unix() != c.CustomTime.Unix() {
		t.Fatalf("exp: %s got: %s", c, c2)
	}
	if c2.CustomTimep.Unix() != c.CustomTimep.Unix() {
		t.Fatalf("exp: %s got: %s", c, c2)
	}
}
