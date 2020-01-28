
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
graphiteURL = "localhost"
graphitePort = 2001

graphite = Graphite.new(graphiteURL, graphitePort, "yuri.tests")
graphite.report("hello.world", 10)
puts "after sned metric!!!!!!!!"



