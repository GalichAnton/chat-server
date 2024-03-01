package services

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i UserService -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i ChatService -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i MessageService -o ./mocks/ -s "_minimock.go"
