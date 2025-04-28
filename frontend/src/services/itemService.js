import axios from 'axios';

const API_URL = 'http://localhost:8080'; 

axios.interceptors.request.use(request => {
  return request;
});

axios.interceptors.response.use(response => {
  return response;
}, error => {
  return Promise.reject(error);
});

export default {
  getItemsForList(listId) {
    return axios.get(`${API_URL}/items/${listId}`);
  },

  createItem(listId, name) {
    return axios.post(`${API_URL}/items/${listId}`, { name });
  },

  updateItem(itemId, data) {
    const payload = {};
    if (data.name !== undefined) payload.name = data.name;
    if (data.isBuy !== undefined) payload.is_buy = data.isBuy; 
    
    return axios.put(`${API_URL}/items/${itemId}`, payload);
  },

  deleteItem(itemId) {
    return axios.delete(`${API_URL}/items/${itemId}`);
  }
}; 