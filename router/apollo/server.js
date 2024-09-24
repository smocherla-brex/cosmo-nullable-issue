"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var server_1 = require("@apollo/server");
var standalone_1 = require("@apollo/server/standalone");
var gateway_1 = require("@apollo/gateway");
var gateway = new gateway_1.ApolloGateway({
    supergraphSdl: new gateway_1.IntrospectAndCompose({
        subgraphs: [
            { name: "employees", url: "http://localhost:4001/graphql" },
            { name: "family", url: "http://localhost:4002/graphql" },
            { name: "hobbies", url: "http://localhost:4003/graphql" },
            { name: "products", url: "http://localhost:4004/graphql" },
        ],
    }),
});
var server = new server_1.ApolloServer({ gateway: gateway });
// Note the top-level await!
(0, standalone_1.startStandaloneServer)(server).then(function (url) {
    console.log("\uD83D\uDE80  Apollo Server ready at ".concat(url));
});
