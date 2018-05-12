package main

import (
	"cloud/config"
	"flag"
	"fmt"
	"os"
)

func main() {

	buildCommand := flag.NewFlagSet("build", flag.ExitOnError)
	clusterName := buildCommand.String("name", "", "mongo cluster name (Required) ")
	noOfClusterNodes := buildCommand.Int("nodes", 3, "No of nodes in the cluster")
	clusterSecurity := buildCommand.String("security", "", "security type of the cluster")
	clusterPort := buildCommand.Int("port", 27017, "mongo server port")
	clusterType := buildCommand.String("type", "replicaset", "Mongo cluster type")

	if len(os.Args) < 2 {
		fmt.Println("subcommand build is required")
		os.Exit(0)
	} else {
		buildCommand.Parse(os.Args[2:])
	}

	if buildCommand.Parsed() {
		// Required Flags
		if *clusterName == "" {
			buildCommand.PrintDefaults()
			os.Exit(1)
		}
		/* validate cluster type and restrict choice */
		clusterTypeChoices := map[string]bool{"replicaset": true, "shard": true, "standalone": false}
		if _, validChoice := clusterTypeChoices[*clusterType]; !validChoice {
			//buildCommand.PrintDefaults()
			fmt.Println("Valid build types are: replicaset, shard")
			os.Exit(1)
		}

		fmt.Printf("clusterName: %s, members: %v, security: %s, port:%v \n", *clusterName, *noOfClusterNodes, *clusterSecurity, *clusterPort)
		if *noOfClusterNodes%2 == 0 {
			fmt.Println("Cluster members must be and odd value")
			os.Exit(1)
		}
	}

	var c config.Configuration = config.GetConfig()
	fmt.Println(c.LogPath)
	fmt.Println(c.RepoURL["Darwin"])

}
