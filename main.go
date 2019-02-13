import (
	"net"
	"io"
	"bytes"
	"strings"
	"os"
	"bufio"
)


func scanIP(ip string) {
	var buf bytes.Buffer
	
	conn, err := net.Dial("tcp",strings.Join([]string{ip,"80"}, ":")

	def conn.Close()
	
	if err != nil {
		log.Print(err)
		return
	}

	io.Copy(&buf, conn)

	fmt.Println("{} @ port 80 open, returns {}", ip, buf)
}

func gen_addresses() []string {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Print(err)
	}

	defer f.close()
	addresses = make([]string,0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		addresses = append(addresses, scanner.Text())
	}
	
	return addresses
}


func main() {
	for ip := range gen_addresses() {
		go scanIP(ip)
	}
}