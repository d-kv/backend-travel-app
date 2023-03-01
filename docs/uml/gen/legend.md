```protobuf
message PointOfInterest{
	required string uuid = 1;
	Address address = 2;
	/* see bellow */

	string name = 3;
	Category category = 4;
	/* see bellow */

	google.type.Latlng lat_lng = 5;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/latlng.proto */

	oneof check{
		AverageCheck average_check = 6;
		/* see bellow */

		CheckInterval check_interval = 7;
		/* see bellow */
	}
	repeated WorkingHours working_hours = 8;
	/* see bellow */

	Rating rating = 9;
	/* see bellow */
}
```
### SearchOpts

```protobuf
message SearchOpts{
	optional Category category = 1;
	/* see bellow */

	message DistanceOpts{
		DistanceInterval distance_interval = 1;
		google.type.Latlng search_area_center = 2;
	}
	optional DistanceOpts distance_opts = 2;
	/* see above */

	optional RatingInterval rating_interval = 3;
	/* see bellow */

	optional CheckInterval check_interval = 4;
	/* see bellow */

	optional WorkingHours working_hours = 5;
	/* see bellow */
}
```

### DistanceInterval

```protobuf
message DistanceInterval{
	uint32 start = 1;
	uint32 stop = 2;
}
```

### RatingInterval

```protobuf
message RatingInterval{
	// TODO
}
```

### Address

```protobuf
message Address{
	required string house = 1;
	required string street = 2;
	required string city = 3;
}
```

### Category

```protobuf
enum Category {
	CATEGORY_UNSPECIFIED = 0;
	AMERICAN = 5;
	GEORGIAN = 4;
	ITALIAN = 1;
	JAPANESE = 3;
	RUSSIAN = 2;
	KAFE = 6;
	STEAK = 7;
	CONFECTIONERY = 8;
	BAR = 9;
	PAB = 10;
	BEER_HOUSE = 12;
	COFFEE_HOUSE = 11;
	VEGAN_MENU = 13;
}
```

### AverageCheck

```protobuf
message AverageCheck{
	google.type.Money average = 1;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/money.proto */
}
```

### CheckInterval

```protobuf
message CheckInterval{
	optional google.type.Money min = 2;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/money.proto */

	optional google.type.Money max = 3;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/money.proto */
}
```

### WorkingHours

```protobuf
message WorkingHours{
	google.type.DayOfWeek day_of_week = 1;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/dayofweek.proto */

	google.type.TimeOfDay open = 2;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/timeofday.proto */

	google.type.TimeOfDay close = 3;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/timeofday.proto */
}
```

### Rating

```protobuf
message Rating{
	uint32 visit = 1;
	uint32 like = 2;
	uint32 dislike = 3;
}
```
