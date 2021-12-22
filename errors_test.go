package cas

import "testing"

// expected invalid token
// {
//	"error":{
//		"code":1000,
//		"message":"Invalid token"
//	}
// }
func TestRequestError_Error(t *testing.T) {
	client := New("abc", "xyz")

	_, err := client.Movie("12313213")

	requestError, ok := err.(*RequestError)

	if !ok {
		t.Errorf("%s is not RequestError", requestError)
		return
	}

	t.Logf("RequestError: %+v", requestError)
}
