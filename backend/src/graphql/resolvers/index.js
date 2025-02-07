import postResolvers from './postResolvers.js';
import userResolvers from './userResolvers.js';

const resolvers = {
  Query: {
    ...postResolvers.Query,
    ...userResolvers.Query
  }
};

export default resolvers;
