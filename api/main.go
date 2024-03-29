/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.17
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"encoding/json"
	"log"

	// WARNING!
	// Pass --git-repo-id and --git-user-id properties when generating the code
	//

	"github.com/shivinkapur/sample-go-api/api/entities"
	sw "github.com/shivinkapur/sample-go-api/api/go"

	_ "embed"
)

//go:embed version.json
var bytesVersion []byte

func main() {

	err := json.Unmarshal(bytesVersion, &entities.VERSION)

	if err != nil {
		log.Fatal(err)
	}

	routes := sw.ApiHandleFunctions{}

	log.Printf("Server started")

	router := sw.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
