import RestToGraphQL from '../dataSources/RestToGraphQL.js';  // Import the REST-to-GraphQL class

// Create an instance of RestToGraphQL for users
const restToGraphQL = new RestToGraphQL('https://jsonplaceholder.typicode.com'); // Replace with your REST API base URL

const userResolvers = {
  Query: {
    getUsers: async () => {
      try {
        const users = await restToGraphQL.generateQuery('/users');
        return users;  // Return users as GraphQL response
      } catch (error) {
        throw new Error(`Error fetching users: ${error.message}`);
      }
    },

    getUser: async (_, { id }) => {
      try {
        const user = await restToGraphQL.generateQuery(`/users/${id}`);
        return user;  // Return a single user as GraphQL response
      } catch (error) {
        throw new Error(`Error fetching user with ID ${id}: ${error.message}`);
      }
    }
  }
};

export default userResolvers;
