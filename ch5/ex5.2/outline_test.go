package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

const input0 = " "
const input1 = `
<script>
(function() {
  var ga = document.createElement("script"); ga.type = "text/javascript"; ga.async = true;
  ga.src = ("https:" == document.location.protocol ? "https://ssl" : "http://www") + ".google-analytics.com/ga.js";
  var s = document.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ga, s);
})();
</script>`

const input2 = `
<footer>
  <div class="Footer ">
    <img class="Footer-gopher" src="/lib/godoc/images/footer-gopher.jpg" alt="The Go Gopher">
    <ul class="Footer-links">
      <li class="Footer-link"><a href="/doc/copyright.html">Copyright</a></li>
      <li class="Footer-link"><a href="/doc/tos.html">Terms of Service</a></li>
      <li class="Footer-link"><a href="http://www.google.com/intl/en/policies/privacy/">Privacy Policy</a></li>
      <li class="Footer-link"><a href="http://golang.org/issues/new?title=x/website:" target="_blank" rel="noopener">Report a website issue</a></li>
    </ul>
    <a class="Footer-supportedBy" href="https://google.com">Supported by Google</a>
  </div>
</footer>

<script>
(function() {
  var ga = document.createElement("script"); ga.type = "text/javascript"; ga.async = true;
  ga.src = ("https:" == document.location.protocol ? "https://ssl" : "http://www") + ".google-analytics.com/ga.js";
  var s = document.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ga, s);
})();
</script>`

var testCases = []struct {
	description string
	input       string
	want        int
	mapv        string
}{
	{
		description: "input zero",
		input:       input0,
		want:        0,
		mapv:        "script",
	},
	{
		description: "input easy",
		input:       input1,
		want:        1,
		mapv:        "script",
	},
	{
		description: "input less easy",
		input:       input2,
		want:        2,
		mapv:        "script",
	},
}

var implementations = []struct {
	descr string
	f     func()
}{
	{
		descr: "Findone",
		f:     main,
	},
}

func TestMain(t *testing.T) {
	for _, impl := range implementations {
		t.Run(impl.descr, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.description, func(t *testing.T) {

					tmpfile, err := ioutil.TempFile("", "example")
					if err != nil {
						log.Fatal(err)
					}
					defer os.Remove(tmpfile.Name())

					source := []byte(input1)
					if _, err := tmpfile.Write(source); err != nil {
						log.Fatal(err)
					}

					if _, err := tmpfile.Seek(0, 0); err != nil {
						log.Fatal(err)
					}

					oldStdin := os.Stdin
					defer func() {
						os.Stdin = oldStdin
					}()

					buf := &bytes.Buffer{}
					out = buf

					os.Stdin = tmpfile

					main()

					strout := buf.String()
					fmt.Println("out ", strout)

					if err := tmpfile.Close(); err != nil {
						log.Fatal(err)
					}

					ss := strings.Split(strout, "&&")
					m := make(map[string]int)
					for _, pair := range ss {
						if strings.Contains(pair, "=") {
							z := strings.Split(pair, "=")
							intval, _ := strconv.Atoi(z[1])
							m[z[0]] = intval
						}
					}

					val := m[tc.mapv]

					fmt.Println("out ", val)

					if val != tc.want {
						t.Fatalf("TestFindone %s \ngot: %d want: %d", tc.input, val, tc.want)
					}
				})
			}
		})
	}
}
