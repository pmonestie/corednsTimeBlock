# timeblock plugin for coreDns
A plugin to selectively block dns lookup for ip and time combination. Inspired from acl plugin

## Description
use this to selectively block IP between a time range

## Syntax
Example below blocks roblox.com every day from 8am to 5pm for a given IP. Outside the range normal forward is applied.
```
roblox.com:53 {
  timeblock 0:7-08:00-17:00 192.168.86.74
  forward . 8.8.8.8:53 
}
```

- **date range** the time range as day of week, startTime, endTime. The time is local time, there is not timeZone translation so make sure your server/configuration coincide

    - **day** First part is day of the week, eg here 0 to 7 - 0 being Sunday

    - **startTime** Second part is start time range as defined HH:MM, here 08:00

    - **endTime** third part is end time range as defined HH:MM, here 17:00

- **IP** Typical CIDR notation and single IP address are supported - In the example above single IP:192.168.86.74 

## install
checkout coreDns at https://github.com/coredns/coredns

edit plugin.cfg and add 
```
...
timeblock:github.com/pmonestie/corednsTimeBlock/timeblock

...

```
Make sure to add that plugin near the top, next to acl plugin is a good place: remember that plugins are executed in order, so it definately needs to be before forward plugin
run
```
go generate
go run coredns.go
```

## example
block roblox from 9am to 5pm
This is combined with cache...
```
roblox.com:53 {
    cache 100
    timeblock 0:7-09:00-17:00 192.168.86.75
    forward . 8.8.8.8:53
}
.:53 {
    forward . 8.8.8.8:53
}
```
