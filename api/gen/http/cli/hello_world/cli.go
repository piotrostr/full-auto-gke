// Code generated by goa v3.8.4, DO NOT EDIT.
//
// hello-world HTTP client CLI support package
//
// Command:
// $ goa gen github.com/piotrostr/full-auto-gke/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	helloc "github.com/piotrostr/full-auto-gke/api/gen/http/hello/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `hello say-hello
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` hello say-hello --name "Ea aliquam qui sunt in maxime."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		helloFlags = flag.NewFlagSet("hello", flag.ContinueOnError)

		helloSayHelloFlags    = flag.NewFlagSet("say-hello", flag.ExitOnError)
		helloSayHelloNameFlag = helloSayHelloFlags.String("name", "REQUIRED", "The name to say hello to")
	)
	helloFlags.Usage = helloUsage
	helloSayHelloFlags.Usage = helloSayHelloUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "hello":
			svcf = helloFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "hello":
			switch epn {
			case "say-hello":
				epf = helloSayHelloFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "hello":
			c := helloc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "say-hello":
				endpoint = c.SayHello()
				data, err = helloc.BuildSayHelloPayload(*helloSayHelloNameFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// helloUsage displays the usage of the hello command and its subcommands.
func helloUsage() {
	fmt.Fprintf(os.Stderr, `The hello service, it says hello
Usage:
    %[1]s [globalflags] hello COMMAND [flags]

COMMAND:
    say-hello: SayHello implements say-hello.

Additional help:
    %[1]s hello COMMAND --help
`, os.Args[0])
}
func helloSayHelloUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] hello say-hello -name STRING

SayHello implements say-hello.
    -name STRING: The name to say hello to

Example:
    %[1]s hello say-hello --name "Ea aliquam qui sunt in maxime."
`, os.Args[0])
}