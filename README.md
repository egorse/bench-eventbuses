Benchmark golang in-process eventbuses
======================================

Stupid tests for next in process event buses:
 - https://github.com/asaskevich/EventBus
 - https://github.com/olebedev/emitter
 - https://github.com/cskr/pubsub

How to run
==========

```
git clone https://github.com/egorse/bench-eventbuses && 
cd bench-eventbuses &&
go test -bench=.
```