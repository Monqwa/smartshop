import axios from 'axios';

const API_URL = 'http://localhost:8080'; 

export default {
  createList(title) {
    return axios.post(`${API_URL}/list/`, { title });
  },

  updateList(listId, title) {
    return axios.put(`${API_URL}/list/${listId}`, { title });
  }
}; 