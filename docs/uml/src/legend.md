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
		italian = 1;
		russian = 2;
		japanese = 3;
		georgian = 4;
		american = 5;
		kafe = 6;
		steak = 7;
		confectionery = 8;
		bar = 9;
		pab = 10;
		coffee_house = 11;
		beer_house = 12;
		vegan_menu = 13;
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
}
```