8g src/mapreduce.go
8g src/file_iter.go
8g -I . src/main.go
8g -I . src/wordcount.go
8l -o main -L . main.8
8l -o wordcount -L . wordcount.8
rm *.8
