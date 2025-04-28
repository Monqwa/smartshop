<template>
  <div class="create-list-view">
    <h2>Create a New Shopping List</h2>
    <div class="form-container">
      <input 
        type="text" 
        v-model="listTitle"
        placeholder="Enter list title"
        @keyup.enter="handleCreateList"
        :disabled="isCreating"
      />
      <button @click="handleCreateList" :disabled="isCreating">
        {{ isCreating ? 'Creating...' : 'Create List' }}
      </button>
      <p v-if="error" class="error-message">{{ error }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import listService from '../services/listService';

const listTitle = ref('');
const isCreating = ref(false);
const error = ref(null);
const router = useRouter();

async function handleCreateList() {
  if (!listTitle.value.trim()) {
    error.value = 'List title cannot be empty.';
    return;
  }

  isCreating.value = true;
  error.value = null;

  try {
    const response = await listService.createList(listTitle.value);
    const newListId = response.data?.id;

    if (!newListId) {
      console.error('Backend did not return a list ID.', response.data);
      throw new Error('Failed to get list ID after creation.');
    }

    router.push({ name: 'ListView', params: { listId: newListId } });

  } catch (err) {
    console.error('Error creating list:', err);
  }
}
</script>

<style scoped>
.create-list-view {
  max-width: 450px;
  margin: 3rem auto;
  padding: 2rem 2.5rem;
  text-align: center;
  background-color: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 4px;
  box-shadow: none;
}

h2 {
  margin-top: 0;
  margin-bottom: 2rem;
  color: #333;
  font-weight: 600;
}

.form-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

input[type="text"] {
  padding: 0.7rem 0.9rem;
  font-size: 1rem;
  border: 1px solid #ccc;
  border-radius: 3px;
  background-color: #fff;
  color: #333;
}
input[type="text"]:focus {
  outline: none;
  border-color: #555;
}

button {
  padding: 0.7rem 1.5rem;
  font-size: 1rem;
  color: #fff;
  background-color: #333;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

button:hover {
  background-color: #555;
}

button:disabled {
  background-color: #ccc;
  color: #777;
  cursor: not-allowed;
}

.error-message {
  color: #333;
  font-weight: bold;
  margin-top: 0.5rem;
  font-size: 0.9em;
}
</style> 