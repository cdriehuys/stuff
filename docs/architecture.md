# Architecture

Discussions about why certain implementations are used, or past decisions that
didn't work out.

## Using Generating API Clients

[`openapi-react-query`][openapi-react-query] is a type-safe wrapper around
[`@tanstack/react-query`][tanstack-query] which uses the types generated from an
Open API spec using [`openapi-typescript`][openapi-ts].

Initially, we were using this wrapper for all queries, but the rigidity of using
exactly the generated client made implementation difficult. For example, when
deleting a resource, the API returns a 204 status if the resource was deleted,
and a 404 if not found. For the web client, either of those responses should be
considered a success, so it ends up being simpler to write our own wrapper
around the [`openapi-fetch`][openapi-fetch] client.

Furthermore, extracting a custom API client makes it easier to fetch data on the
server.

[openapi-react-query]: https://openapi-ts.dev/openapi-react-query/
[openapi-fetch]: https://openapi-ts.dev/openapi-fetch/
[openapi-ts]: https://openapi-ts.dev/
[tanstack-query]: https://tanstack.com/query/latest/docs/framework/react/overview
