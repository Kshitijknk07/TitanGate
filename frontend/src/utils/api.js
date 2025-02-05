// frontend/src/utils/api.js
export const login = async (username, password) => {
    const response = await fetch('http://localhost:3000/login', { // Update the URL to your backend
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });
  
    const data = await response.json();
  
    if (data.token) {
      localStorage.setItem('jwtToken', data.token); // Store token in localStorage
      return true;
    } else {
      return false;
    }
  };