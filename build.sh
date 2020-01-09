GOOS="linux" GOARCH="amd64" go build -o "./bin/ccmu" 
GOOS="windows" GOARCH="amd64" go build -o "./bin/ccmu.exe"
GOOS="darwin" GOARCH="amd64" go build -o "./bin/ccmu_mac"

cd bin


sha256sum ccmu.exe ccmu ccmu_mac > checksum_sha256.txt

cd ..