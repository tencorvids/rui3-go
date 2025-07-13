# rui3-go

Go implementation of the RUI3 library AT commands, currently incomplete but very easy to extend (open a pr!).
Tested against the RAK3172 module via an ESP32 but should support all RUI3 modules.

## Usage

```go
func main() {
	portName := "/dev/ttyS0"
	rui, err := rui3.New(portName)
	if err != nil {
		slog.Error("Failed to create RUI3 instance", "error", err)
		os.Exit(1)
	}
	defer rui.Close()

	attention, err := rui.Attention()
	if err != nil {
		slog.Error("Failed to get attention", "error", err)
		os.Exit(1)
	}
	slog.Info("Attention", "attention", attention)
```

More examples found in `/cmd`.

## Resources

https://docs.rakwireless.com/product-categories/software-apis-and-libraries/rui3/at-command-manual/
https://store.rakwireless.com/products/wisduo-lpwan-module-rak3172?utm_source=rak3172landingpage&utm_medium=header&utm_campaign=RAKwireless&variant=44068554473670
https://github.com/beegee-tokyo/RUI3-Arduino-Library
