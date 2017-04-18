package main

import (
  // "bytes"
  "database/sql"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/itsjamie/gin-cors"
  _ "github.com/mattn/go-sqlite3"
  "net/http"
  "sort"
  "strconv"
  "strings"
)

func main() {
  db, err := sql.Open("sqlite3", "./task")
  checkErr(err)
  defer db.Close()
  type Country struct {
    Id   int
    Name string
    Code string
  }
  type State struct {
    Id         int
    Name       string
    Code       string
    Country_Id int
  }
  type City struct {
    Id       int
    Name     string
    Code     string
    State_Id int
  }
  type Distrubutor struct {
    Id             int
    Name           string
    Distrubutor_Id int
  }
  router := gin.Default()
  router.Use(cors.Middleware(cors.Config{
    Origins:         "*",
    ValidateHeaders: false,
    RequestHeaders:  "Origin, Authorization, Content-Type",
  }))
  // Get all countries
  router.GET("/countries", func(c *gin.Context) {
    var (
      country   Country
      countries []Country
    )
    rows, err := db.Query("select id, name, code from countries;")
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&country.Id, &country.Name, &country.Code)
      countries = append(countries, country)
      checkErr(err)
    }
    defer rows.Close()
    c.JSON(http.StatusOK, gin.H{
      "countries": countries,
    })
  })
  // Create countries
  router.POST("/countries", func(c *gin.Context) {
    name := c.PostForm("name")
    code := c.PostForm("code")
    stmt, err := db.Prepare("insert into countries (name, code) values(?,?);")
    checkErr(err)
    _, err = stmt.Exec(name, code)

    checkErr(err)

    defer stmt.Close()
    c.JSON(http.StatusOK, gin.H{
      "message": fmt.Sprintf(" %s, %s successfully created", name, code),
    })
  })
  // Get country and its states with id
  router.GET("/countries/:id", func(c *gin.Context) {
    var (
      country Country
      state   State
      states  []State
      result  gin.H
    )
    id := c.Param("id")
    row := db.QueryRow("select id, name, code from countries where id = ?;", id)
    err = row.Scan(&country.Id, &country.Name, &country.Code)
    if err != nil {
      // If no results send null
      result = gin.H{
        "country": nil,
      }
    } else {
      rows, err := db.Query("select id, name, code,country_id from states where country_id=?;", id)
      checkErr(err)
      for rows.Next() {
        err = rows.Scan(&state.Id, &state.Name, &state.Code, &state.Country_Id)
        states = append(states, state)
        checkErr(err)
      }
      defer rows.Close()
      result = gin.H{
        "country": country,
        "states":  states,
      }
    }
    c.JSON(http.StatusOK, result)
  })
  router.GET("/states", func(c *gin.Context) {
    var (
      state  State
      states []State
    )
    rows, err := db.Query("select id, name, code, country_id from states;")
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&state.Id, &state.Name, &state.Code, &state.Country_Id)
      states = append(states, state)
      checkErr(err)
    }
    defer rows.Close()
    c.JSON(http.StatusOK, gin.H{
      "states": states,
    })
  })
  // Create states
  router.POST("/states", func(c *gin.Context) {
    name := c.PostForm("name")
    code := c.PostForm("code")
    country_id := c.PostForm("country_id")
    stmt, err := db.Prepare("insert into states (name, code, country_id) values(?,?,?);")
    checkErr(err)
    _, err = stmt.Exec(name, code, country_id)

    checkErr(err)

    defer stmt.Close()
    c.JSON(http.StatusOK, gin.H{
      "message": fmt.Sprintf(" %s, %s successfully created", name, code),
    })
  })
  // Get state and its cities with id
  router.GET("/states/:id", func(c *gin.Context) {
    var (
      country Country
      state   State
      city    City
      cities  []City
      result  gin.H
    )
    id := c.Param("id")
    row := db.QueryRow("select id, name, code, country_id from states where id = ?;", id)
    err = row.Scan(&state.Id, &state.Name, &state.Code, &state.Country_Id)
    if err != nil {
      // If no results send null
      result = gin.H{
        "state": nil,
      }
    } else {
      row := db.QueryRow("select id, name, code from countries where id = ?;", state.Country_Id)
      _ = row.Scan(&country.Id, &country.Name, &country.Code)

      rows, err := db.Query("select id, name, code,state_id from cities where state_id=?;", id)
      checkErr(err)
      for rows.Next() {
        err = rows.Scan(&city.Id, &city.Name, &city.Code, &city.State_Id)
        cities = append(cities, city)
        checkErr(err)
      }
      defer rows.Close()
      result = gin.H{
        "country": country,
        "state":   state,
        "cities":  cities,
      }
    }
    c.JSON(http.StatusOK, result)
  })
  router.GET("/cities", func(c *gin.Context) {
    var (
      city   City
      cities []City
    )
    rows, err := db.Query("select id, name, code, state_id from cities;")
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&city.Id, &city.Name, &city.Code, &city.State_Id)
      cities = append(cities, city)
      checkErr(err)
    }
    defer rows.Close()
    c.JSON(http.StatusOK, gin.H{
      "cities": cities,
    })
  })
  // Create cities
  router.POST("/cities", func(c *gin.Context) {
    name := c.PostForm("name")
    code := c.PostForm("code")
    state_id := c.PostForm("state_id")
    stmt, err := db.Prepare("insert into cities (name, code, state_id) values(?,?,?);")
    checkErr(err)
    _, err = stmt.Exec(name, code, state_id)

    checkErr(err)

    defer stmt.Close()
    c.JSON(http.StatusOK, gin.H{
      "message": fmt.Sprintf(" %s, %s successfully created", name, code),
    })
  })
  // get all distributors
  router.GET("/distributors", func(c *gin.Context) {
    var (
      distrubutor  Distrubutor
      distrubutors []Distrubutor
    )
    rows, err := db.Query("select id, name, distributor_id from distributors;")
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&distrubutor.Id, &distrubutor.Name, &distrubutor.Distrubutor_Id)
      checkErr(err)
      distrubutors = append(distrubutors, distrubutor)
    }
    defer rows.Close()
    c.JSON(http.StatusOK, gin.H{
      "distrubutors": distrubutors,
    })
  })
  // Create distributors
  router.POST("/distributors", func(c *gin.Context) {
    save := true
    name := c.PostForm("name")
    distributor_id := c.PostForm("distributor_id")
    included_countries := strings.Split(c.PostForm("included_countries"), ",")
    included_states := strings.Split(c.PostForm("included_states"), ",")
    included_cities := strings.Split(c.PostForm("included_cities"), ",")
    excluded_states := strings.Split(c.PostForm("excluded_states"), ",")
    excluded_cities := strings.Split(c.PostForm("excluded_cities"), ",")
    if distributor_id != "" {
      save = check_includion_exclution(db, distributor_id, included_countries, included_states, included_cities)
    } else {
      distributor_id = "0"
    }
    if save {
      stmt, err := db.Prepare("insert into distributors (name, distributor_id) values(?,?);")
      checkErr(err)
      res, err := stmt.Exec(name, distributor_id)
      checkErr(err)
      dist_id, _ := res.LastInsertId()
      add_includion_exclution(db, "included_countries", "countries", "country_id", dist_id, included_countries)
      add_includion_exclution(db, "included_states", "states", "state_id", dist_id, included_states)
      add_includion_exclution(db, "included_cities", "cities", "city_id", dist_id, included_cities)
      add_includion_exclution(db, "excluded_states", "states", "state_id", dist_id, excluded_states)
      add_includion_exclution(db, "excluded_cities", "cities", "city_id", dist_id, excluded_cities)

      defer stmt.Close()
      c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf(" %s, %s, successfully created", name, distributor_id),
      })
    } else {
      c.JSON(http.StatusOK, gin.H{
        "message": "distribution permision issue for sub-distrubutor",
      })
    }
  })
  // get permisions
  router.POST("/permisions", func(c *gin.Context) {
    distributor_id := c.PostForm("distributor_id")
    country_ids := strings.Split(c.PostForm("country_ids"), ",")
    state_ids := strings.Split(c.PostForm("state_ids"), ",")
    city_ids := strings.Split(c.PostForm("city_ids"), ",")
    message := ""
    if check_includion_exclution(db, distributor_id, country_ids, state_ids, city_ids) {
      message = "distributor has permisions"
    } else {
      message = "distributor doesnot have permisions"
    }
    c.JSON(http.StatusOK, gin.H{
      "message": message,
    })
  })
  // Add API handlers here
  router.Run(":3000")
}

func check_includion_exclution(db *sql.DB, distributor_id string, country_ids []string, state_ids []string, city_ids []string) bool {
  return_obj := true
  if len(country_ids) > 0 && !validate_distrubutor_country(db, distributor_id, country_ids) {
    return_obj = false
  }
  if len(state_ids) > 0 && !validate_distrubutor_state(db, distributor_id, state_ids) {
    return_obj = false
  }
  if len(city_ids) > 0 && !validate_distrubutor_city(db, distributor_id, city_ids) {
    return_obj = false
  }
  return return_obj
}
func validate_distrubutor_country(db *sql.DB, distributor_id string, country_ids []string) bool {
  country_ids_string := strings.Join(country_ids, ", ")
  var id_count int
  row := db.QueryRow("select count(id) from included_countries where country_id in (?) and distributor_id = ?;", country_ids_string, distributor_id)
  err := row.Scan(&id_count)
  checkErr(err)
  return id_count == len(country_ids)
}
func validate_distrubutor_state(db *sql.DB, distributor_id string, state_ids []string) bool {
  return_obj := false
  state_ids_string := strings.Join(state_ids, ", ")
  var (
    id_count          int
    inc_state_id      int64
    inc_country_id    int64
    inc_country_ids   []string
    not_inc_state_ids []string
  )
  row := db.QueryRow("select count(id) from excluded_states where state_id in (?) and distributor_id = ?;", state_ids_string, distributor_id)
  err := row.Scan(&id_count)
  checkErr(err)
  if id_count == 0 {
    rows, err := db.Query("select state_id from included_states where state_id in (?) and distributor_id = ?;", state_ids_string, distributor_id)
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&inc_state_id)
      checkErr(err)
      if !find(state_ids, strconv.FormatInt(inc_state_id, 10)) {
        not_inc_state_ids = append(not_inc_state_ids, strconv.FormatInt(inc_state_id, 10))
      }
    }
    not_inc_state_ids_string := strings.Join(not_inc_state_ids, ", ")
    rows, err = db.Query("select distinct country_id from states where id in (?);", not_inc_state_ids_string)
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&inc_country_id)
      checkErr(err)
      inc_country_ids = append(inc_country_ids, strconv.FormatInt(inc_country_id, 10))
    }
    return_obj = validate_distrubutor_country(db, distributor_id, inc_country_ids)
  }
  return return_obj
}
func validate_distrubutor_city(db *sql.DB, distributor_id string, city_ids []string) bool {
  city_ids_string := strings.Join(city_ids, ", ")
  var (
    id_count         int
    inc_city_id      int64
    inc_state_id     int64
    inc_state_ids    []string
    not_inc_city_ids []string
  )
  row := db.QueryRow("select count(id) from excluded_cities where city_id in (?) and distributor_id = ?;", city_ids_string, distributor_id)
  err := row.Scan(&id_count)
  checkErr(err)
  if id_count == 0 {
    rows, err := db.Query("select city_id from included_cities where city_id in (?) and distributor_id = ?;", city_ids_string, distributor_id)
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&inc_city_id)
      checkErr(err)
      if !find(city_ids, strconv.FormatInt(inc_city_id, 10)) {
        not_inc_city_ids = append(not_inc_city_ids, strconv.FormatInt(inc_city_id, 10))
      }
    }
    not_inc_city_ids_string := strings.Join(not_inc_city_ids, ", ")
    rows, err = db.Query("select distinct state_id from cities where id in (?);", not_inc_city_ids_string)
    checkErr(err)
    for rows.Next() {
      err = rows.Scan(&inc_state_id)
      checkErr(err)
      inc_state_ids = append(inc_state_ids, strconv.FormatInt(inc_state_id, 10))
    }
    return validate_distrubutor_state(db, distributor_id, inc_state_ids)
  } else {
    return false
  }
}
func add_includion_exclution(db *sql.DB, table string, check_table string, column string, distributor_id int64, regions []string) {
  ins, err := db.Prepare(fmt.Sprintf("insert into %s (%s, distributor_id) values(?,?);", table, column))
  checkErr(err)
  for _, region := range regions {
    if region != "" {
      _, err = ins.Exec(region, distributor_id)
      checkErr(err)
    }
  }
}
func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
func find(elements []string, element string) bool {
  i := sort.Search(len(elements),
    func(i int) bool { return elements[i] >= element })
  return i < len(elements) && elements[i] == element
}
