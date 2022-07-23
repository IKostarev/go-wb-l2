package ntpmodule

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func NtpModule() {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(404)
	}

	fmt.Printf("Текущее время: %v \n", time.Now())
	fmt.Println("Ntp время:", t.String())
}
