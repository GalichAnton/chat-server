package root

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"

	descAuth "github.com/GalichAnton/auth/pkg/auth_v1"
	"github.com/GalichAnton/platform_common/pkg/closer"
)

func login(ctx context.Context, address string) error {
	client, err := authClient(address)
	if err != nil {
		return err
	}

	var username string

	fmt.Print(color.CyanString("Username: "))
	_, err = fmt.Scan(&username)
	if err != nil {
		return err
	}

	fmt.Print(color.CyanString("Password (no echo): "))
	password, err := term.ReadPassword(0)
	if err != nil {
		return err
	}

	res, err := client.Login(
		ctx, &descAuth.LoginRequest{
			Email:    username,
			Password: string(password),
		},
	)
	if err != nil {
		return err
	}

	resAccessToken, err := client.GetAccessToken(
		ctx, &descAuth.GetAccessTokenRequest{
			RefreshToken: res.GetRefreshToken(),
		},
	)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	closer.Add(file.Close)

	wr := bufio.NewWriter(file)
	_, err = wr.WriteString(resAccessToken.GetAccessToken())
	if err != nil {
		return err
	}

	err = wr.Flush()
	if err != nil {
		return err
	}
	return nil
}
