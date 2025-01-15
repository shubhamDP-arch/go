package main

// type Server struct {
// 	accountClient
// 	catalogUrl
//		orderUrl
//}

func newGraphQLServer(accountUrl, catalogUrl, orderUrl) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {

		return nil, error
	}
	catalogClient, err := catalog.NewClient(catalogUrl)
	if err != nil {
		accountClient.Close()
		return nil, error
	}
	orderClient, err := order.NewClient(orderUrl)
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return nil, error
	}
	return &Server{
		accountClient,
		catalogClient,
		orderClient
	}, nil
}

func ()  {
	
}
