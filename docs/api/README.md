```proto3
message PointOfInterest{
	optional string uuid = 1; // required
	optional Address address = 2;
	/* see bellow */

	optional string name = 3;
	optional Category category = 4;
	/* see bellow */

	optional google.type.Latlng lat_lng = 5;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/latlng.proto */

	oneof check{
		optional AverageCheck average_check = 6;
		/* see bellow */

		optional CheckInterval check_interval = 7;
		/* see bellow */
	}
	repeated WorkingHours working_hours = 8;
	/* see bellow */

	optional Rating rating = 9;
	/* see bellow */
}
```
### SearchOpts

```proto3
message SearchOpts{
	optional Category category = 1;
	/* see bellow */

	message DistanceOpts{
		optional DistanceInterval distance_interval = 1;
		optional google.type.Latlng search_area_center = 2;
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

```proto3
message DistanceInterval{
	optional uint32 start = 1;
	optional uint32 stop = 2;
	/* meters */
}
```

### RatingInterval

```proto3
message RatingInterval{
	// TODO
}
```

### Address

```proto3
message Address{
	optional string house = 1;
	optional string street = 2;
	optional string city = 3;
}
```

### Category

```proto3
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

```proto3
message AverageCheck{
	optional google.type.Money average = 1;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/money.proto */
}
```

### CheckInterval

```proto3
message CheckInterval{
	optional google.type.Money min = 2;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/money.proto */

	optional google.type.Money max = 3;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/money.proto */
}
```

### WorkingHours

```proto3
message WorkingHours{
	optional google.type.DayOfWeek day_of_week = 1;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/dayofweek.proto */

	optional google.type.TimeOfDay open = 2;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/timeofday.proto */

	optional google.type.TimeOfDay close = 3;
	/* see https://github.com/googleapis/googleapis/blob/master/google/type/timeofday.proto */
}
```

### Rating

```proto3
message Rating{
	optional uint32 visit = 1;
	optional uint32 like = 2;
	optionaluint32 dislike = 3;
}
```
