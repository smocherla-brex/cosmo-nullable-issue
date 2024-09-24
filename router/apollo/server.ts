import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from '@apollo/server/standalone';
import { ApolloGateway, IntrospectAndCompose } from "@apollo/gateway";

const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
            { name: "employees", url: "http://localhost:4001/graphql" },
            { name: "family", url: "http://localhost:4002/graphql" },
            { name: "hobbies", url: "http://localhost:4003/graphql" },
            {name: "products", url: "http://localhost:4004/graphql"},
        ],
    }),
});

const server = new ApolloServer({ gateway });

startStandaloneServer(server).then((res) => {
    console.log(`🚀  Apollo Server ready at ${res.url}`);
})
