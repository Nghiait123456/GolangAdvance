/*
ideal:
    many process run with many routine and many chanel stream data
    we need timeout for stop read data
    for all job runs concurrency in routine, time is same asynchronous and concurrency.
    in golang, perfect support it with "select case"
*/
