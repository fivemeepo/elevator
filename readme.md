# Elevator SDK
---

# Get Started
- Build: go build && ./elevator
- Usage: [up|down|goto|history] [floor_number]

We provide 3 buttons:
- up
- down
- goto

Example:
```
$ up 10
$ down 2
$ goto 5
$ history
```

# Design
This projects implements a elevator schedule system. Here's the project directory:
```
internal/ # elevator's low-level api
queue/ # a queue data structure, used for elevator's taks queue
sdk/ # including interfaces and core schedule algorithm
``` 

I use two task queues for the elevator, one upward queue and one downward queue. Every time you press the button, it will add a task to one of the queues, which depends of the elevator's current status. 
