```protobuf

message PointOfInterest{
	required string uuid = 1;

	message Address{
		required string house = 1;
		required string street = 2;
		required string city = 3;
	}
	Address address = 2;
	string name = 3;

	enum Category {
		UNDEFINED = 0;
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
	Category category = 4;
	google.type.Latlng lat_lng = 5;

	message AveragePrice{
		google.type.Money min = 1;
		google.type.Money max = 2;
	}
	AveragePrice average_price = 6;

	message WorkingHours{
		google.protobuf.Timestamp from = 1;
		google.protobuf.Timestamp to = 2;
	}
	WorkingHours working_hours = 7;
		
	message Rating{
		uint32 dislike = 1;
		uint32 like = 2;
	}
	Rating rating = 8;
}
```