GOOS="linux" GOARCH="amd64" go build -o "./bin/ccms" 
GOOS="windows" GOARCH="amd64" go build -o "./bin/ccms.exe"
GOOS="darwin" GOARCH="amd64" go build -o "./bin/ccms_mac"

cd bin



export PATH=$PATH:"/C/Program Files/7-zip/"

mkdir checksums

sha256sum ccms_mac > ./checksums/checksum_sha256_mac.txt
7z.exe a -ttar -so -an ccms_mac ./checksums/checksum_sha256_mac.txt ../LICENSE | 7z.exe a -si ./compressed/ccms_mac.tgz


sha256sum ccms > ./checksums/checksum_sha256_linux.txt
7z.exe a -ttar -so -an ccms ./checksums/checksum_sha256_linux.txt ../LICENSE | 7z.exe a -si ./compressed/ccms_linux.tgz

sha256sum ccms.exe > ./checksums/checksum_sha256_windows.txt

7z.exe a ./compressed/ccms_windows.zip ccms.exe ./checksums/checksum_sha256_windows.txt ../LICENSE



cd ..
