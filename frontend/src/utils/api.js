// frontend/src/utils/api.js
const API_URL = import.meta.env.VITE_API_URL;

export const login = async (username, password) => {
    const response = await fetch(`${API_URL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });
  
    const data = await response.json();
  
    if (data.token) {
      localStorage.setItem('jwtToken', data.token);
      return true;
    } else {
      return false;
    }
  };

export const authenticatedFetch = async (url, options = {}) => {
  const token = localStorage.getItem('jwtToken');
  return fetch(url, {
    ...options,
    headers: {
      ...options.headers,
      'Authorization': `Bearer ${token}`,
    },
  });
};