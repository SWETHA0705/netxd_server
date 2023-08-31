package main

import (
	"context"
	"fmt"
	"net"
	netxddalservices "netxd_project/dal/netxd_dal/netxd_dal_services"
	netxdcustomerconfig "netxd_project/server/netxd_server/netxd_customer_config"
	netxdcustomerconstants "netxd_project/server/netxd_server/netxd_customer_constants"

	controller "netxd_project/server/netxd_server/netxd_customer_controller"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
netxdcustomer	"netxd_project/netxd_customer/customer"
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