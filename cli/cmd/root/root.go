package root

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	filename = ".access_token"

	tokenHeader = "Bearer "

	lineUp = "\033[1A"

	addressFlag = "address"
	idFlag      = "id"
	usersFlag   = "users"

	addressFlagShort = "a"
	idFlagShort      = "n"
	usersFlagShort   = "u"
)

var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Chat client",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to chat service",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString(addressFlag)
		if err != nil {
			log.Fatalf("failed to get %s: %v", addressFlag, err)
		}

		err = login(context.Background(), addr)
		if err != nil {
			log.Printf("failed to login: %v", err)
		} else {
			fmt.Println(color.GreenString("\n\n[Successfully logged in]\n"))
		}
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from chat service",
	Run: func(_ *cobra.Command, _ []string) {
		err := logout()
		if err != nil {
			log.Printf("failed to logout: %v", err)
		} else {
			fmt.Println(color.GreenString("\n\n[Successfully logged out]\n"))
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create object",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete object",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to object",
}

var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Create new chat",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString(addressFlag)
		if err != nil {
			log.Fatalf("failed to get %s: %v", addressFlag, err)
		}

		users, err := cmd.Flags().GetString(usersFlag)
		if err != nil {
			log.Fatalf("failed to get users: %v", err)
		}

		id, err := createChat(context.Background(), addr, strings.Split(users, ","))
		if err != nil {
			log.Printf("failed to create chat: %v", err)
		} else {
			fmt.Printf("[%s %s]\n", color.CyanString("Created chat with id"), color.BlueString(strconv.Itoa(int(id))))
		}
	},
}

var deleteChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Delete chat",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString(addressFlag)
		if err != nil {
			log.Fatalf("failed to get %s: %v", addressFlag, err)
		}

		id, err := cmd.Flags().GetString(idFlag)
		if err != nil {
			log.Fatalf("failed to get %s: %v", idFlag, err)
		}

		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Fatalf("failed to convert id to int %v", err)
		}

		err = deleteChat(context.Background(), addr, intID)
		if err != nil {
			log.Printf("failed to delete chat: %v", err)
		} else {
			fmt.Printf("[%s %s]\n", color.CyanString("Deleted chat (if existed)"), color.YellowString(id))
		}
	},
}

var connectChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Connect to chat",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString(addressFlag)
		if err != nil {
			log.Fatalf("failed to get %s: %v", addressFlag, err)
		}

		id, err := cmd.Flags().GetString(idFlag)
		if err != nil {
			log.Fatalf("failed to get %s: %v", idFlag, err)
		}

		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Fatalf("failed to convert id to int %v", err)
		}

		err = connectChat(context.Background(), addr, intID)
		if err != nil {
			log.Printf("failed to connect to chat: %v", err)
		}
	},
}

// Execute ...
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(connectCmd)

	createCmd.AddCommand(createChatCmd)
	deleteCmd.AddCommand(deleteChatCmd)
	connectCmd.AddCommand(connectChatCmd)

	// Add flags to commands
	loginCmd.Flags().StringP(addressFlag, addressFlagShort, "", "`IP:port` of authentication service")
	createChatCmd.Flags().StringP(usersFlag, usersFlagShort, "", "chat participants")
	createChatCmd.Flags().StringP(addressFlag, addressFlagShort, "", "`IP:port` of chat service")
	deleteChatCmd.Flags().StringP(idFlag, idFlagShort, "", "ID of the chat to delete")
	deleteChatCmd.Flags().StringP(addressFlag, addressFlagShort, "", "`IP:port` of chat service")
	connectChatCmd.Flags().StringP(idFlag, idFlagShort, "", "ID of the chat to connect")
	connectChatCmd.Flags().StringP(addressFlag, addressFlagShort, "", "`IP:port` of chat service")

	// Mark required flags in commands
	err := loginCmd.MarkFlagRequired(addressFlag)
	if err != nil {
		log.Fatalf("failed to mark %s flag as required: %v", addressFlag, err)
	}

	err = createChatCmd.MarkFlagRequired(addressFlag)
	if err != nil {
		log.Fatalf("failed to mark %s flag as required: %v", idFlag, err)
	}

	err = deleteChatCmd.MarkFlagRequired(idFlag)
	if err != nil {
		log.Fatalf("failed to mark %s flag as required: %v", idFlag, err)
	}

	err = deleteChatCmd.MarkFlagRequired(addressFlag)
	if err != nil {
		log.Fatalf("failed to mark %s flag as required: %v", addressFlag, err)
	}

	err = connectChatCmd.MarkFlagRequired(idFlag)
	if err != nil {
		log.Fatalf("failed to mark %s flag as required: %v", idFlag, err)
	}

	err = connectChatCmd.MarkFlagRequired(addressFlag)
	if err != nil {
		log.Fatalf("failed to mark %s flag as required: %v", addressFlag, err)
	}
}
