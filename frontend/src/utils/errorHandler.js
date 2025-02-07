export const handleApiError = (error, fallbackValue = null) => {
  console.error("API Error:", error);
  // Add user notification system here
  return fallbackValue;
}; 