GENERATED_FILE_PATH := pkg/dynsoap/services.go

clean:
	rm -f $(GENERATED_FILE_PATH)

install-go-wsdl:
	go get github.com/fiorix/wsdl2go

generate-code-from-wsdl:
	gowsdl -o services.go -p dynsoap https://api2.dynect.net/wsdl/current/Dynect.wsdl
	mv dynsoap ./pkg

replace-any-type:
	sed -i 's/Data AnyType/Data interface{}/g' $(GENERATED_FILE_PATH)
	sed -i 's/xml:\"msgs/xml\:\"messages/g' $(GENERATED_FILE_PATH)

all: clean generate-code-from-wsdl replace-any-type