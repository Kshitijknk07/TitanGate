const fetch = require('node-fetch');

class RestToGraphQL {
  constructor(baseURL) {
    this.baseURL = baseURL;  // The base URL of your REST API
  }

  // Generic method to convert a REST endpoint to a GraphQL query
  async generateQuery(endpoint) {
    try {
      const response = await fetch(`${this.baseURL}${endpoint}`);
      const data = await response.json();
      if (!response.ok) throw new Error('Failed to fetch data from REST API');
      return data;  // Return the response data in a format compatible with GraphQL
    } catch (error) {
      throw new Error(`Error in REST-to-GraphQL conversion: ${error.message}`);
    }
  }

  // Additional utility method for handling specific API logic (if needed)
  async generateQueryWithParams(endpoint, params) {
    const url = new URL(`${this.baseURL}${endpoint}`);
    Object.keys(params).forEach(key => url.searchParams.append(key, params[key]));

    try {
      const response = await fetch(url);
      const data = await response.json();
      if (!response.ok) throw new Error('Failed to fetch data from REST API');
      return data;
    } catch (error) {
      throw new Error(`Error in REST-to-GraphQL conversion: ${error.message}`);
    }
  }
}

module.exports = RestToGraphQL;
