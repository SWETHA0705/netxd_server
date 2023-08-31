package netxdcustomercontroller

import ( //netxddalinterfaces "github.com/SWETHA0705/netxd_dal/netxd_dal_interfaces"
	"context"
	"fmt"

	pro "github.com/SWETHA0705/netxd_customer/customer"
	netxddalinterfaces "github.com/SWETHA0705/netxd_dal/netxd_dal_interfaces"
)

type TransactionServer struct{
	pro.UnimplementedCustomerServiceServer
}

var(
	TransactionService netxddalinterfaces.Itransaction
)

func (s*TransactionServer)Transaction(ctx context.Context,req * pro.Transaction)(*pro.TransactionResponse,error){
	fmt.Println(req.FromAccount)
	
	res,err := TransactionService.Transaction(req.FromAccount,req.ToAccount,req.Amount)
	if err!= nil{
		fmt.Println("s")
	  return nil,err
	}else{
	  Response := &pro.TransactionResponse{
	  	Message: res,
	  }
	  return Response,nil
	}
}
