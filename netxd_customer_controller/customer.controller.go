package netxdcustomercontroller

import (
	// "dal/netxd_dal_models"
	"context"
	netxddalinterfaces "netxd_project/dal/netxd_dal/netxd_dal/interfaces"
	"netxd_project/dal/netxd_dal/netxd_dal_models"
	pro "netxd_project/netxd_customer/customer"
)

type CustomerServer struct{
 pro.UnimplementedCustomerServiceServer
}

var(
  CustomerService netxddalinterfaces.ICustomer

)

func (c * CustomerServer) CreateCustomer(ctx context.Context,req * pro.Customer)(*pro.CustomerResponse,error){
  dbcustomer := &netxddal.Customer{FirstName : req.FirstName,
  LastName:  req.LastName,
  BankId: int(req.BankId),
  Balance: int(req.Balance),
}
res,err := CustomerService.CreateCustomer(dbcustomer)
if err!= nil{
  return nil,err
}else{
  Response := &pro.CustomerResponse{CustomerId: int64(res.CustomerId),
    CreatedAt: res.CreatedAt,
  }
  return Response,nil
}
}

