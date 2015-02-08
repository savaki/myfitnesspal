# myfitnesspal

[![GoDoc](https://godoc.org/github.com/savaki/myfitnesspal?status.svg)](https://godoc.org/github.com/savaki/myfitnesspal)

[![Build Status](https://snap-ci.com/savaki/myfitnesspal/branch/master/build_image)](https://snap-ci.com/savaki/myfitnesspal/branch/master)

A client library and command line tool for myfitnesspal.

## Installation

```
go get github.com/savaki/myfitnesspal/myfitnesspal
```

## Command Line Usage

In addition to being a go library, myfitnesspal can be used from the command line to retrieve data from myfitnesspal.com in json format.

### MyFitnessPal Authentication

All myfitnesspal command line requests need to be authenticated.  You can either authenticate by tacking on ```--username``` and ``--password`` to each request as follows:

```
myfitnesspal food-diary --username YOUR-USERNAME --password YOUR_PASSWORD
--
```

Or by setting environment variables:

```
export MYFITNESSPAL_USERNAME=YOUR-USERNAME
export MYFITNESSPAL_PASSWORD=YOUR-PASSWORD
```

## Food Diary

### Retrieve my macro intake for today

Assuming you've set your username and password in the environment (see MyFitnessPal Authentication above).

```
myfitnesspal food-diary 
```

Will return something like: 

```
{
  "breakfast": [
    {
      "label": "Breakfast Burrito",
      "calories": 450,
      "carbs": 250,
      "fat": 150,
      "protein": 50,
      "sodium": 600,
      "sugar": 0
    }
  ],
  "totals": {
    "label": "Totals",    
    "calories": 450,
    "carbs": 250,
    "fat": 150,
    "protein": 50,
    "sodium": 600,
    "sugar": 0
  },
  "goal": {
    "label": "Your Daily Goal",    
    "calories": 450,
    "carbs": 250,
    "fat": 150,
    "protein": 50,
    "sodium": 600,
    "sugar": 0
  },
  "remaining": {
    "label": "Remaining",    
    "calories": 0,
    "carbs": 0,
    "fat": 0,
    "protein": 0,
    "sodium": 0,
    "sugar": 0
  }
}
```

### Retrieve my macro intake for a specific date

To retrieve my intake on for a specific date, like ```Feb 1st, 2015```, I can type:

```
myfitnesspal food-diary --date 2015-02-01
```

## Future Plans

* Looking for a feature, but don't see it?  Ping me and I'll see about adding it.

