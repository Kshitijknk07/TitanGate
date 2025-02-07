import { gql } from 'apollo-server-fastify';

const postSchema = gql`
  type Post {
    id: Int!
    title: String!
    body: String
  }

  type Query {
    getPosts: [Post]
    getPost(id: Int!): Post
  }
`;

export default postSchema;
