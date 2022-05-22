package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
    "encoding/json"
    "runtime"
    "regexp"
)

var publicKey []byte = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0Gkcvz0zpYIjcXS+alKf
WeAJtsctBN/E3n8BXpg3f0yX3hBgYFmZz5b9T6EZa791uqQpvIQVn39EZSmNt1mR
fnNvggEchkmj68xibo3EaYO27rBdHWKJk2ECaX2KKKBQ6TUMpudBPtRpsPblWEOu
I8Wo6q3WH5YrxHmy9C+BiAOOkuyfhpy5SY2BEmRyeE2QP2m/O+W2ACjBpPuz+wDT
YVmbuEvKXZ/SHwvUYUpoew2r5r6flRO7W+tJCQO7EecMwpEzNUCUGEf+UsUvC+z9
/QCXrH1bpAhvmwJmSmkwNqTOy35SvDgw8gkfWPhJcha0xskVNbWg61cfOrxg1l86
eQIDAQAB
-----END PUBLIC KEY-----`)

func ValidateLicense(signed []byte) (bool, string) {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(signed)))

	_, err := base64.StdEncoding.Decode(decoded, signed)
	if err != nil {
		return false, ""
	}

	if len(decoded) <= 256 {
		return false, ""
	}

	// remove null terminator
	for decoded[len(decoded)-1] == byte(0) {
		decoded = decoded[:len(decoded)-1]
	}

	plaintext := decoded[:len(decoded)-256]
	signature := decoded[len(decoded)-256:]

	block, _ := pem.Decode(publicKey)

	public, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, ""
	}

	rsaPublic := public.(*rsa.PublicKey)

	h := sha512.New()
	h.Write(plaintext)
	d := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(rsaPublic, crypto.SHA512, d, signature)
	if err != nil {
		return false, ""
	}

	return true, string(plaintext)
}

func GetFile(name string) ([]byte) {
    file, _ := os.Open(name)
    defer file.Close()
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    return fileBytes
}

func genFlag() (string) {
    /** Omitted **/
}

func main() {
	if success, plaintext := ValidateLicense(GetFile("license.out")); success {
        fmt.Println("Signature: Good")

        var f interface{}
        json.Unmarshal([]byte(plaintext), &f)
        m := f.(map[string]interface{})
        can_view := m["can_view_flag"].(bool)

        if (can_view) {
            fmt.Println("Flag Viewing Licensed: true")
            fmt.Println("Flag: " + genFlag())
        } else {
            fmt.Println("Flag Viewing Licensed: false")
        }

    } else {
        fmt.Println("Signature: Bad")
    }
}