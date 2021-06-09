
require 'socket'

#Graphite Class
class Graphite
    def initialize(host, port, prefix)
      @host = host
      @port = port
      @prefix = prefix
    end
  
    def socket
      return @socket if @socket && !@socket.closed?
      @socket = TCPSocket.new(@host, @port)
    end
  
    def report(key, value, time = Time.now)
      begin
        socket.write("#{@prefix}.#{key} #{value.to_f} #{time.to_i}\n")
      rescue Errno::EPIPE, Errno::EHOSTUNREACH, Errno::ECONNREFUSED
        @socket = nil
        nil
      end
    end

    def close_socket
      @socket.close if @socket
      @socket = nil
    end
end


# Graphite connection
graphiteURL = "carbonrelay-20001-001-test"
graphitePort = 2009
$graphiteMetricPrefix = "yuri.tests"

graphite = Graphite.new(graphiteURL, graphitePort, $graphiteMetricPrefix)

$i = 0
$num = 100000
$m = "hello.world"
$mtag = "hello.world;tagname=test"

while $i < $num  do
  
  graphite.report($m, $i)
  puts "send metrics: #$graphiteMetricPrefix.#$m, value: #$i"
  $i +=1
  sleep 1
end

# graphite.report("hello.world", 10)
# puts "after sned metric!!!!!!!!"

