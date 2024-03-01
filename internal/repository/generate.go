package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i UserRepository -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i LogRepository -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i ChatRepository -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i MessageRepository -o ./mocks/ -s "_minimock.go"
