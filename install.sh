mkdir -p ~/pim/build/ && rm -rf ~/pim/build/* && cp -rf ../pim ~/pim/build
go build -o pim main.go
mv pim /usr/local/bin/