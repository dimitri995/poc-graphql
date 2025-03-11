
![img.png](graphql/img.png)


# GraphQL Gateway in Golang

This project implements a GraphQL gateway in Golang that unifies data from multiple backend services (such as Product and User APIs) into a single GraphQL endpoint. This gateway allows clients to query all necessary data in one request, abstracting away the complexity of managing multiple services.

## Overview

The GraphQL gateway is built with Golang and the [gqlgen](https://gqlgen.com/) library. It serves as a centralized API that fetches and merges data from various backend services, such as:
- **Product Service:** Provides product-related data (e.g., product ID, name, price).
- **User Service:** Provides user-related data (e.g., user ID, first name, last name).

By using this gateway, clients benefit from a unified API that returns exactly the data they need with minimal overhead.

## Benefits

- **Unified API Endpoint:**  
  Clients only need to call one endpoint, simplifying integration and reducing the number of network calls.

- **Flexible and Efficient Data Fetching:**  
  GraphQL allows clients to request only the fields they need, optimizing data transfer and reducing payload sizes.

- **Scalability and Performance:**  
  Implemented in Golang, the gateway leverages Go's concurrency model to handle multiple backend API calls concurrently, ensuring high performance even under heavy load.

- **Decoupled Architecture:**  
  Each backend service can evolve independently. The gateway handles data merging and transformation, providing a consistent API to the clients.

- **Simplified Client Logic:**  
  The client does not need to know about the underlying services or their endpoints. All complexities are hidden behind the GraphQL layer.

## Architecture

- **GraphQL Schema:**  
  The schema defines types such as `Product`, `User`, and a composite type like `Dashboard` (or allows separate queries) that aggregates data from multiple services.

- **Resolvers:**  
  Each resolver calls the appropriate backend API (via HTTP) and merges the responses as needed. For instance, the `Dashboard` resolver makes concurrent calls to both Product and User services.

- **Backend Services:**  
  These services provide domain-specific data and can be maintained independently from the gateway.



