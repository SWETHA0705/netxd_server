package main

import (
	"context"
	"fmt"
	"net"

	netxddalservices "github.com/SWETHA0705/netxd_dal/netxd_dal_services"
	netxdcustomerconfig "github.com/SWETHA0705/netxd_server/netxd_customer_config"
	netxdcustomerconstants "github.com/SWETHA0705/netxd_server/netxd_customer_constants"

	controller "github.com/SWETHA0705/netxd_server/netxd_customer_controller"

	netxdcustomer "github.com/SWETHA0705/netxd_customer/customer"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	//"golang.org/x/vuln/client"
)

func InitialiseDatabase(client *mongo.Client){
	customercollection := netxdcustomerconfig.GetCollection(client,"bankdb","customer")

	controller.CustomerService = netxddalservices.InitialiseCustomerService(customercollection,context.Background())


}
func main() {

	mongoclient, err := netxdcustomerconfig.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	
	InitialiseDatabase(mongoclient)
	

	lis, err := net.Listen("tcp", netxdcustomerconstants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	fmt.Println("ser")

	s := grpc.NewServer()
	netxdcustomer.RegisterCustomerServiceServer(s, &controller.CustomerServer{})

	fmt.Println("Server listening on", netxdcustomerconstants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
     

	
}