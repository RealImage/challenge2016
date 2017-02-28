package main

import (
	db "./db"
	"github.com/go-macaron/session"
    "github.com/go-macaron/pongo2"
    "gopkg.in/macaron.v1"
    "fmt"
    "encoding/json"
    "strconv"
)

func main() {
    fmt.Println("Useless statement to ignore fmt unued error!")
    m := macaron.Classic()
    m.Use(macaron.Static("public"))
    
    //Tell macoron to use Pongoer template engine
    m.Use(pongo2.Pongoer(pongo2.Options{
        Directory: "views",
        IndentJSON: true,
    }))

    m.Use(func(ctx *macaron.Context) {
	    ctx.Header().Set("Access-Control-Allow-Origin", "*")
	    ctx.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")	    
	})

    //Use Macaron Session
    m.Use(session.Sessioner())

    //This will read the CSV file and store it in DB
    db.SeedPlaceDB()
    db.InitDB()    

    //Render homePage
    m.Get("/", homePage)

    //Get all places
    m.Get("/places", listPlaces) 

    //Search for autocomplete
    m.Get("/places/search/:type/:q", searchPlaces)

    //Get all places
    m.Get("/distributors", listDistributors)

    //Get all places
    m.Post("/distributors", createDistributor)

    //Options
    m.Options("/distributors", options)    

    //Add place to a distributor includes
    m.Get("/distributors/:id/include/:code/:name", includePlace)

    //Add place to a distributor excludes
    m.Get("/distributors/:id/exclude/:code/:name", excludePlace)

    m.Run()
}

func includePlace(ctx *macaron.Context, sess session.Store) {
	i, _ := strconv.Atoi(ctx.Params(":id"))

	if(db.AddPlace(sess.ID(), i, ctx.Params(":code"), "include", ctx.Params(":name"))){		
		ctx.JSON(200, map[string]interface{}{"data": []interface{}{db.SerializeDistributor(db.FindDistributor(db.DistributorDB[sess.ID()], i))}})
	}else{
		ctx.Status(422)
	}
	
}

func excludePlace(ctx *macaron.Context, sess session.Store) {
	i, _ := strconv.Atoi(ctx.Params(":id"))

	if(db.AddPlace(sess.ID(), i, ctx.Params(":code"), "exclude", ctx.Params(":name"))){
		ctx.JSON(200, map[string]interface{}{"data": []interface{}{db.SerializeDistributor(db.FindDistributor(db.DistributorDB[sess.ID()], i))}})
	}else{
		ctx.Status(422)
	}
	
}

func searchPlaces(ctx *macaron.Context, sess session.Store) {	
	ctx.JSON(200, db.SearchPlace(ctx.Params(":q"), ctx.Params(":type")))
}

func options(ctx *macaron.Context, sess session.Store) {
	ctx.Status(200)
}

func createDistributor(ctx *macaron.Context, sess session.Store) {
	body, err := ctx.Req.Body().Bytes();
	if err != nil {
        fmt.Println("Error")
    }

    //TODO: Needs to be serialized in a better way
	var raw map[string]interface{}
    json.Unmarshal(body, &raw)
	
	var data map[string]interface{}
	data = raw["data"].(map[string]interface{})

	var attributes map[string]interface{}
	attributes = data["attributes"].(map[string]interface{})
	
	var id int = db.NextId()
	parentDistributorId := int(attributes["parentDistributorId"].(float64))

	var distributor db.Distributor = db.Distributor{id, attributes["name"].(string), parentDistributorId, []string{}, []string{}, []string{}, []string{}}

	db.DistributorDB[sess.ID()] = append(db.DistributorDB[sess.ID()], distributor)

	ctx.JSON(201, map[string]interface{}{"data": db.SerializeDistributor(distributor)})
}

func homePage(ctx *macaron.Context, sess session.Store) {	
    ctx.HTML(200, "index")
}

func listPlaces(ctx *macaron.Context, sess session.Store) {
	ctx.JSON(200, db.SerialzeAllPlace(db.PlaceDB))
}

func listDistributors(ctx *macaron.Context, sess session.Store) {
	db.SeedDBForUser(sess.ID())
	fmt.Printf("%+v\n", db.DistributorDB[sess.ID()])
	ctx.JSON(200, db.SerialzeAllDistributor(db.DistributorDB[sess.ID()]))
}