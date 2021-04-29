# hiddenGoWebShell
 A web shell written used in KingOfTheHill matches

Once you have a shell on the machine, upload,rename and launch the binary.

A webserver will launch that will be launched, and all directories will return 75 bytes and a 200 response. To run execution, send a get request to the following.

```
http:\\<IPHERE>:8000\candlelight?masterplan=hostname
```

The response will be the output from that command. 

*Do not use in production environments. The webserver is running HTTP and information can be intercepted.*