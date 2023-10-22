package main

import (
	"fmt"
)






func main() {
	// Load region data from CSV
    regions := LoadRegionData()

    // Create the Gin router and API routes
    r := SetupRouter(regions)

    // Start the web server
    if err := r.Run(":8080"); err != nil {
        fmt.Println(err)
    }
}
