package server

import (
	"context"
	"fmt"

	"github.com/anz-bank/sysl-template/gen/jsonplaceholder"
	"github.com/anz-bank/sysl-template/gen/simple"
)

// GetFoobarList refers to the endpoint in our sysl file
func GetFoobarList(ctx context.Context, req *simple.GetFoobarListRequest, client simple.GetFoobarListClient) (*jsonplaceholder.TodosResponse, error) {

	// Here we can make a request on the client object which was generated from the call to "myDownstream" in the sysl model
	// We will get the id equal to one, which was generated from out {id} from /todos/{id<:int}
	ans, err := client.GetTodos(ctx, &jsonplaceholder.GetTodosRequest{ID: 1})
	fmt.Println(ans)
	return ans, err
}
