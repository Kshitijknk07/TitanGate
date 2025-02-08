import RestToGraphQL from '../dataSources/RestToGraphQL.js';  // Correct import for default export

// Create an instance of RestToGraphQL for posts
const restToGraphQL = new RestToGraphQL('https://jsonplaceholder.typicode.com'); // Replace with your REST API base URL

const postResolvers = {
  Query: {
    getPosts: async () => {
      try {
        const posts = await restToGraphQL.generateQuery('/posts');
        return posts;  // Return posts as GraphQL response
      } catch (error) {
        throw new Error(`Error fetching posts: ${error.message}`);
      }
    },

    getPost: async (_, { id }) => {
      try {
        const post = await restToGraphQL.generateQuery(`/posts/${id}`);
        return post;  // Return a single post as GraphQL response
      } catch (error) {
        throw new Error(`Error fetching post with ID ${id}: ${error.message}`);
      }
    }
  }
};

export default postResolvers;
