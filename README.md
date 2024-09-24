## Example to illustrate difference between cosmo and apollo on strictness regarding nullable field of types which have non-nullable fields

This repository encompasses reproduction of an issue that illustrates a difference between the apollo and cosmo implementations of the GraphQL router
with regards to handling nullable fields with non-nullable fields in their types.

This example was built on top of the [cosmo-federation-demos](https://github.com/wundergraph/cosmo-federation-demos) repository.

A new entity type `Foo` is added to both the `employees` and `family` subgraph
```graphql
# in employees
type Foo @key(fields: "id") {
    id: Int!
}

# in family
type Foo @key(fields: "id") {
    id: Int!
    message: String!
}
```

The entity `Foo` has two non-nullable fields. With the above example, the following query behaves differently with `cosmo` and `apollo-gateway` routers


```graphql
query {
 	foo {
        id
        message
    }
}
```

With apollo-gateway, we get

```json
{
  "data": null,
  "extensions": {
    "valueCompletion": [
      {
        "message": "Cannot return null for non-nullable field Foo.message.",
        "path": [
          "foo",
          "message"
        ],
        "extensions": {
          "code": "INVALID_GRAPHQL"
        }
      }
    ]
  }
}
```

and 

```json
{
  "errors": [
    {
      "message": "Failed to fetch from Subgraph 'family' at Path 'foo', Reason: no data or errors in response.",
      "extensions": {
        "statusCode": 200
      }
    },
    {
      "message": "Cannot return null for non-nullable field 'Query.foo.message'.",
      "path": [
        "foo",
        "message"
      ]
    }
  ],
  "data": null,
```

One of the subgraphs (`employees`) returns a non-null `Foo` entity in the `_entities` resolver but in the `family` subgraph it resolves to `null`. The `message` field on `Foo` type is non-nullable
which is what `Cannot return null for non-nullable field` refers to. In `apollo-gateway` (and `apollo-router`), this has been enforced as such as described [here](https://github.com/apollographql/federation/issues/2374).


# Running locally

Make sure you have `node`, `npm` and `ts-node` installed.

Run the below to start the cosmo router with all the subgraphs
```bash
DEV_MODE=true npm start
```

Then separately run the below 
```bash
ts-node router/apollo/server.ts
```

to start the apollo-gateway server which talks to the same subgraphs as the cosmo router.

Now, you can go to `http://localhost:4000` to run queries against the apollo sandbox and 
`http://localhost:3002` as the playground and run the above query. You'll observe the difference.