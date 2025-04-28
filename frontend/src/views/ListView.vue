<template>
  <div class="list-view-container">
    <div class="list-view">
      <div class="list-header">
          <h2 class="list-title">{{ list.name }}</h2> 
          <p class="list-description">{{ list.description }}</p>
          <div class="header-actions">
              <button @click="copyLink" class="action-btn icon-btn" title="Copy shareable link">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-link-45deg" viewBox="0 0 16 16">
                    <path d="M4.715 6.542 3.343 7.914a3 3 0 1 0 4.243 4.243l1.828-1.829A3 3 0 0 0 8.586 5.5L8 6.086a1.002 1.002 0 0 0-.154.199 2 2 0 0 1 .861 3.337L6.88 11.45a2 2 0 1 1-2.83-2.83l.793-.792a4.018 4.018 0 0 1-.128-1.287z"/>
                    <path d="M6.586 4.672A3 3 0 0 0 7.414 9.5l.775-.776a2 2 0 0 1-.896-3.346L9.12 3.55a2 2 0 1 1 2.83 2.83l-.793.792c.112.42.155.855.128 1.287l1.372-1.372a3 3 0 1 0-4.243-4.243L6.586 4.672z"/>
                  </svg>
                  <span>Copy Link</span>
              </button>
          </div>
      </div>

      <div class="items-section">
        <div v-if="isLoading" class="loading-container">
          <div class="loading-spinner"></div>
          <p>Loading your shopping list...</p>
        </div>

        <div v-else>
          <div class="list-stats">
            <span class="stat-item">
              <span class="stat-value">{{ items.length }}</span> total items
            </span>
            <span class="stat-item">
              <span class="stat-value">{{ boughtItemsCount }}</span> purchased
            </span>
          </div>

          <div class="items-container">
            <div v-if="items.length === 0" class="empty-list">
              <p>Your shopping list is empty</p>
              <p class="empty-list-hint">Add your first item below</p>
            </div>
            
            <transition-group name="list" tag="ul" class="items-list" v-else>
              <ItemComponent
                v-for="item in items"
                :key="item.id"
                :item="item"
                @toggle-bought="toggleItemBought"
                @delete-item="deleteItem"
                @update-item-name="updateItemName"
              />
            </transition-group>
          </div>

          <form @submit.prevent="addItem" class="add-item-form">
            <input
              v-model="newItemName"
              type="text"
              placeholder="Add a new item..."
              class="add-item-input"
              required
              :disabled="isAddingItem"
              ref="newItemInput"
            />
            <button type="submit" class="add-button" :disabled="isAddingItem || !newItemName.trim()">
              <span v-if="isAddingItem">Adding...</span>
              <span v-else>Add Item</span>
            </button>
          </form>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, defineProps, watch, computed } from 'vue'; 
import itemService from '../services/itemService';
import ItemComponent from '../components/ItemComponent.vue';

const props = defineProps({
  listId: {
    type: String,
    required: true
  }
});

const list = ref({
  name: "Shopping List",
  description: "Your collaborative shopping list"
});

const items = ref([]);
const newItemName = ref('');
const isLoading = ref(true);
const isAddingItem = ref(false);
const isPollingActive = ref(false);
let pollTimeoutId = null;

const boughtItemsCount = computed(() => {
  return items.value.filter(item => item.is_buy).length;
});

function processBackendItems(data) {
    if (Array.isArray(data)) {
        return data.map(item => ({
            id: item.id || item.ID,
            name: item.name || item.Name,
            is_buy: item.is_buy !== undefined ? item.is_buy : item.IsBuy,
            list_id: item.list_id || item.ListID
        }));
    }
    if (typeof data === 'object') {
        return Object.values(data).map(item => ({
            id: item.id,
            name: item.name,
            is_buy: item.is_buy,
            list_id: item.list_id
        }));
    }
    return [];
}

async function pollItems() {
  if (!isPollingActive.value) return;
  try {
    const response = await itemService.getItemsForList(props.listId);
    if (isPollingActive.value) {
        items.value = processBackendItems(response.data);
    }
  } catch (err) {
     if (isPollingActive.value) {
        console.error(`Error polling items for list ${props.listId}:`, err);
        if (err.response?.status === 404) {
          stopPolling();
        }
     }
  } finally {
     if (isPollingActive.value) {
        clearTimeout(pollTimeoutId);
        pollTimeoutId = setTimeout(pollItems, 1500);
        if (isLoading.value) isLoading.value = false;
     }
  }
}

async function addItem() {
  isAddingItem.value = true;
  try {
    await itemService.createItem(props.listId, newItemName.value);
    newItemName.value = '';
    clearTimeout(pollTimeoutId); 
    pollItems(); 
  } catch (err) {
     console.error(`Error adding item to list ${props.listId}:`, err);
  }
  isAddingItem.value = false;
}

async function toggleItemBought(itemId, newStatus) {
  const itemRef = items.value.find(item => item.id === itemId);
  
  const originalStatus = itemRef.is_buy;
  itemRef.is_buy = newStatus;
  
  try {
    
    const response = await itemService.updateItem(itemId, { isBuy: newStatus });
    
    if (isPollingActive.value) {
      clearTimeout(pollTimeoutId);
      pollItems();
    }
  } catch (err) {
    console.error(`Error updating item ${itemId}:`, err);
  }
}

async function deleteItem(itemId) {
  if (!confirm('Delete this item?')) return;
  const originalItems = [...items.value];
  items.value = items.value.filter(item => item.id !== itemId);
  isAddingItem.value = true;
  try {
    await itemService.deleteItem(itemId);
    clearTimeout(pollTimeoutId); 
    pollItems();
  } catch (err) {
    console.error(`Error deleting item ${itemId}:`, err);
  }
  isAddingItem.value = false;
}

async function updateItemName(itemId, newName) {
    const itemRef = items.value.find(item => item.id === itemId);
    if (!itemRef || itemRef.name === newName) return;
    const originalName = itemRef.name;
    itemRef.name = newName;
    try {
        await itemService.updateItem(itemId, { name: newName });
    } catch (err) {
        console.error(`Error updating item ${itemId}:`, err);
    }
}

function copyLink() {
    const url = window.location.href;
    navigator.clipboard.writeText(url).then(() => {
        alert('Link copied to clipboard!');
    }, () => {
        alert('Failed to copy link.');
    });
}

async function fetchItems() {
  isLoading.value = true;
  try {
    const response = await itemService.getItemsForList(props.listId);
    items.value = processBackendItems(response.data);
  } catch (err) {
    console.error(`Error fetching items for list ${props.listId}:`, err);
  } finally {
    isLoading.value = false;
  }
}

function initializeList() {
    startPolling();
}

function startPolling() {
    if (!props.listId || isPollingActive.value) return;
    isLoading.value = true;
    isPollingActive.value = true;
    clearTimeout(pollTimeoutId);
    pollItems();
}

function stopPolling() {
    if (isPollingActive.value) {
        isPollingActive.value = false;
        clearTimeout(pollTimeoutId);
    }
}

onMounted(initializeList);
onUnmounted(stopPolling);

watch(() => props.listId, (newListId, oldListId) => {
  if (newListId && newListId !== oldListId) {
    stopPolling();
    items.value = []; 
    initializeList();
  }
});
</script>

<style scoped>
.list-view-container {
  padding: var(--space-md);
}

.list-view {
  max-width: 700px;
  margin: var(--space-xl) auto;
  padding: var(--space-xl);
  background-color: var(--color-white);
  border-radius: var(--border-radius-lg);
  box-shadow: var(--shadow-lg);
}

.list-header {
  display: flex;
  flex-direction: column;
  margin-bottom: var(--space-xl);
  padding-bottom: var(--space-md);
  border-bottom: var(--border-width-thin) solid var(--color-gray-200);
}

.list-title {
  margin: 0 0 var(--space-xs) 0;
  font-size: var(--font-size-xxl);
  font-weight: var(--font-weight-bold);
  color: var(--color-gray-900);
  letter-spacing: -0.03em;
}

.list-description {
  color: var(--color-gray-600);
  font-size: var(--font-size-md);
  margin-bottom: var(--space-md);
}

.header-actions {
  display: flex;
  gap: var(--space-sm);
  align-self: flex-start;
}

.action-btn {
  display: inline-flex;
  align-items: center;
  gap: var(--space-xs);
  padding: var(--space-sm) var(--space-md);
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-medium);
  background-color: var(--color-white);
  border: var(--border-width-thin) solid var(--color-gray-300);
  color: var(--color-gray-700);
  border-radius: var(--border-radius-md);
  transition: all var(--transition-normal);
}

.action-btn:hover {
  background-color: var(--color-gray-100);
  border-color: var(--color-gray-400);
  transform: translateY(-1px);
}

.action-btn svg {
  flex-shrink: 0;
}

.list-stats {
  display: flex;
  justify-content: space-between;
  margin-bottom: var(--space-md);
  background-color: var(--color-gray-100);
  padding: var(--space-sm) var(--space-md);
  border-radius: var(--border-radius-md);
  font-size: var(--font-size-sm);
  color: var(--color-gray-700);
}

.stat-item {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
}

.stat-value {
  font-weight: var(--font-weight-bold);
  font-size: var(--font-size-md);
  color: var(--color-gray-900);
}

/* Items Container */
.items-section {
  margin-top: var(--space-lg);
}

.items-container {
  margin-bottom: var(--space-xl);
  min-height: 100px;
}

.empty-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: var(--space-xl) 0;
  color: var(--color-gray-500);
  text-align: center;
  background-color: var(--color-gray-50);
  border-radius: var(--border-radius-md);
  border: 1px dashed var(--color-gray-300);
}

.empty-list p {
  font-size: var(--font-size-md);
  margin-bottom: var(--space-xs);
}

.empty-list-hint {
  font-size: var(--font-size-sm);
  opacity: 0.8;
}

.items-list {
  padding: 0;
  margin: 0;
}

.add-item-form {
  display: flex;
  gap: var(--space-sm);
  margin-top: var(--space-lg);
  padding-top: var(--space-lg);
  border-top: var(--border-width-thin) solid var(--color-gray-200);
}

.add-item-input {
  flex: 1;
  padding: var(--space-md);
  border: var(--border-width-thin) solid var(--color-gray-300);
  border-radius: var(--border-radius-md);
  font-size: var(--font-size-md);
  box-shadow: var(--shadow-sm);
  transition: all var(--transition-normal);
}

.add-item-input:focus {
  outline: none;
  border-color: var(--color-gray-600);
  box-shadow: var(--shadow-md);
}

.add-button {
  padding: var(--space-sm) var(--space-lg);
  background-color: var(--button-background);
  color: var(--button-text);
  border-radius: var(--border-radius-md);
  font-weight: var(--font-weight-medium);
  font-size: var(--font-size-md);
  transition: all var(--transition-normal);
  min-width: 120px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.add-button:hover:not(:disabled) {
  background-color: var(--color-black);
  transform: translateY(-1px);
}

.add-button:disabled {
  background-color: var(--color-gray-300);
  cursor: not-allowed;
  color: var(--color-gray-500);
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: var(--space-xl) 0;
  color: var(--color-gray-600);
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid var(--color-gray-200);
  border-radius: 50%;
  border-top-color: var(--color-gray-800);
  animation: spin 1s linear infinite;
  margin-bottom: var(--space-md);
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}

.list-enter-from {
  opacity: 0;
  transform: translateX(-20px);
}

.list-leave-to {
  opacity: 0;
  transform: translateX(20px);
}

</style> 