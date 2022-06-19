/*
ideal:
1) every job run concurrency, we create one channel for it,
2) in functions run job, we return channel is a endpoint for other action get data stream
3) in point get data, we get data from this channel, order use chanel is order get data.
*/
