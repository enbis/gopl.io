package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

const input1 = `
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
	want        string
}{
	{
		description: "input easy",
		input:       input1,
		want:        "[script]",
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

					fmt.Println("out ", buf.String())

					if err := tmpfile.Close(); err != nil {
						log.Fatal(err)
					}

					if strings.Trim(buf.String(), " ") != strings.Trim(tc.want, " ") {
						t.Fatalf("TestFindone %s \ngot: %s\nwant: %s", tc.input, out, tc.want)
					}
				})
			}
		})
	}
	/*
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

		fmt.Println("out ", buf.String())

		if err := tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
	*/
}
