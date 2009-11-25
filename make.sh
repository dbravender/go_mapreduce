8g src/mapreduce.go
8g src/file_iter.go
8g -I . src/main.go
8g -I . src/wordcount.go
8l -o main mapreduce.8 main.8
8l -o wordcount wordcount.8 mapreduce.8
rm *.8
