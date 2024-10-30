package main

import (
	"fall-cloud/database"
	"fall-cloud/routes"
)



func main()  {
	//docker.ListAllContainerByID()
	//containers := docker.ListAllRunningByID()
	//docker.StopAll(containers)
	//docker.DeleteById("2b04140257d3")
	//docker.CreateContainer()
	database.StartDB()
	routes.HandleRequest()
}