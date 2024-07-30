## NodeJs and Go Communication app in a stream-like way

This app is get data from a txt file contain some numbers separate by space, read it in stream way by NodeJs and pass it to
`normalizer` app which was written by Go lang.
The Go program read the data from `stdin` and do some process on it then saved it on `dest.txt` file.

The input file example > `2134214 124 12341324324 32432432432432 4324324324`

The output file example > `$2.134.214 $124 $12.341.324.324 $32.432.432.432.432 $4.324.324.324`

Because this program reads and writes data in a stream, it performs the processes in the most optimal way.
In the photo below, you can see the amount of CPU and memory usage in the operations performed on a `10GB` file including the input sample numbers.
![Usage of app](https://raw.githubusercontent.com/dinno7/node_go_stream/master/usage.png)

#### Run app
```bash
node ./app.js ./src.txt
```
