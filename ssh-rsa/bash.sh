openssl genrsa -out rsa.priv 2048

openssl rsa -in rsa.priv -pubout -out rsa.pub

cat encrypted.bin | openssl pkeyutl -decrypt -inkey rsa.priv \
  -pkeyopt rsa_padding_mode:oaep \
  -pkeyopt rsa_oaep_md:sha256 \
  -pkeyopt rsa_mgf1_md:sha256
#rsa_padding_mode:oaep — OAEP padding.
#rsa_oaep_md:sha256 — хэш внутри OAEP.
#rsa_mgf1_md:sha256 — MGF1 тоже на SHA-256.

# шифруем текст через Go и сразу дешифруем через OpenSSL
echo "foo" | go run rsa_oaep.go | openssl pkeyutl -decrypt -inkey rsa.priv \
  -pkeyopt rsa_padding_mode:oaep \
  -pkeyopt rsa_oaep_md:sha256 \
  -pkeyopt rsa_mgf1_md:sha256
