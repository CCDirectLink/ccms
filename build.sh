GOOS="linux" GOARCH="amd64" go build -o "./bin/ccms" 
GOOS="windows" GOARCH="amd64" go build -o "./bin/ccms.exe"
GOOS="darwin" GOARCH="amd64" go build -o "./bin/ccms_mac"

cd bin



export PATH=$PATH:"/C/Program Files/7-zip/"

sha256sum ccms_mac > checksum_sha256_mac.txt
7z.exe a -ttar -so -an ccms_mac checksum_sha256_mac.txt | 7z.exe a -si ccms_mac.tgz


sha256sum ccms  > checksum_sha256_linux.txt
7z.exe a -ttar -so -an ccms checksum_sha256_linux.txt | 7z.exe a -si ccms_linux.tgz




cd ..
