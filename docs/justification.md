# Justifications

## Data format

I have chosen a relational database (Postgres) because the data follows a rigid format. I could have used a no sql databased seeing as the data is already a map, but I find no benefit in this other than not having to change the format of the data when initially reading from a file.

## Documentation

I would have liked to use OpenAPI given more time.

## Design

Because this API is so simple, there is no need to create separate Service and Persistence layers, because the persistence layer would almost be non-existent.