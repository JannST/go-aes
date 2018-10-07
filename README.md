# go-aes
This is my implementation of the Rijndael Algorithm (AES) in Go. It can encrypt and decrypt strings (yeah).
It was implemented using this spec: https://nvlpubs.nist.gov/nistpubs/fips/nist.fips.197.pdf

It uses github.com/urfave/cli as command line application framework.

## Examples: 
./go-aes encrypt 1234567887654321 "hello world"
Ciphertext (HEX): f33e4af222e89f54df51315093bd1168

Do not forget to set the -hex flag when decrypting hex values
- './go-aes -hex decrypt 1234567887654321 f33e4af222e89f54df51315093bd1168'
Plaintext: hello world


It is also possible to set the key length, but it must be devidable by 32:
- ./go-aes -klength 192 encrypt 111111112222222233333333 "hello world"
- ./go-aes -klength 32 encrypt 1234 "hello world"
- ./go-aes -klength 64 encrypt 12341234 "hello world"

* You can also set the number of encryption rounds 
- ./go-aes -klength 64 -rounds 17 encrypt 12345678 McMurphey
Ciphertext (HEX): be3392467b39e2f825e6b8fb6403304a

* Make sure to set the same number of rounds when decrypting

./go-aes -klength 64 -rounds 17 -hex decrypt 12345678 be3392467b39e2f825e6b8fb6403304a
Plaintext: McMurphey





