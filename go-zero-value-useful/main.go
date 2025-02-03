package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Config struct {
	output      io.Writer
	readTimeout time.Duration
}

// init runs before each method that wants a default value. This makes the zero
// value useful, having a valid default config. This is essentially the same as
// having a [NewConfig] function that returns a valid Config.
func (c *Config) init() {
	// The if could be something else like `if c.initialized == false`. Using an interface
	// field is useful as they have a `nil` default value but bear in mind if
	// you allow setting `nil` config.output as a possibility it could lead to bugs.
	if c.output == nil {
		c.output = os.Stdout
		c.readTimeout = 60 * time.Second
	}
}

// returns output
func (c *Config) Output() io.Writer {
	// if no output is
	c.init()
	return c.output
}

// SetOutput sets log ouput. If w is nil it does nothing.
func (c *Config) SetOutput(w io.Writer) {
	// If we allow `w` to be nil a second call to `c.ReadTimeout()` would
	// re-initialize the config, overwriting values.
	if w != nil {
		c.output = w
	}
}

func (c *Config) ReadTimeout() time.Duration {
	c.init()
	return c.readTimeout
}

func (c *Config) SetReadTimeout(d time.Duration) {
	c.readTimeout = d
}

func main() {
	// or c := Config{}
	var c Config
	fmt.Printf("zero value:\tc = %+v\n", c)

	fmt.Println("call to c.Output()")
	c.Output()
	fmt.Printf("value:\t\tc = %+v\n", c)

	fmt.Println("call to c.SetReadTimeout()")
	c.SetReadTimeout(20 * time.Second)

	fmt.Printf("value|\t\tc = %+v\n", c)
	fmt.Printf("c.ReadTimeout():\t%+v\n", c.ReadTimeout())
}
