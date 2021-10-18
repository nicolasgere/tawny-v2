docker run  -v /Users/nicolas/go/src/tawny/consul/:/consul/config consul agent  -node=client-1 -join=172.17.0.2
docker run \                                                                                                                                
-d \
-p 8500:8500 \
-p 8600:8600/udp \
--name=badger \
consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
