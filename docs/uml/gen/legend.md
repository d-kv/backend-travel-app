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

	oneof Check{
		AverageCheck average_check = 6;
		/* see bellow */

		BoundedCheck bounded_check = 7;
		/* see bellow */
	}
	repeated WorkingHours working_hours = 8;
	/* see bellow */

	Rating rating = 9;
	/* see bellow */
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
	ITALIAN = 1;
	RUSSIAN = 2;
	JAPANESE = 3;
	GEORGIAN = 4;
	AMERICAN = 5;
	KAFE = 6;
	STEAK = 7;
	CONFECTIONERY = 8;
	BAR = 9;
	PAB = 10;
	COFFEE_HOUSE = 11;
	BEER_HOUSE = 12;
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

### BoundedCheck

```protobuf
message BoundedCheck{
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
	uint32 like = 1;
	uint32 dislike = 2;
}
```