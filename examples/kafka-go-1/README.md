https://developer.confluent.io/get-started/go/

cd .\examples\kafka-go-1\


go mod init kafka-go-1
go get github.com/confluentinc/confluent-kafka-go/kafka

go build -o out/producer util.go producer.go

go build -o out/consumer util.go consumer.go

./out/producer getting-started.properties
./out/consumer getting-started.properties 
./out/consumer.exe getting-started.properties 
## Troubleshooting


cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%

gcc (the GNU Compiler Collection) provides a C compiler. On Windows, install TDM-GCC. The github.com/miekg/pkcs11 package uses cgo. Cgo enables the creation of Go packages that call C code.

https://jmeubank.github.io/tdm-gcc/

$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")

echo $Env:PATH