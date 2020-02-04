# metric_intercept
A small project to work with metrics locally.

## Content
* [TCP Server](#tcp-server-graphite) - bind `localhost:2001` to simulate [graphite](https://graphite.readthedocs.io/en/latest/feeding-carbon.html) server
* [UDP Server](#udp-server-statsd) - bind `localhost:8125` to simulate [statsd](https://github.com/statsd/statsd) server
* [Ruby Graphite client](#ruby-graphite-client)

## Instalattion
### Golang
```
go get github.com/yorikya/metric_intercep
```

### Git 
```
git clone https://github.com/yorikya/metric_intercept
```

## Usage
### TCP Server Graphite
```
go run tcpserver.go

out=> Message app.my.metric.gauge 7 1580849242
```

### UDP Server Statsd
```
go run udpserver.go

out=> addr:10.10.1.1, get message: app.my.metric.counter:1|c

```

### Ruby graphite client
See example in `graphiteExporter.rb` file
