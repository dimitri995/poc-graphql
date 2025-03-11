const { ApolloServer } = require('apollo-server');
const { ApolloGateway, IntrospectAndCompose } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
            { name: 'products', url: 'http://localhost:8081/graphql' },
            { name: 'users', url: 'http://localhost:8088/query' }
        ]
    })
});

const server = new ApolloServer({
    gateway,

    subscriptions: false,
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});