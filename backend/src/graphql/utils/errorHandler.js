class ErrorHandler {
    static handleError(error, context) {
      if (error instanceof Error) {
        console.error(error.message, error.stack);
      } else {
        console.error(`[ERROR] - ${error}`);
      }
  
      return {
        success: false,
        message: error.message || 'Something went wrong. Please try again later.',
      };
    }
  
    static throwError(message) {
      throw new Error(message);
    }
  }
  
  module.exports = ErrorHandler;
  