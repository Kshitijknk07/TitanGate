import { gql } from 'apollo-server-fastify';

const userSchema = gql`
  type User {
    id: Int!
    name: String!
    email: String!
  }

  type Query {
    getUsers: [User]
    getUser(id: Int!): User
  }
`;

export default userSchema;
