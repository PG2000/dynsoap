# Dynsoap

## Purpose

dynsoap is a helper library to read all given records at dyn.com via SOAP with a single request, 
as the REST interface does not support this operation and need multiple requests.

When using external-dns with a lot of DNS records (dyn.com) it could take a long time to 
request all dyn.com records via the REST Api.
 
Unfortunately the maintainer of [dynectsoap](https://github.com/sanyu/dynectsoap) has not responded to requests regarding 
accept a [pull request](https://github.com/sanyu/dynectsoap/pull/2) nor for changes to the currently not available [license](https://github.com/sanyu/dynectsoap/issues/3)
which doesn't allow to fork the repo and do some custimzations.

## Generating code

The models and SOAP commands were generated with https://github.com/fiorix/wsdl2go
After that we need some tiny modifications so that the code can be used with dyn.com
and can be used from external-dns

The changes are:
```bash
Before: "Data AnyType" After: "Data interface{}"
```
and

```bash
Before: `xml:"msgs,omitempty" json:"msgs,omitempty"`
After : `xml:"messages,omitempty" json:"msgs,omitempty"`
```

This can be easily done with some make/sed commands:

```bash
make install-go-wsdl
make generate-code-from-wsdl
make replace-any-type
```
