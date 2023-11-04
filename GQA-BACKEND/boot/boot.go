package boot

import (
	"context"
	"errors"
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Boot(router *gin.Engine, projectName string, port string) {
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf(projectName+" listen: %s\n", err)
		}
	}()
	logo(projectName, port)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println(projectName + " start to exit...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(projectName+" force exit：", err)
	}

	log.Println(projectName + " exit complete！")
}

var logoText = "  ________.__                 ________                                                 _____       .___      .__        \n /  _____/|__| ____           \\_____  \\  __ _______    ___________ _______            /  _  \\    __| _/_____ |__| ____  \n/   \\  ___|  |/    \\   ______  /  / \\  \\|  |  \\__  \\  /  ___/\\__  \\\\_  __ \\  ______  /  /_\\  \\  / __ |/     \\|  |/    \\ \n\\    \\_\\  \\  |   |  \\ /_____/ /   \\_/.  \\  |  // __ \\_\\___ \\  / __ \\|  | \\/ /_____/ /    |    \\/ /_/ |  Y Y  \\  |   |  \\\n \\______  /__|___|  /         \\_____\\ \\_/____/(____  /____  >(____  /__|            \\____|__  /\\____ |__|_|  /__|___|  /\n        \\/        \\/                 \\__>          \\/     \\/      \\/                        \\/      \\/     \\/        \\/"

func logo(projectName, port string) {
	fmt.Println(logoText)
	fmt.Println("Welcome to Gin-Quasar-Admin !")
	fmt.Println("Github: https://github.com/Junvary/gin-quasar-admin ")
	fmt.Println("Expecting Your Star!")
	fmt.Println(projectName + " started, listening " + port + "...")
	global.GqaSLogger.Info(projectName + " started, listening " + port + "...")
}
