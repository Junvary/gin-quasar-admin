package boot

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Boot(router *gin.Engine, projectName string, port string, logger *slog.Logger) {
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("【gin-quasar-admin】 "+projectName+" listen: %s\n", err)
		}
	}()
	logo(projectName, port, logger)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("【gin-quasar-admin】 " + projectName + " start to exit...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("【gin-quasar-admin】 "+projectName+" force exit：", err)
	}

	log.Println("【gin-quasar-admin】 " + projectName + " exit complete！")
}

var logoText = "  ________.__                 ________                                                 _____       .___      .__        \n /  _____/|__| ____           \\_____  \\  __ _______    ___________ _______            /  _  \\    __| _/_____ |__| ____  \n/   \\  ___|  |/    \\   ______  /  / \\  \\|  |  \\__  \\  /  ___/\\__  \\\\_  __ \\  ______  /  /_\\  \\  / __ |/     \\|  |/    \\ \n\\    \\_\\  \\  |   |  \\ /_____/ /   \\_/.  \\  |  // __ \\_\\___ \\  / __ \\|  | \\/ /_____/ /    |    \\/ /_/ |  Y Y  \\  |   |  \\\n \\______  /__|___|  /         \\_____\\ \\_/____/(____  /____  >(____  /__|            \\____|__  /\\____ |__|_|  /__|___|  /\n        \\/        \\/                 \\__>          \\/     \\/      \\/                        \\/      \\/     \\/        \\/"

func logo(projectName, port string, logger *slog.Logger) {
	fmt.Println(logoText)
	fmt.Println("Welcome to Gin-Quasar-Admin !")
	fmt.Println("Github: https://github.com/Junvary/gin-quasar-admin ")
	fmt.Println("Expecting Your Star!")
	fmt.Println("【gin-quasar-admin】 " + projectName + " started, listening " + port + "...")
	logger.Info("【gin-quasar-admin】 " + projectName + " started, listening " + port + "...")
}
