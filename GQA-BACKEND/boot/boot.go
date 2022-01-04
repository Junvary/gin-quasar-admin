package boot

import (
	"context"
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Boot()  {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", global.GqaConfig.System.Port),
		Handler: Router(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("【gin-quasar-admin】监听：%s\n", err)
		}
	}()
	logo(fmt.Sprintf(":%d", global.GqaConfig.System.Port))
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("【gin-quasar-admin】开始退出...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("【gin-quasar-admin】强制退出：", err)
	}

	log.Println("【gin-quasar-admin】退出完成！")
}

func logo(port string)  {
	fmt.Print("\n______________                 _______                                                 _______ _________            _____         \n__  ____/___(_)_______         __  __ \\____  ________ _______________ _________        ___    |______  /_______ ___ ___(_)_______ \n_  / __  __  / __  __ \\_________  / / /_  / / /_  __ `/__  ___/_  __ `/__  ___/__________  /| |_  __  / __  __ `__ \\__  / __  __ \\\n/ /_/ /  _  /  _  / / /_/_____// /_/ / / /_/ / / /_/ / _(__  ) / /_/ / _  /    _/_____/_  ___ |/ /_/ /  _  / / / / /_  /  _  / / /\n\\____/   /_/   /_/ /_/         \\___\\_\\ \\__,_/  \\__,_/  /____/  \\__,_/  /_/             /_/  |_|\\__,_/   /_/ /_/ /_/ /_/   /_/ /_/ \n                                                                                                                                  \n")
	fmt.Println("欢迎使用Gin-Quasar-Admin!")
	fmt.Println("项目地址: https://github.com/Junvary/gin-quasar-admin ")
	fmt.Println("欢迎交流，感谢Star!")
	fmt.Println("系统已启动,开始监听 " + port + " 端口...")
	global.GqaLog.Info("欢迎使用Gin-Quasar-Admin,系统已启动,开始监听 " + port + " 端口...")
}
