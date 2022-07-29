package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
)

var myTelnet tcpCustom

type tcpCustom struct {
	timeout     time.Duration
	host        string
	port        string
	connect     net.Conn
	err         error
	wg          *sync.WaitGroup
	stopChannel []chan struct{}
}

type takeArgs struct {
	timeout time.Duration
	host    string
	port    string
}

func findArguments(args []string, str string) string {
	reg := regexp.MustCompile(str)

	for _, j := range args {
		if reg.MatchString(j) {
			return j
		}
	}

	return ""
}

func (t *tcpCustom) Listen() {
	t.wg.Add(1)
	fmt.Println("Слушаем хост...")

	for {
		mess, err := bufio.NewReader(t.connect).ReadString('\n')

		if err != nil {
			fmt.Println("Ошибка со стороны хоста - ", err)
		}

		fmt.Printf("Ответ от хост %s: %s", t.host, mess)
	}

	t.wg.Done()
}

func (t *tcpCustom) Write() {
	t.wg.Add(1)
	fmt.Println("Пишем в хост...")

	for {
		r := bufio.NewReader(os.Stdin)
		text, _ := r.ReadString('\n')

		_, err := fmt.Fprintf(t.connect, text+"\n")

		if err != nil {
			t.Stop(1)
			fmt.Println("Есть ошибка, завершаю работу - ", err)
			break
		}
	}
	t.wg.Done()
}

func (t *tcpCustom) Stop(num int) {
	t.connect.Close()
	t.wg.Add(-1)
	os.Exit(num)
}

func (t *tcpCustom) TimeOut() error {
	c, err := net.DialTimeout("tcp", t.host+":"+t.port, t.timeout)

	if err != nil {
		fmt.Println("Ошибка при инициалзиации Dial Timeout")
		return err
	}

	t.connect = c
	return nil
}

func calcTimeOut(arg string) time.Duration {
	if !strings.Contains(arg, "=") {
		res, _ := time.ParseDuration("10s")
		return res
	}

	str := strings.Split(arg, "=")[1]
	res, err := time.ParseDuration(str)

	if err != nil {
		log.Fatalf("Ошибка при указании таймаута - %s, введено - %s", err, str)
	}

	return res
}

func MakeArgs() takeArgs {
	var arg takeArgs

	args := os.Args[1:]
	arg.timeout = calcTimeOut(findArguments(args, `--timeout=`))
	arg.host = args[len(args)-2]
	arg.port = args[len(args)-1]

	return arg
}

func prepTcpDialer(arg takeArgs) *tcpCustom {
	tcpCustoms := &tcpCustom{
		timeout: arg.timeout,
		host:    arg.host,
		port:    arg.port,
		wg:      new(sync.WaitGroup),
	}

	return tcpCustoms
}

func main() {
	par := MakeArgs()
	myTelnet = *prepTcpDialer(par)
	err := myTelnet.TimeOut()
	sign := make(chan os.Signal, 1)

	if err != nil {
		log.Fatalf("Ошибка при создании таймаута - %s", err)
	}

	go myTelnet.Listen()
	go myTelnet.Write()

	signal.Notify(sign,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGHUP,
	)

	go func() {
		<-sign
		fmt.Println("Программа завершена вручную")
		myTelnet.Stop(0)
	}()
	myTelnet.wg.Wait()
}
