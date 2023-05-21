# Introduction:
---
This repo is a simple re-implementation of a take-home assignment that implements two containerized services: a key value store, and a client to test the implementation of said KV store.

This implementation builds upon my initial submission in a couple of ways:
- First I wanted to take a stab at using Golang in my implementation as opposed to Python which I used in my first submission. While python is my go-to language I wanted to challenge myself while not under the pressure of submitting this as part of an interview.
- Secondly I wanted to try to recall and implement all the improvements that we talked about in the assessment review, namely using a service like Redis to handle the implementation of the Key Value store.

# Redis as a Key Value Store:
---
This is my first time interacting with Redis, and my second time building containerized applications from scratch on my own, the first being my initial submission for the assessment. Redis seems like a nice option for implementing a quick and easy KV store. I containerized an instance of it and played with some basic functionality using the `redis-cli` tool and with that information coupled with the SDK documentation I don't foresee having issues with implementation.

# Developing a client in Go:
---
As exists with many applications out in the world Redis comes with a nifty SDK that the client can use to interact with Redis in a programmatic way. As previously mentioned Python is my go-to language for development so the implementation in Go will likely be simpler than someone with a more advanced understanding of the language but I think this will serve as a fun educational challenge.


