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

	t.Logf("RequestError: %s", requestError.Error())

	// expected 1000 unauthorized token
	if requestError.Content.Code != 1000 {
		t.Errorf("expected 1000 code, got %v", requestError.Content.Code)
		return
	}

	if requestError.Content.Message != "Invalid token" {
		t.Errorf("expected '%s' message, got %s", "Invalid token", requestError.Content.Message)
		return
	}

	t.Log("Test Pass.")
}
