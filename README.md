# timeblock plugin for coreDns
A plugin to selectively block dns lookup for ip and time combination. Inspired from acl plugin

## Description
use this to selectively block IP between a time range

##Syntax/Example
Example below blocks roblox.com every day from 8am to 5pm for a given IP. Outside the range normal forward is applied.
```
roblox.com:53 {
  timeblock 0:7-08:00-17:00 192.168.86.74
  forward . 8.8.8.8:53 
}
```

- **date range** the time range as day of week, startTime, endTime:

    - **day** First part is day of the week, eg here 0 to 7 - 0 being Sunday

    - **startTime** Second part is start time range as defined HH:MM, here 08:00

    - **endTime** third part is end time range as defined HH:MM, here 17:00

- **IP** Typical CIDR notation and single IP address are supported - In the example above single IP:192.168.86.74 

##install
checkout coreDns at https://github.com/coredns/coredns

edit plugin.cfg and add 
```
...
timeblock:https://github.com/pmonestie/corednsTimeBlock.git
...

```

Make sure to add that plugin near the top, probably next to acl plugin    
run
```
go generate
go run coredns.go
```

