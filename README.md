# Family Home Evening Schedule Generator

`fhe` is a go program for generating schedules for family home evening. 

Given a list of names of your family members who wish to participate in family home evening, a year and a month, `fhe` will generate a schedule like this,

```
===============================
Schedules
===============================
Date               2024-11-03
Spiritual Thought  Victor
Message            Peter
Game               Diana
-------------------------------
Date               2024-11-10
Spiritual Thought  Peter
Message            Diana
Game               Sariah
-------------------------------
Date               2024-11-17
Spiritual Thought  Diana
Message            Sariah
Game               Victor
-------------------------------
Date               2024-11-24
Spiritual Thought  Sariah
Message            Victor
Game               Peter
-------------------------------
```

## Installation

To build a `fhe` binary, you need to install the [Go](https://go.dev/doc/install) programming language. Then use make target `fhe`.
```
make fhe
```
This will create a `fhe` binary under `build/`. Move this to your `PATH` if you wish.

## Usage

```
fhe 2024 11 Diana,Victor,Peter,Sariah
```
