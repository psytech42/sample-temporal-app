package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Withdraw(ctx context.Context, transferDetails TransferDetails) error {
	fmt.Printf(
		"\nWithdrawing $%f from account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ReferenceID,
	)

	url := "https://v1.genr.ai/api/circuit-element/translate-code"
	payload := strings.NewReader("{\n  \"code\": \"using namespace std;int main(){vector <int> arr1 = {1, 2, 3, 4};cout<< arr1.size() << endl;return 0;}\",\n  \"temperature\": 0,\n  \"input_language\": \"C++\",\n  \"desired_language\": \"Haskell\"\n}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err == nil {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(res)
		fmt.Println(string(body))
	}
	return nil
}

func Deposit(ctx context.Context, transferDetails TransferDetails) error {
	fmt.Printf(
		"\nDepositing $%f into account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	// Switch out comments on the return statements to simulate an error
	//return fmt.Errorf("deposit did not occur due to an issue")
	return nil
}
