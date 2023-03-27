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

	url := "https://v1.genr.ai/api/circuit-element/generate-image"
	payload := strings.NewReader("{\n  \"prompt\": \"detailed realistic image of elements of  penguins and rainbows and water slides\",\n  \"height\": 512,\n  \"width\": 512,\n  \"model\": \"stable-diffusion-2\",\n  \"n_images\": 1\n}")
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
