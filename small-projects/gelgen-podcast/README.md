# NHL API project

This is cloned from [Learn Go by building small projects - P02 GraphQL itunes podcast](https://www.youtube.com/watch?v=doLybrf3w-o&list=PLzQWIQOqeUSOM9D9yk4mQ25ZiZEahJ3XQ&index=3)

## Executing

`go run main.go`

This will launch the GraphQL playground at localhost:8080

## Query in GraphQL playground

```
query my_search{
  search(
    term:"Full Stack Radio"
  ) {
    artist
    podcastName
    episodesCount
    feedUrl
  }
}

query get_feed{
  feed(
    feedUrl: "https://feeds.simplecast.com/Gd37VcDw"
  ) {
    title
  }
}
```

## Resources

* [iTunes Search API](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/)
