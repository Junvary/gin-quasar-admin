package boot

import (
	"context"
	"errors"
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Boot() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.GqaConfig.System.Port),
		Handler: Router(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("【gin-quasar-admin】listen: %s\n", err)
		}
	}()
	logo(fmt.Sprintf(":%d", global.GqaConfig.System.Port))
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("【gin-quasar-admin】start to exit...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("【gin-quasar-admin】force exit：", err)
	}

	log.Println("【gin-quasar-admin】 exit complete！")
}

var logoText = "  ________.__                 ________                                                 _____       .___      .__        \n /  _____/|__| ____           \\_____  \\  __ _______    ___________ _______            /  _  \\    __| _/_____ |__| ____  \n/   \\  ___|  |/    \\   ______  /  / \\  \\|  |  \\__  \\  /  ___/\\__  \\\\_  __ \\  ______  /  /_\\  \\  / __ |/     \\|  |/    \\ \n\\    \\_\\  \\  |   |  \\ /_____/ /   \\_/.  \\  |  // __ \\_\\___ \\  / __ \\|  | \\/ /_____/ /    |    \\/ /_/ |  Y Y  \\  |   |  \\\n \\______  /__|___|  /         \\_____\\ \\_/____/(____  /____  >(____  /__|            \\____|__  /\\____ |__|_|  /__|___|  /\n        \\/        \\/                 \\__>          \\/     \\/      \\/                        \\/      \\/     \\/        \\/"

func logo(port string) {
	fmt.Println(logoText)
	fmt.Println("Welcome to Gin-Quasar-Admin !")
	fmt.Println("Github: https://github.com/Junvary/gin-quasar-admin ")
	fmt.Println("Expecting Your Star!")
	fmt.Println("System started, listening " + port + "...")
	global.GqaSLogger.Info("System started, listening " + port + "...")
}
