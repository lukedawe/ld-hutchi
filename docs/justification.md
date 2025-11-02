# Justifications

## Data format

I have chosen a relational database (Postgres) because the data follows a rigid format. I could have used a no sql databased seeing as the data is already a map, but I find no benefit in this other than not having to change the format of the data when initially reading from a file.

## Design

Because this API is so simple, there is no need to create separate service and persistence layers, because the persistence layer would almost be non-existent.

## Technologies

I have used Hutchi's tech normal stack because:

- It enables Hutchi to see the standard of work that I would produce on day 1.
- It's a modern and proven stack with good tooling and documentation.

For development I used Fedora Linux, VSCode.

## Hosting

I have self hosted the API and NGinx instance on a Raspberry PI.

The database is hosted on Supabase.

## What I'm not happy with

I am not paricularly happy with the website codebase, especially the sending and receiving of requests. I have used Zod to make handling casting better, but there is still a lot of repeated code for sending requests. Given more tiem I would create a library for sending requests build on top of zod that would expose http methods (`POST()`, `GET()` etc.) and handle the casting and error handling.

Some of the Gorm codebase I am also not happy with. I don't feel that this is a failure with my code particularly, but more a failing of the Gorm library.