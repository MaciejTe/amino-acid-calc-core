module github.com/MaciejTe/amino-acid-calc

go 1.13.1

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/go-resty/resty/v2 v2.1.0
	github.com/lib/pq v1.3.0
	github.com/mitchellh/panicwrap v0.0.0-20190228164358-f67bf3f3d291
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
